ES 是一个分布式搜索和分析引擎，用于高效全文搜索。
与 Mysql 存储后查询相比，ES支持用户快速搜索“包含某关键词的文章，以及实现“附近动态”或“热门话题”等复杂排序和过滤
## 经典例子
GO 相关包：`go get -u github.com/elastic/go-elasticsearch/v8`
docker 快速部署：`docker run -d --name elasticsearch -p 9200:9200 -e "discovery.type=single-node" elasticsearch:8.14.0`

```bash
# 启用现有容器
docker start elasticsearch

#如果希望重新部署：
# 首先停止容器（如果它在运行）
docker stop elasticsearch

# 然后移除容器
docker rm elasticsearch

# 查看容器
docker ps -a | grep elasticsearch
```

启动后你会发现，访问不行。这是因为 es 设置了密码和安全验证，那么我们在开发环境下可以这么启动：
``` bash
docker run -d --name elasticsearch \
  -p 9200:9200 \
  -e "discovery.type=single-node" \
  -e "xpack.security.enabled=false" \
  elasticsearch:8.14.0
```
记得先用上面的指令暂停之前的容器并移除，重新 docker 启动

- 示范 demo 代码
``` go
package main

import (
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	// 索引一个文档
	_, err = es.Index("posts", strings.NewReader(`{"title":"Test Post","content":"Hello world"}`))
	if err != nil {
		log.Fatalf("Error indexing document: %v", err)
	}

	// 搜索
	res, err := es.Search(
		es.Search.WithIndex("posts"),
		es.Search.WithQuery("hello"),
	)
	if err != nil {
		log.Fatalf("Error searching: %v", err)
	}
	defer res.Body.Close()
	log.Println(res.String())
}

```

发现看到的结果不太能看懂，进行进一步解析：
- 优化后的代码
``` Go
package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	// 在搜索前刷新索引
	_, err = es.Indices.Refresh(es.Indices.Refresh.WithIndex("posts"))
	if err != nil {
		log.Fatalf("Error refreshing index: %v", err)
	}

	// 索引一个文档
	_, err = es.Index("posts", strings.NewReader(`{"title":"Test Post","content":"Hello world"}`))
	if err != nil {
		log.Fatalf("Error indexing document: %v", err)
	}

	// 搜索
	res, err := es.Search(
		es.Search.WithIndex("posts"),
		es.Search.WithBody(strings.NewReader(`{
			"query": {
				"match": {
					"content": "hello"
				}
			}
		}`)),
	)
	if err != nil {
		log.Fatalf("Error searching: %v", err)
	}
	defer res.Body.Close()

	// 解析搜索结果 （这个算是一个通用的模板）
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// 现在可以访问解析后的结果
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		log.Printf("Found document: %v", source)
	}
}

```
这样就可以看到很好的输出结果。
不过记得如果不注释，每次执行程序都会新建一个索引的

es 返回格式示例：
``` json
{
  "took": 10,                  // 查询耗时（毫秒）
  "timed_out": false,          // 是否超时
  "hits": {                    // 命中结果
    "total": {                 // 总匹配数
      "value": 2,              // 具体数量
      "relation": "eq"         // 计数关系（eq 表示精确值）
    },
    "max_score": 1.0,          // 最高相关性得分
    "hits": [                  // 文档数组
      {
        "_index": "my_index",  // 索引名
        "_id": "1",            // 文档ID
        "_score": 1.0,         // 当前文档得分
        "_source": {           // 原始文档数据   这是我们需要的！！
          "title": "Elasticsearch入门",
          "content": "学习ES基础用法"
        }
      }
    ]
  }
}
```


## Gin，gorm，es 项目
下面这个函数在解析 es 结果的时候会报空指针的错误，分析如下原因
```go
func SearchPosts(query string) ([]models.Post, error) {
	var posts []models.Post
	res, err := esClient.Search(
		esClient.Search.WithIndex("posts"),
		esClient.Search.WithQuery(`{"match": {"title": {"query": "`+query+`"}}}`), //查询针对的是文档的 title 字段，使用的是 match 查询，这是一种全文检索查询，用变量 query 的值作为搜索词
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 解析结果（简化处理，仅提取 hits）
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil { //将结果放入 result 中
		return nil, err
	}
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		posts = append(posts, models.Post{
			Title:   source["title"].(string), //结构化输出，放入结构体，然后放入列表中
			Content: source["content"].(string),
		})
	}
	return posts, nil
}
```
这个函数存在空指针错误，原因如下：
- 如果 ES 返回的响应中缺少 `hits` 字段（如查询语法错误或索引不存在），`result["hits"].(map[string]interface{})` 会触发 `panic`。
- 即使 `hits` 存在，若搜索结果为空（`hits.hits` 为空数组），循环内的 `hit(map[string]interface{})` 虽不会报错，但后续对 `_source` 的访问仍需校验
- 进一步的，如果某文档无 `title` 或 `content` 字段，`source["title"].(string)` 会因类型断言失败或字段不存在而崩溃


参照 es 的返回格式更好理解


所以可以采用逐层校验的方式：
``` go
hits, ok := result["hits"].(map[string]interface{})
if !ok {
    return nil, fmt.Errorf("invalid hits format in ES response")
}

hitList, ok := hits["hits"].([]interface{})
if !ok {
    return nil, fmt.Errorf("invalid hits.hits format")
}

for _, hit := range hitList {
    hitMap, ok := hit.(map[string]interface{})
    if !ok {
        continue // 跳过无效条目
    }
    source, ok := hitMap["_source"].(map[string]interface{})
    if !ok {
        continue
    }
    // 安全获取字段（支持缺省值）
    title, _ := source["title"].(string)    // 若字段不存在，title=""
    content, _ := source["content"].(string)
    posts = append(posts, models.Post{
        Title:   title,
        Content: content,
    })
}
```
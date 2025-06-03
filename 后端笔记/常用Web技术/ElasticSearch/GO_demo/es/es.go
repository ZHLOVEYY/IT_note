package es

import (
	"bytes"
	"encoding/json"
	"esdemo/models"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

func InitES() {
	esHost := os.Getenv("ELASTICSEARCH_HOST")
	if esHost == "" {
		esHost = "elasticsearch" // Docker环境中的默认值
	}
	cfg := elasticsearch.Config{
		Addresses: []string{"http://" + esHost + ":9200"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}
	esClient = client

	// 创建索引和映射   定义索引映射
	// 包含：title：类型为text，用于存储和搜索文本内容   content：类型为text，同样用于存储和搜索文本内容
	//注意：和 model 定义的需要一致！！
	mapping := `{    
        "mappings": {
            "properties": {
                "title": {"type": "text"},
                "content": {"type": "text"}
            }
        }
    }`
	//使用客户端的Indices.Create方法创建名为"posts"的索引，并应用上面定义的映射。如果索引已存在，会打印警告信息但不会中断程序执行
	_, err = esClient.Indices.Create("posts", esClient.Indices.Create.WithBody(strings.NewReader(mapping)))
	//可以理解成“posts”类的索引的结构，是上面那样的
	if err != nil {
		log.Printf("Index 'posts' may already exist: %v", err)
	}
}

func IndexPost(post *models.Post) error {
	data, err := json.Marshal(post) //这里可以看出是是用 json 解码，进一步验证了 model 定义的需要一致
	if err != nil {
		return err
	}
	_, err = esClient.Index("posts", bytes.NewReader(data)) //新建立存储
	return err
}

func SearchPosts(query string) ([]models.Post, error) {
	// 构建搜索请求 - 使用 multi_match 查询同时搜索 title 和 content 字段
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"title", "content"},  //同时查询 title 和 content 中的内容
			},
		},
	}

	// 将查询转换为 JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	// 执行搜索
	res, err := esClient.Search(
		esClient.Search.WithIndex("posts"),
		esClient.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, fmt.Errorf("search error: %w", err)
	}
	defer res.Body.Close()

	// 解析响应
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	// 提取结果 (使用安全类型断言)
	var posts []models.Post

	// 依次安全地提取 hits 对象和 hits 数组
	hitsObj, exists := result["hits"]
	if !exists {
		return posts, nil
	}

	hitsMap, ok := hitsObj.(map[string]interface{})
	if !ok {
		return posts, nil
	}

	hitsList, exists := hitsMap["hits"]
	if !exists {
		return posts, nil
	}

	hitsArray, ok := hitsList.([]interface{})
	if !ok {
		return posts, nil
	}

	// 处理搜索结果
	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}

		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		// 安全地提取字段值
		title, _ := source["title"].(string)
		content, _ := source["content"].(string)

		posts = append(posts, models.Post{
			Title:   title,
			Content: content,
		})
	}

	return posts, nil
}

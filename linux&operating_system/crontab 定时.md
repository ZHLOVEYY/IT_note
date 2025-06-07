## 基础
`crontab -e` 会进入设置窗口
crontab 的设定：
``` text
* * * * *
│ │ │ │ │
│ │ │ │ └─ 星期几 (0-7) (0 或 7 是星期日)
│ │ │ └─── 月份 (1-12)
│ │ └───── 日期 (1-31)
│ └─────── 小时 (0-23)
└───────── 分钟 (0-59)
```

例子： （结合 es 删除脚本例子理解）
``` bash
# 每天午夜执行
0 0 * * * /root/xxxx/cleanshell/clean-es-posts.sh

# 每周一凌晨执行
0 3 * * 1 /root/xxxxx/cleanshell/clean-es-posts.sh

# 每月一号执行
0 3 1 * * /root/xxxxx/cleanshell/clean-es-posts.sh

# 每天下午三点执行
0 15 * * * /root/xxxxx/cleanshell/clean-es-posts.sh

# 每小时执行
0 * * * * /root/ xxxxx/cleanshell/clean-es-posts.sh
```
## ES删除脚本例子
####  es 脚本
定时任务 -> 使用 crontab 
```  bash
#!/bin/bash
# 清理posts索引的内容（保留索引结构，仅删除数据
# 注意localhost还是在docker网络中（比如容器中）就需要是容器的名字
response=$(curl -s -XPOST "http://localhost:9200/（你的索引比如 indexs）/_delete_by_query" -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} }
}')

if [ $? -eq 0 ]; then  # - `$?`是获取上一个命令的退出状态码 ，看上面执行成功了没
    echo "$(date "+%Y-%m-%d %H:%M:%S") - 清理成功: $response" >> /root/xxxxxx/cleanshell/es_clear.log
else
    echo "$(date "+%Y-%m-%d %H:%M:%S") - 清理失败: $response" >> /root/xxxxxx/cleanshell/es_clear_error.log
    exit 1
fi
```
解释：
- `response=$( ... )` - 将命令执行结果保存到变量 `response` 中
- `curl` - 用于发送 HTTP 请求的命令行工具  （可以看我讲 curl 的一部分）
- `-s` - 静默模式，不显示进度条或错误信息
- `-XPOST` - 指定使用 HTTP POST 方法
- `"http://localhost:9200/posts/_delete_by_query"` - 请求的 URL
- `-H 'Content-Type: application/json'` - 设置请求头，指定内容类型为 JSON
- `-d'...'` - 指定请求体数据，后面跟着的是 JSON 格式的查询条件
请求体：
- `query` - 指定查询条件
- `match_all: {}` - 匹配所有文档的查询

\_delete_by_query是标准的 api 端点，专门用于批量删除文档 
#### 文件路径问题
需要考虑到工作目录的原因，所以不要用./ 
 比如当我手动执行：
``` bash
bash -x /root/xxxxx/cleanshell/clean-es-posts.sh
```
当我在不同的文件夹下执行，就会在不同的文件夹下输出新的文件了，所以还是采用绝对路径！



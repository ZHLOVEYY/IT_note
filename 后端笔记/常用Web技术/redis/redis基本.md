更多个人笔记：（仅供参考，非盈利）
gitee： https://gitee.com/harryhack/it_note
github： https://github.com/ZHLOVEYY/IT_note
- 其实redis并不是主要学习的中心，还是要基于语言进行学习掌握实现不同的模式和用法，redis-cli主要是辅助和调试，知道基本的功能就行
- （基于mac展示，别的可以参考）接下来将直接通过案例进行学习

## 简单了解概念
简单几条快速了解
- 属于非关系型数据库，no-sql
- redis用于缓解频繁访问sql带来的性能问题。redis性能高/
- 操作redis访问包括：CLI，API，GUI 三种方式访问
- 涵盖的数据类型包括：字符串，列表，集合，哈希，消息队列，位图，有序集合等
- 支持数据持久化，主从复制，哨兵模式等高可用特性
## 安装
这方面我用的是mac，直接homebrew 下载redis就可以了。这方面大伙稍微搜一下就有，下载很方便的
后面的展示也是基于mac

UI界面有一个叫“redisinsight”的，不过本着无非必要，勿增实体的想法，等发现确实有需要了再看看，基本的用终端操作就可以的

## 基础操作

``` bash
redis-server # 启动服务
redis-cli   # 启动客户端

# string
set user1 "Alice" # 设置键值对
get user1
setex session1 10 "token123" # 设置有时间限制的键值对
TTL session1 # 查看剩余的时间，可以多尝试几次
DEL user1 # 删除特定的键
get user1 # 发现是nil
DEL user # 我们没有设置过user，对比发现上面DEL输出1，这里输出0，可以理解成操作是否执行成功

# 操作Hash
HSET user2 name "Bob" age 30 # 返回值表示​​成功新增的字段数量​​（不包括被覆盖的已有字段） 这里会返回2   integer通常用于表示命令执行后影响的数据数量或状态变化​
HGET user2 name # "bob"
HGET user2 age  # "30"
HGETALL user2

# 操作List
RPUSH message_queue "msg1" "msg2" "msg3" # 推入消息
LPOP message_queue # 弹出消息，msg1
LRANGE message_queue 0 -1 # 查看队列所有元素

#需要唯一性 → ​​Set​​。
#需要顺序或重复元素 → ​​List​​。

# 操作set
SADD myset "apple" "banana" "apple" # 添加元素，自动去重
SMEMBERS myset # 查看集合

# 操作soeted Set
ZADD leaderboard 100 "player1" 200 "player2" # 添加带分数的元素
ZRANGE leaderboard 0 -1 WITHSCORES # 连着分数一起列出

```
redis 的键值对都是二进制安全的，所以设置数字，bool等也都会转为字符串的形式输出

可以发现其实redis主要就是对于主流的数据结构类型进行操作，简单了解后发现后面的使用是直接有需要查就可以的。主要还是要基于自己的语言进行学习

## 补充
######  redis 结合 sessionid对比直接cookie存储：
Redis 使用 Session ID 作为 key，存储对应的 session 数据，浏览器发送请求时带上 Cookie 中的 Session ID还是一样的
如果是cookie方法就会用“secret”加密然后直接放在cookie中
- 比如：
session.Set("userId", user.Id) session.Set("email", user.Email) 
mysession:abc123def456  （Redis 中的 session key） 对应就是-> {"userId": "10086","email": "test@example.com"} 



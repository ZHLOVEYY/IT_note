# Resources

# Core features
简单几条快速了解
- 属于非关系型数据库，no-sql
- redis用于缓解频繁访问sql带来的性能问题。redis性能高/
- 操作redis访问包括：CLI，API，GUI 三种方式访问
- 涵盖的数据类型包括：字符串，列表，集合，哈希，消息队列，位图，有序集合等
- 支持数据持久化，主从复制，哨兵模式等高可用特性

redis 的键值对都是二进制安全的，所以设置数字，bool等也都会转为字符串的形式输出
###### 数据持久化（RDB + AOF）
这个也是八股中很多的
Redis 的数据持久化通过 **RDB**（快照）和 **AOF**（追加日志）实现。Go 代码无需直接控制持久化（由 Redis 配置文件管理），但可以通过命令触发快照或检查持久化状态。
- RDB（Redis Database Backup） （案例）
	- 原理：
		- RDB 通过定期生成内存数据的快照（二进制文件，.rdb），保存到磁盘。
		- 快照是某一时刻的完整数据副本，文件体积较小，恢复速度快。
	- 触发方式：
		 - **自动触发**：根据 save 配置（如 save 900 1 表示 900 秒内至少 1 次变更触发）。
		- **手动触发**：
		    - SAVE：阻塞主线程，生成快照（生产慎用）。
		    - BGSAVE：后台异步生成快照，常用。
	- 优点：恢复速度比AOF快文件紧凑，适合备份和快速恢复
	- 缺点：可能丢失数据（两次快照间隔内的变更会丢失！！！），BGSAVE 需要 fork 进程，内存占用较高
	- **适用场景**：
		- 定期备份（如每天全量备份）。
		- 对少量数据丢失可接受的场景（如缓存）。
RDB的配置文件 redis.config (示范)
``` bahs
save 900 1    # 900秒内至少1次变更触发快照
save 300 10   # 300秒内至少10次变更
save 60 10000 # 60秒内至少10000次变更  这几个同时生效保证不同场景
dir /var/redis # 快照文件存储路径 
dbfilename dump.rdb # 快照文件名
```
触发快照的意思就是会进行一次更改更新（类似虚拟机的快照，但是这个直接覆盖）

- AOF（Append-Only File）
	- 原理
		- AOF 记录每次写操作（如 SET、DEL）到日志文件（.aof），每个写命令都会追加到 AOF 文件！！！类似数据库的 WAL（Write-Ahead Log）。
		- 重启时，Redis 重放 AOF 文件重建数据
	- **同步策略**（appendfsync）
		- always：每次写操作同步到磁盘，数据最安全但性能最低。
		- everysec：每秒同步，折中方案（最多丢 1 秒数据）。
		- no：依赖操作系统同步，性能最高但数据丢失风险大。
	- AOF重写
		- AOF 文件会随写操作不断增长，占用磁盘空间。
		- BGREWRITEAOF 合并冗余命令（如多次 INCR 合并为一个 SET），生成更小的 AOF 文件。
	- 优点：数据可靠性高，支持增量记录，适合高一致性场景
	- 缺点：文件体积较大恢复速度慢（需重放所有命令）。写频繁的时候性能开销高
	- **适用场景**：
		- 数据一致性要求高的场景（如订单、会话）。
		- 与 RDB 结合使用，兼顾恢复速度和可靠性。

AOF配置文件redis.config （示范）
``` bash
appendonly yes        # 启用 AOF
appendfsync everysec  # 每秒同步
dir /var/redis        # AOF 文件存储路径
appendfilename appendonly.aof # AOF 文件名
auto-aof-rewrite-percentage 100 # AOF 文件增长100%时触发重写
auto-aof-rewrite-min-size 64mb  # AOF 文件至少64MB时触发重写
```

redis-cli中手动输入可以出发RDB和AOF：
``` bash
BGSAVE # 手动触发快照
BGREWRITEAOF # 手动触发 AOF 重写
```


我的是mac，如果是homebrew启动的redis即通过指令 `brew services start redis ` 那么即使执行 `redis-cli shutdown` 也无法关闭，homebrew启动的是类似全局的redis，需要使用 `brew services start redis` 进行关闭
```  bash
# 查看当前运行的 Redis   可以用于检查问题
ps aux | grep redi
# 快速检查端口
lsof -i :6379  # 如果没有被占用就没有输出（显示为错误输出但实际没有输出的）
```

`redis-cli config get dir`  获取存储对应的目录，如果是homebrew启动的一眼就能看出来
`redis-cli config get dbfilename` 获取对应的dbfilename

redis.conf的demo：（结合了RDB和AOF）
``` bash
save 900 1
save 300 10
dir ./redisstorage
dbfilename dump.rdb
appendonly yes
appendfsync everysec
appendfilename appendonly.aof
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb
```
dir是目标存储文件夹，需要新建对应的文件夹
接着我们需要在当前文件夹的终端下执行 `redis-server ./redis.conf` 根据配置文件启动
# Related operations
``` bash
redis-server # 启动服务
redis-cli   # 启动客户端
```
#### Data type operation
``` bash
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
# Multilingual implementation

#### GO

安装redis客户端：`go get github.com/redis/go-redis/v9`
注意go mod tidy的时候不要import成了"github.com/go-redis/redis" 需要检查一下
######  connect redis && CRUD
``` GO
package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default 
		//password和db如果都是默认值可以不设置，这里展示一下
	})
	ctx := context.Background() //创建上下文

	//测试连接
	_,err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("连接失败",err)
	}
	fmt.Println("连接成功")

	///设置键值对
	err = client.Set(ctx,"user1","Alice",0).Err()  //0表示永不过期
	if err!=nil {
		log.Fatal("设置键值对失败",err)
	}
	fmt.Println("设置了user1=Alice")

	//获取键值对
	name,err := client.Get(ctx,"user1").Result()
	if err!=nil {
		log.Fatal("获取键值对失败",err)
	}
	fmt.Println("user1的值是",name)

	//设置过期时间（10秒）
	err = client.SetEx(ctx,"session1","123456",10*time.Second).Err()
	if err!=nil {
		log.Fatal("设置过期时间失败",err)
	}
	fmt.Println("设置了session1=123456,过期时间为10秒")

	//等待2秒
	fmt.Println("等待2秒...")
	time.Sleep(2*time.Second)

    // 获取 session 的剩余时间
    ttl, err := client.TTL(ctx, "session1").Result()
    if err != nil {
        log.Fatal("获取过期时间失败", err)
    }
    fmt.Printf("session1 剩余时间: %v\n", ttl)

	//删除键
	_,err = client.Del(ctx,"user1").Result()
	if err!=nil {
		log.Fatal("删除键失败",err)
	}
	fmt.Println("删除了user1")

	//再次获取键值对
	name, err = client.Get(ctx, "user1").Result()
    if err == redis.Nil {  //如果不存在会返回redis.Nil
        fmt.Println("user1 不存在")  
    } else if err != nil {
        log.Fatal("获取键值对失败", err)
    } else {
        fmt.Printf("user1的值是: %s\n", name)
    }
}


```
- mind ctx context
- Note that at the end there is.result () or. The characteristics of Err()
- is like redis original grammar


###### Operating complex data structures (Hash and List)
``` GO
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background() //创建上下文

	//操作Hash（存储用户对象）
	user := map[string]interface{}{ //interface{} 表示任意类型，就很方便
		"name": "Bob",
		"age":  30,
	}
	err := client.HSet(ctx, "user2", user).Err()
	if err != nil {
		log.Fatal("Hset失败", err)
	}
	fmt.Println("存储用户user2")

	//获取Hash
	name, err := client.HGet(ctx, "user2", "name").Result()
	if err != nil {
		log.Fatal("Hget失败", err)
	}
	fmt.Printf("获取用户user2的name:%s\n", name)

	//操作List（消息队列）
	queue := "message_queue"
	err = client.RPush(ctx, queue, "msg1", "msg2", "msg3").Err() //Rpush从列表右端（尾部）插入元素，和Lpush相反
	if err != nil {
		log.Fatal("RPush失败", err)
	}
	fmt.Println("消息入队到message_queue")

	//弹出
	msg, err := client.LPop(ctx, queue).Result()
	if err != nil {
		log.Fatal("LPop失败", err)
	}
	fmt.Printf("从message_queue弹出消息：%s\n", msg)
}

```
（一些设置要是输出怪可能是和已经存在的变量名有关，改一下就可以了）

- **Hash**：HSet 和 HGet 适合存储结构化数据（如用户信息）。
- **List**：RPush 和 LPop 实现 FIFO 队列，常用作任务队列。
- 实际项目中，Hash 常用于缓存对象，List 用于异步任务处理


######   Publish/Subscribe (Pub/ Sub)
``` GO
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background() //创建上下文

	//订阅者
	go func(){  //通过协程持续监听
		pubsub := client.Subscribe(ctx,"channel1")
		defer pubsub.Close()

		for msg := range pubsub.Channel(){ //调用channel方法，获取消息
			fmt.Printf("收到消息:%s(频道:%s)\n",msg.Payload,msg.Channel)
		}
	}()

	//发布者
	time.Sleep(1*time.Second) //确保订阅者启动
	err := client.Publish(ctx,"channel1","hello,Redis").Err()
	if err != nil{
		log.Fatal("发布失败",err)
	}
	fmt.Println("发布消息成功")

	//保持运行
	time.Sleep(2 * time.Second)
}

```

###### 分布式锁
通过lua脚本等简单了解分布式锁的概念
``` Go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func acquireLock(client *redis.Client, ctx context.Context, lockKey string, value string, ttl time.Duration) bool {
	// 使用 SETNX 命令尝试获取锁
	result, err := client.SetNX(ctx, lockKey, value, ttl).Result()
	if err != nil {
		log.Println("Failed to acquire lock:", err)
		return false
	}
	return result
	//和SetNX返回值有关，redis中当键不存在时，设置成功，返回 1
	//SetNX：set if not exist
}

func releaseLock(client *redis.Client, ctx context.Context, lockKey, value string) bool {
	// Lua 脚本确保原子性释放
	script := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then  
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end  
	` // KEYS[1] 是锁的键，ARGV[1] 是锁的值，检查锁的值是否匹配，如果匹配则删除锁
	result, err := client.Eval(ctx, script, []string{lockKey}, value).Int() //[]string{lockKey} 是 Lua 脚本的参数，value 是锁的值
	if err != nil {
		return false
	}
	return result == 1
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	lockKey := "mylock"
	value := "unique_123"
	ttl := 20 * time.Second
	if acquireLock(client, ctx, lockKey, value, ttl) {
		fmt.Println("获取锁成功")
		// 模拟业务操作
		time.Sleep(2 * time.Second)
		//释放锁
		if releaseLock(client, ctx, lockKey, value) {
			fmt.Println("释放锁成功")
		} else {
			fmt.Println("释放锁失败")
		}
	}else {
		fmt.Println("获取锁失败")
	}
}

```
- 关键在于做到只能释放自己的锁，防止竞争
- 可以用于库存扣减、分布式任务调度

###### 连接池和性能优化
简单了解
包括通过pipeline进行批量输入
``` GO
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		PoolSize:     10,
		MinIdleConns: 2,
		MaxRetries:   3,
		DialTimeout:  time.Second * 5,    
		ReadTimeout:  time.Second * 3,    
		WriteTimeout: time.Second * 3,    
	})
	ctx := context.Background()

	// 批量写入（使用 Pipeline 提升性能）
	pipe := client.Pipeline()
	for i := 0; i < 100; i++ { //将命令加入到管道中，但不立即执行
		pipe.Set(ctx, fmt.Sprintf("key:%d", i), i, time.Second*60)
		//sprinf形成新的变量比如key:0，key:1，key:2
	}
	_, err := pipe.Exec(ctx) //一次性执行所有命令
	if err != nil {
		log.Fatal("Pipeline 执行失败:", err)
	}
	fmt.Println("批量写入 100 个键完成")

	// 并发读取
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加等待组中的计数
		go func(id int) {
			defer wg.Add(-1)
			val, err := client.Get(ctx, fmt.Sprintf("key:%d", i)).Result()
			if err != nil {
				log.Printf("读取 key:%d 失败: %v", id, err)
				return
			}
			fmt.Printf("读取 key:%d = %s\n", id, val)
		}(i)
	}
	wg.Wait() // 等待所有协程完成
}

```
- **接池**：PoolSize 和 MinIdleConns 控制连接复用，适合高并发。
- **Pipeline**：批量执行命令，减少网络开销。
- **并发**：使用 goroutine 并发操作，结合 sync.WaitGroup 同步。
- **超时配置**：防止网络问题导致阻塞。
###### AOF/RDB
同样当前文件夹下，运行下面的go文件： （go run xxx.go）
```  GO
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 连接 Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	// 测试连接
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	fmt.Println("连接成功")

	// 写入数据（触发 AOF 记录）
	err = client.Set(ctx, "persistent_key", "critical_data", 0).Err()
	if err != nil {
		log.Fatal("设置失败:", err)
	}
	fmt.Println("设置 persistent_key = critical_data")

	// 触发 RDB 快照
	err = client.BgSave(ctx).Err()
	if err != nil {
		log.Fatal("触发 RDB 快照失败:", err)
	}
	fmt.Println("触发 RDB 快照")

	// 触发 AOF 重写（优化 AOF 文件）
	err = client.BgRewriteAOF(ctx).Err()
	if err != nil {
		log.Fatal("触发 AOF 重写失败:", err)
	}
	fmt.Println("触发 AOF 重写")

	// 等待持久化完成
	time.Sleep(2 * time.Second)

	// 检查持久化状态
	info, err := client.Info(ctx, "persistence").Result()
	if err != nil {
		log.Fatal("获取持久化信息失败:", err)
	}
	fmt.Println("持久化状态:\n", info)

	// 验证数据
	value, err := client.Get(ctx, "persistent_key").Result()
	if err != nil {
		log.Fatal("读取失败:", err)
	}
	fmt.Printf("读取 persistent_key = %s\n", value)
}
```
- 运行后可以发现redisstorage中有添加存储数据
- `redis-cli` 进入redis服务 `get persistent_key` 发现可以看到键对应的值结果critical_data
- 然后可以通过Ctrl-c退出运行redis的终端，或者`redis-cli shutdown` 关闭服务
- 接着再次`redis-server ./redis.conf`启动服务，可以发现`get persistent_key`还是可以得到对应的结果，说明成功（如果不配置持久化的话，就是不写配置文件，redis也会有默认配置文件存储在运行redis的文件夹下 ）
###### 哨兵模式
**哨兵模式**（Sentinel）用于 Redis 高可用，监控主从节点，自动故障转移。Go 客户端通过 go-redis 的 FailoverClient 连接哨兵

新建sentinel.conf 文件: 
``` bash
port 26379
sentinel monitor mymaster 127.0.0.1 6379 1
sentinel down-after-milliseconds mymaster 5000
sentinel failover-timeout mymaster 15000
```
- 设置端口为26379 
- 设置监控的主节点，以及只用一个哨兵就可以完成选举（多配置哨兵可以实现高可用）
- 主节点响应超过 5000 毫秒（5秒）就认为主观下线
- 障转移超时时间为 15000 毫秒（15秒），时间内没完成转移认为转移失败

准备好GO文件：
``` GO
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 连接哨兵
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",           // 哨兵监控的主节点名称
		SentinelAddrs: []string{"localhost:26379"}, // 哨兵地址
	})
	ctx := context.Background()

	// 测试连接
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	fmt.Println("哨兵模式连接成功")

	// 写入数据
	err = client.Set(ctx, "sentinel_key", "high_availability", 0).Err()
	if err != nil {
		log.Fatal("设置失败:", err)
	}
	fmt.Println("设置 sentinel_key = high_availability")

	// 读取数据
	value, err := client.Get(ctx, "sentinel_key").Result()
	if err != nil {
		log.Fatal("读取失败:", err)
	}
	fmt.Printf("读取 sentinel_key = %s\n", value)

	// 模拟主节点故障（手动停止主节点）
	fmt.Println("请手动停止主节点（6379），然后再次读取")
	// 等待用户操作
	// 假设主节点故障，哨兵自动切换到从节点
	time.Sleep(10 * time.Second)

	// 再次读取，验证故障转移
	value, err = client.Get(ctx, "sentinel_key").Result()
	if err != nil {
		log.Fatal("故障转移后读取失败:", err)
	}
	fmt.Printf("故障转移后读取 sentinel_key = %s\n", value)
}
```


需要开启多个终端：
配置主节点`redis-server --port 6379` 
配置从节点：`redis-server --port 6380 --replicaof 127.0.0.1 6379`
启动哨兵：`redis-sentinel ./sentinel.conf`
运行GO文件：`go run xxx.go`

接着自己手动关闭主节点的redis服务，等待后可以发现故障转移后可以读取！

- 通过redis-cli验证
``` bash
redis-cli -p 26379
SENTINEL get-master-addr-by-name mymaster  # 查看当前主节点
# 输出：127.0.0.1 6379
# 停止主节点（另一个终端）
redis-cli -p 6379 SHUTDOWN
# 再次检查主节点（哨兵应切换到 6380）
SENTINEL get-master-addr-by-name mymaster
# 输出：127.0.0.1 6380
```

# Advanced features




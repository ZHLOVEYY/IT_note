## 基本理解
概念：
- 类似缓存区
- 用于协程之间通信
基本操作：
 - msgchannel := make(chan, int ) 
- 写入：msgchannel <- i     写出：  j := <- msgchannel

不用关心通道是否关闭，没有被引用的时候会被垃圾回收机制自动处理

## 常见示例
#### 遍历channel
``` Go
package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}
```
#### 通道传递
``` Go
package main
		  import "fmt"
		  func ping(pings chan<- string, msg string) {
		      pings <- msg
		  }
		  func pong(pings <-chan string, pongs chan<- string) {
		      msg := <-pings
		      pongs <- msg
		  }
		  //the direction is set
		  func main() {
		      pings := make(chan string, 1)
		      pongs := make(chan string, 1)
		      ping(pings, "passed message")
		      pong(pings, pongs)
		      fmt.Println(<-pongs)
		      //output: passed message
		  }
```


#### channel循环输入输出
``` GO
package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <=3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		//time.Sleep(time.Millisecond) // 添加短暂延迟
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}

```
可以发现当 j 的循环范围增加的时候，会出现一些问问题
- 当改为 7 时，发现会有先 received 后 sent 的现象，这是因为：主 goroutine 在发送 job 7 后被阻塞，等待 worker 接受了之后才打印出sent！！ 可以取消注释添加打印延迟使得更加真实
#### 通道死锁
无缓冲通道的读取和写入应该位于不同的协程中，不然死锁
``` GO
func main() {
    var c = make(chan int)
	c <- 1
	<-c
}
//有死锁

//改进方案：
//增加了缓冲
func main() {
    c := make(chan int, 1) // 带1个缓冲的通道
    c <- 1
    <-c
    fmt.Println("Hello!")
}

//增加异步处理
func main() {
    c := make(chan int) //没有缓冲需要协程引入
    go func() {
        c <- 1
    }()
    <-c
}

```

#### 通道关闭
利用close关闭
主通道关闭的时候也会收到通知
```GO
func main(){
	var c  = make(chan int)
	go func() {
		data,ok:= <-c
		fmt.Println("goroutine one: ",data,ok)
	}()
	go func() {
		data,ok:= <-c
		fmt.Println("goroutine two: ",data,ok)
	}()
	close(c)
	time.Sleep(1*time.Second)

}
// goroutine two:  0 false
// goroutine one:  0 false
```


#### 通道作为一等公民分配
- goroutine的顺序不固定但是工人分配到的工作是固定的
- 通道可以作为类型放到数组中
``` GO
package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c { // // 这个for循环会一直运行，等待channel中有数据
		fmt.Printf("Worker %d received %d\n", id, n)
	}
    //只有当channel被关闭时，循环才会退出
}
func CreateWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan int // create 10 channels
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i) // create 10 workers
        //此时 channel 中i 对应的位置还是一个空的channel c
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
    // // 发送完所有数据后关闭channel (关闭会更加规范，虽然 main 函数结束也会释放)
    for i := 0; i < 10; i++ {
        close(channels[i])  // 显式关闭每个channel
    }
    //close(channels) //channels 是一个数组（[10]chan int），而不是一个 channel 所以不用 close

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	//'a' 97 开始往后 i 写入  'A'65 开始
}

```

#### 通道访问理解
类似爬虫的逻辑，一个网站访问结束后开启新的协程检查
``` GO
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com/",
		"https://www.taobao.com/",
		"https://www.163.com/",
	}
	var c = make(chan string)
	for _, link := range links {
		go checkLink(link, c) // 并发执行
	}
	<-c //如果只有一个只会返回最先结束的子协程
	<-c
	<-c
	<-c
	// <-c //会卡死

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
	c <- "is up" // 通知主线程
}

```

## select 结合使用

#### select的随机性
``` GO
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	c <- 1
	select {
	//体现了
	case <-c:
		fmt.Println("random 01")
	case <-c:
		fmt.Println("random 02")
	}

}

//01 and 02

```
- 可以使用default来避免阻塞 ： 
case <-time.After(800 * time.Millisecond)
- 当channel为nil时，它将永远不会被选中

循环 select 输出：
``` Go
func main() {
	c := make(chan int,1)
	tick := time.Tick(time.Second) //创建了一个定时器通道，每隔一秒向通道发送一个当前时间
	for {
		select{
		case <-c:
			fmt.Println("random 01")
		case <-tick:
			fmt.Println("tick 01")
		case <-time.After(800*time.Millisecond): //设置1s就会只有tick
		//这个就是交替的
			fmt.Println("timeout 01")
		}
	}	
}
```


## 其他功能函数加入

#### TImer
在一定时间后将值发送到channel中  后面记得带上.C
``` GO
package main
  
import (
    "fmt"
    "time"
)

func main() {

    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C // 阻塞等待，直到2秒后定时器触发
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C // 在goroutine中等待定时器
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop() // 立即尝试停止定时器
    if stop2 {
        fmt.Println("Timer 2 stopped")// 停止成功则打印
    }

    time.Sleep(2 * time.Second)
}
```
 Stop（）函数阻止计时器触发，如果成功停止计时器则返回true


#### Ticker
帮助你定期重复做某事
``` Go
package main
	  import (
	      "fmt"
	      "time"
	  )
	  func main() {
	      ticker := time.NewTicker(500 * time.Millisecond)
	      done := make(chan bool)
	      go func() {
	          for {
	              select {
	              case <-done:
	                  return
	              case t := <-ticker.C:
	              //will send repeatedly
	                  fmt.Println("Tick at", t)
	              }
	          }
	      }()
	      time.Sleep(1600 * time.Millisecond)
	      ticker.Stop()
	      done <- true
	      fmt.Println("Ticker stopped")
	  }
```

#### 原子钟  Atomic Counter 
可以看： https://gobyexample.com/atomic-counters 

## 进阶示例

#### 工人池
```
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// to fill in the job
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	// to ensure the certain amount of the job is done
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

```
使用相同的通道来精确任务

#### 速率限流
``` GO
package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// like a array
	limiter := time.Tick(200 * time.Millisecond)
	//Ticker,every 200ms receive a value (for channel)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	// print the formal rate

	burstyLimiter := make(chan time.Time, 3)
	// a bursty handler
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	//to fill in the limiter(3) after the initual output
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		// has three value at first, so handle at the same time
		fmt.Println("request", req, time.Now())
	}
}

```
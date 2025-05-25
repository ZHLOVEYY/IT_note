
开启进程的消耗比线程大很多。进程享有独立空间，进程之间的通信比较复杂
  
## 协程基本理解
如果是下面这样线形访问，需要一个个访问请求，较慢：
``` GO
func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com/",
		"https://www.taobao.com/",
		"https://www.163.com/",
		"https://www.sohu.com/",
	}
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

```
于是在 main 函数中引入协程，但是由于主协程结束子协程也会结束，所以需要预留运行时间

``` GO
func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com/",
		"https://www.taobao.com/",
		"https://www.163.com/",
		"https://www.sohu.com/",
	}
	for _, link := range links {
		go checkLink(link) // 并发执行
	}
	time.Sleep(1 * time.Second)  //需要一定的时间，1/2就不太行
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
//每次打印的结果都不太一样的
```

## GMP模型
G：协程   M：线程 P：Go的逻辑处理器
- 一个P可以包含多个G协程（多汇聚为1），然后一个P任意时刻只能绑定一个M线程
- 不一定一直绑定同一个P处理器，同时P处理器也不是总对应同一个M线程





WaitGroup
``` Go
func main() {
	//主协程（main 函数）在启动子协程后立即退出，会强制关闭
	var wg sync.WaitGroup
	values := []string{"apple", "banana", "cherry"}
	for _, val := range values {
		wg.Add(1) // 增加计数
		go func(v string) {
			defer wg.Done() //单独的foroutine完成后减少计数
			fmt.Println(v)
		}(val)
	}
	wg.Wait() // 阻塞直到所有子协程完成
}
```



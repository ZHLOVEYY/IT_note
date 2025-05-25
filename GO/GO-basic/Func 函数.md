函数是一等公民！！ 可以作为返回值，用作变量或者赋值等
###### 函数栈概念
函数调用的时候消耗栈空间，因为调用完函数需要返回原来地方继续执行   栈说明是后进先出的方式
同时每个协程都有一个栈，并且初始化大概2KB，可以扩容
#### 函数表示类型
``` GO
	func visit (numbers []int,callback func(int)) {
		//func(int) represents it receives an interger but return nothing
		for _,num := range numbers {
			callback(num)
		}
	}
```
callback 就是一个传入 int 的函数

## Anonymous func 匿名函数
示例： 就是看立即调用还是手动调用
``` Go
package main

import "fmt"

func main() {
    func(){
        fmt.Println("Hello, world!")
    }()   //立刻调用
    myfunc := func() {
        fmt.Println("Hello, myfunc!")
    }
    myfunc()    //手动调用
    dosomething(func(){
        fmt.Println("Hello, dosomething!")
    })
}

func dosomething(f func()) {
    f()  //内部手动调用
}


```

## closure 闭包
闭包可以修改之中的局部变量，在调用的时候同时故事保持之中的局部变量不变
#### 基础示例
``` GO
package main

import "fmt"

func main(){
    nextInt:= intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    // 1 2 3
    newInts := intSeq()
    fmt.Println(newInts())
    //1

    //虽然 addr 不接受参数，但是它仍然可以访问外部变量 sum。
    pos,neg := adder(),adder()  //这里已经调用了 adder 函数
	for i :=0 ; i< 5; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}


}

//intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。返回的函数使用闭包的方式 隐藏 变量 i。
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
    // the func capture the "i"
}

//sum 
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
  }
```
（回调函数，函数式编程中常用）

#### http利用闭包裹挟中间件
设计到 gin 等 web 框架的概念
``` GO
package main

import (
	"fmt"
	"net/http"
	"time"
)

func timed(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r) //执行原本的hello函数
		end := time.Now()
		fmt.Println("Time taken:", end.Sub(start)) //打印在控制台
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", timed(hello)) //访问http://localhost:8080/hello
	http.ListenAndServe(":8080", nil)
}

```
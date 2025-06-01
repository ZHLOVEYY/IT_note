## 基础知识
defer一般用于资源释放和异常捕获
defer有三种：堆上，栈上以及内联  内联相当于代码放到函数最后执行
[defer必须掌握7ge](https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmee0q64ij5p)

## 基础案例
#### 多个defer倒着执行的顺序
``` GO
func main() {
	for i:= 1;i<5;i++{
		defer fmt.Println("start",i)
	}
	//start 4
	//start 3
	//。。。
}
```
从后面的开始往前

#### defer返回数值问题


## defer 的作用

#### 资源释放
``` GO
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	// written, err = io.Copy(dst, src) //第一个参数是写入部分,这个直接return
	// dst.Close() //如果没有defer中间有问题，就不会执行close了
	// src.Close()
	return io.Copy(dst, src)
}
```
defer close 类的

#### 异常捕获
用于打印：
``` GO
func excute() {
	defer fmt.Println("defer func")
	panic("this is a panic")
	fmt.Println("completed")
}
func main() {
	excute()
	fmt.Println("main func") //main函数也将不能运行
}
```

或者recover 接收 panic：
``` GO
func excute() {
	defer func() {
		if errMsg := recover(); errMsg!= nil {  //用recover接收
			fmt.Println(errMsg)
		fmt.Println("this is a recovery function")}
	}()
	panic("this is an error")
	fmt.Println("completed")
}
func main(){
	excute()
	fmt.Println("main function completed")  //可以正常执行
}
```

#### 中间件的一中实现方式
``` GO
package main

import (
	"fmt"
	"log"
	"net/http"
)

func recoverHandler(next http.Handler) http.Handler { //需要配合handle的类型
	fn := func(w http.ResponseWriter, r *http.Request) { //签名与 http.HandlerFunc 的定义完全一致：
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic:%v", err)
				http.Error(w, http.StatusText(500), 500)
			} else{
				next.ServeHTTP(w, r) //调用对应的ServeHTTP方法，再打印helloworld
            }
		}()
		next.ServeHTTP(w, r) //调用对应的ServeHTTP方法，这里是打印helloworld
	}
	return http.HandlerFunc(fn) //转换为 http.Handler 类型.fn 是一个符合 http.HandlerFunc 签名的函数就可以
	//由于HandlerFunc实现了ServeHTTP方法，所以 http.HandlerFunc(fn) 就是一个实现了 http.Handler 接口的实例！！！

}

type MyHandler struct{}

// 定义MyHandler结构体的ServerHTTP方法，接收一个http.ResponseWriter和一个http.Request作为参数
func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 将字符串"hello world"写入http.ResponseWriter
	fmt.Fprintf(w, "hello world")
}

// 主函数
func main() {
	// 创建一个MyHandler实例
	handler := MyHandler{}
	// 创建一个http.Server实例，指定地址和处理器
	server := http.Server{
		Addr:    ":8080",
		Handler: recoverHandler(handler),
	}
	// 启动服务器
	server.ListenAndServe()
}

```

#### 参数预计算
会在执行到的时候计算
``` GO
func main() {
	a := 1
	defer func(b int) {
		//首先预编译执行
		fmt.Println("defer b:", b)
	}(a + 1)
	a = 99
}
```

#### defer返回数值问题
- 结合全局变量和局部变量的存放问题理解，全局存在堆上，局部存在函数栈上
- 先return？先defer？ ->先 return 再 defer
看下面两个例子对比：
``` Go
var g = 100 //全局变量
func f() (r int){
	defer func(){
		g = 200  //对全局变量进行赋值
	} ()
	fmt.Printf("f:g = %d\n",g)  //\n换行符
	//f:g = 100
	return g
}

func main() {
	i:=f()
	fmt.Printf("main:i = %d,g = %d\n",i,g)
	//main:i = 100,g = 200
}
//流程
//g = 100
//r = g=100 (r就是return的)
//g = 200
// return 得到r
```
以及
``` GO
//另外一种可能：
var g = 100 //全局变量
func f() (r int){
	r = g
	defer func(){
		r = 200
	}()
	r = 0
	return r
}

func main() {
	i := f()
	fmt.Printf("main:i = %d, g = %d\n",i,g)
	//main:i = 200, g = 100  
	//上一个例子中就是i是100 
}
//流程：
// g = 100 //全局变量
// r = g = 100
// r = 0
// r =200
// r = r返回
```

本质是：返回值return（r）不是全局变量，所以存在栈上，执行defer函数后返回 （所以如果修改r就是可以体现的）
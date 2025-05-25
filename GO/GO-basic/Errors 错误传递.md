## 自己创建错误
error 信息可以自己定义
``` GO
func main() {
	arg := 2
    arg2 := 4
	err := makeTea(arg)
	if err != nil {
		fmt.Println(err)
	}
	err2 := makeTea(arg2)
	if err2!= nil {
		fmt.Println(err2)
	}
}

//两种创建错误的方式
func makeTea(arg int) error {
	if arg == 2 {
		return fmt.Errorf("no more tea available")
	} else if arg == 4 {
		return errors.New("no more power available")
	}
	return nil
}
```

## panic和recover
#### 基础示范
panic 会报错然后程序返回，recover 结合defer 放在最后匿名函数中执行用于接受 panic
``` Go
package main

import (
	"fmt"
)

func mayPanic() {
	panic("a problem") //手动 panic
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("成功接收Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()") // won't output
}
//output：
//成功接收Recovered. Error:
//a problem

```

#### panic的返回
通过嵌套函数理解返回方式
``` Go
func a(){
	defer fmt.Println("defer a")
	b()
	fmt.Println("after a")
}
func b(){
	defer fmt.Println("defer b")
	c()
	fmt.Println("after b")
}
func c(){
	defer fmt.Println("defer c")
	panic("this is a panic")
	fmt.Println("after c")
}

func main() {
	a()
}
// defer c
// defer b
// defer a
// panic: this is a panic
```

修改b部分：可以保证回到a
``` GO
func b() {
	defer func() { //变为立即执行的函数
		fmt.Println("defer b")
		if x := recover(); x != nil {
			fmt.Println("run time panic: %v \n", x)
		}
	}()
	c()
	fmt.Println("after b")
}
//defer c
// defer b
// run time panic: %v 
//  this is a panic
// after a
// defer a

```
能回到 A 就会先打印出 after a 然后再结束就打印出 defer a

###### 用于保证g（）函数执行的中间件
``` GO
func protect(g func()) {
	defer func(){
		log.Println("done")
		if x:=recover();x!=nil{
			log.Printf("run time panic:%v",x)
		}
	}()
	log.Println("start")
	g()   //执行g   如果有问题会返回上去的
}
```


#### panic和recover的嵌套顺序
###### 打印最先 panic
如果在defer调用的函数中也有panic，会打印最早的panic （没有recover的时候
``` GO
func main() {
	a()
}

func a() {
	defer b()
	panic("a panic")
}

func b() {
	defer fb()
	panic("b panic")
}

func fb() {
	panic("fb panic")
}
// panic: a panic
// panic: b panic
// panic: fb panic

```
这个例子中，先 panic a 然后调用 b，继续向下迭代


###### recover 捕获最近 panic
``` GO
func main() {
	defer catch("main")  //增加捕获
	a()
}

func catch(name string) {
	if r := recover(); r != nil {
		fmt.Println(name, "recover:", r)
		//println会打印地址
	}
}

func a() {
	defer b()
	panic("a panic")
}

func b() {
	defer fb()
	panic("b panic")
}

func fb() {
	defer catch("fb")    // //增加捕获
	panic("fb panic")
}

// fb recover: fb panic
// main recover: b panic
```
执行顺序： a 中 panic ->defer 调用 b ->b 中 panic->调用 fb->fb
中 panic->defer 调用 catch，抓住的是 fb panic ->回到 main 中 defer->抓住的是b panic （因为 fb 的 panic 已经被捕获，会向前推）



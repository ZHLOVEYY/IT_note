- 求长度还是utf库中求，因为中文有占3/4字节的，用len不准确
- 符串和数组不能直接123 + “456”这样拼接
	``` GO
		con := fmt.Sprintf("123%d",456)
	  	println(con)
	```

## 区分 rune，byte，string
``` GO
package main

import (
	"fmt"
)

func main() {
	s := "hello world" //直接字符串形式
	for _, v := range s {
		fmt.Printf("value is %c,type is %T\n", v, v)
		//typeis：int32value (rune)
	}

	for _, v := range s {
		fmt.Println(v)  //是 int32 形式的数字
		//is int32vlaue (rune)
		fmt.Println(string(v))
		//string
	}

	tokens := []string{"2", "1", "+", "3", "*"} //用 string处理
	for i := 0; i < len(tokens); i++ {
		fmt.Printf("type is %T\n", tokens[i])
		//type is string
	}

	bytes := []byte("hello world")  //用 byte 处理
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("value is %c,type is %T", bytes[i], bytes[i])
		//type is uint8
		fmt.Println(bytes[i])
		//value is 104,type is uint8
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("value is %c,type is %T\n", runes[i], runes[i])
		//value is h,type is rune
		fmt.Println(string(runes[i]))
		//string
	}
}

```
总结：
- 原始字符串之中 的内容每一个是 rune 存储，rune 是 int32
- byte 就是转为 int8
- string类型是不同于 rune 的，stirng 的列表中的类型就是 string 而不是 rune 的 int32

## 字符串操作
#### strings 库相关
常用：
``` GO
func main() {
    fmt.Println(strings.ToUpper("hello"))
    fmt.Println(strings.ToLower("HELLO"))
    fmt.Println(strings.Replace("hello world", "world", "go", -1)) //hello go
    fmt.Println(strings.Split("hel-lo-w-rld", "-"))  //[hel lo w rld]
    fmt.Println(strings.Join([]string{"hel", "lo", "w", "rld"}, "-")) //hel-lo-w-rld
}
```

## fmt.Printf语法
```GO
//输出形式
fmt.Printf("整数：%d\n", 123)      // %d：十进制整数
fmt.Printf("字符：%c\n", 'A')      // %c：字符
fmt.Printf("字符串：%s\n", "hello") // %s：字符串
fmt.Printf("布尔值：%t\n", true)    // %t：布尔值
fmt.Printf("浮点数：%f\n", 3.14)    // %f：浮点数
fmt.Printf("二进制：%b\n", 15)      // %b：二进制
fmt.Printf("十六进制：%x\n", 15)    // %x：十六进制（小写） 

//类型相关
fmt.Printf("类型：%T\n", 123)       // %T：类型
fmt.Printf("值：%v\n", 123)         // %v：默认格式
fmt.Printf("Go语法：%#v\n", "hello") // %#v：Go语法格式
fmt.Printf("p=%+v\n", p) // p={x:1 y:2} 打印结构体字段和名
fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2} 打印更具体结构体名称

```
Println 会自动回车，Printf 需要\n


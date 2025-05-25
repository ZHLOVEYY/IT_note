Json化后方便网络通信等传输，主要特点其实就是对第一个字母进行处理

## 基础案例
``` GO
type userInfo struct {
	Name  string   `json:"Name"`
	Age   int      `json:"age"`
	Hobby []string `json:"Hobby"`
}

func main() {
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}

	// 普通JSON编码
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

	// 带缩进的JSON编码
	indentBuf, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indentBuf)) //如果不是 string 会以数字形式输出
	/*
		{
			"Name": "wang",
			"age": 18,
			"Hobby": [
				"Golang",
				"TypeScript"
			]
		}
	*/

	// JSON解码
	var b userInfo
	err = json.Unmarshal(buf, &b) //解码然后放入 b 中
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}
```

 - 似乎在定义结构体的 json 的时候出一些错误都是可以被 Unmarshal 正确重新放回去的（不过 Marshal 的时候就会带有一点错误地放入）
 还是老实定义 json 就好
 - 记住需要编为 json 后，string 来输出，不然就是一串数字
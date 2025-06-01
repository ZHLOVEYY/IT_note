更多个人笔记：（仅供参考，非盈利）
gitee： https://gitee.com/harryhack/it_note
github： https://github.com/ZHLOVEYY/IT_note
（里面也有我的git教学）

针对GO中net/http包的学习笔记
### 基础快速了解

创建简单的GOHTTP服务
``` GO
func main() {  
    http.HandleFunc("/hello", sayHello)  
    http.ListenAndServe(":8080", nil) //创建基本服务  
}  
  
func sayHello(w http.ResponseWriter, r *http.Request) {  
    w.Write([]byte("Hello, World!"))  
}
```
访问8080/hello进行测试

Handler接口定义：（这部分后面又详细解释）
``` GO
type Handler interface {  
    ServeHTTP(ResponseWriter, *Request)  
}
//只要有ServeHTTP方法就行
```
可以自己实现这个接口

同时http提供了handlerFunc结构体
``` GO
type HandlerFunc func(ResponseWriter, *Request)
// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
f(w, r)
}
//本质上就是调用自身，因为也是一个函数，不过serveHTTP的内容自己可以定义改动
```
和之前的HandleFunc区分，HandleFunc是用来给不同路径绑定方法的，少一个r

那么只要满足Handlerfunc 的签名形式，就可以进行类型转换

- 类型转换的正常的理解例子
``` GO
type yes func(int, int) int  
func (y yes) add(a, b int) int {  
    return a + b  
}  
func multiply(a, b int) int {  
	fmt.Println(a * b)
    return a * b  
}  
func main() {  
    multiply(1, 2)   //2
    ans := yes(multiply) //将multiply进行转换  
    res := ans.add(1, 2)  
    fmt.Println(res)   //3
}

```


``` GO
http.HandleFunc("/hello", hello)
```
这个后面的签名只要是 `func(ResponseWriter, *Request)`就可以了
但是
``` GO
http.ListenAndServe(":8080", referer)
```
这个后面的函数需要是满足Handler接口，有serveHTTP方法


尝试搭建检测是在query中有name = red
即http://localhost:8080/hello?name=red   
发现会有重复覆盖路由的问题，因为listenandServe会拦截所有的路由，后面再解决
``` GO

type CheckQueryName struct {
	wantname string
	handler  http.Handler
}

func (this *CheckQueryName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query() //获取get请求的query
	name := queryParams.Get("name")
	if name == "red" {
		this.handler.ServeHTTP(w, r) //其实就是调用本身，下面变为checkforname了
	} else {
		w.Write([]byte("not this name"))
	}
}

func checkforname(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check is ok"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	thecheck := &CheckQueryName{ //用&因为serveHTTP方法定义在指针接收器上
		wantname: "red",
		handler:  http.HandlerFunc(checkforname),
	}
	http.HandleFunc("/hello", hello)       //满足func(ResponseWriter, *Request)签名就可以
	http.ListenAndServe(":8080", thecheck) //直接监视8080的所有端口，拦截所有路由
}
```

#### 编写简单的GET请求客户端
利用defaultclient或者自己定义client都可以
``` Go
func main() {
	resp, err := http.DefaultClient.Get("https://api.github.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	time.Sleep(time.Second * 2)  //等待获取请求
}
```
但是如果把网址换成baidu.com就会获取不到，这是因为转发，以及没有User-agent的问题

#### 编写自定义的GET请求客户端
利用http.Client可以进行自定义
``` GO
func main() {
	client := &http.Client{
		// 允许重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}
	// 添加请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req) 
	//do执行HTTP请求的整个周期包括请求准备，建立连接，发送请求，请求重定向，接收响应等等
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	time.Sleep(time.Second * 2)
}
```
发现可以接受到baidu的网页html信息

#### 编写默认的post请求客户端
``` GO
func main() {
	postData := strings.NewReader(`{"name": "张三", "age": 25}`)
	resp, err := http.DefaultClient.Post(
		"http://localhost:8080/users",
		"application/json",  
		postData,
	)
	if err != nil {
		fmt.Printf("POST请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("POST响应: %s\n", string(body))
}
```
string.NewReader是一种方法将格式改为io reader可以读取的形式，接收的话也可以postData.Read读取。
这一类的方式比较多，不一一汇总

对应的server.go  （需要在终端中go run server.go ）两个终端分别运行服务端，客户端
``` Go
func main() {
	// 处理 /users 路径的 POST 请求
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// 只允许 POST 方法
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "只支持 POST 方法")
			return
		}
		// 读取请求体
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "读取请求失败: %v", err)
			return
		}
		defer r.Body.Close()
		// 解析 JSON 数据，放入user中
		var user User
		if err := json.Unmarshal(body, &user); err != nil {
			w.WriteHeader(http.StatusBadRequest) //写入状态码
			fmt.Fprintf(w, "JSON 解析失败: %v", err)
			return
		}
		// 设置响应头
		w.Header().Set("Content-Type", "application/json")
		// 构造响应数据
		response := map[string]interface{}{
			"message": "success",
			"data": map[string]interface{}{
				"name": user.Name,
				"age":  user.Age,
			},
		}

		// 返回 JSON 响应
		json.NewEncoder(w).Encode(response) //将 response 对象转换为 JSON 格式并写入响应
		//等价于：
		// jsonData, err := json.Marshal(response)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		// w.Write(jsonData)
	})

	// 启动服务器
	fmt.Println("服务器启动在 :8080 端口...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

### 多路复用器
DefaultServeMux一般不会使用，因为会有冲突等等问题，所以一般用NewServeMux直接创建


``` GO
type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "API response"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})  //多引入结构体，后面会知道有好处
	// mux.HandleFunc("/api/", func(w http.ResponseWriter, req *http.Request) {
    //     w.Header().Set("Content-Type", "application/json")
    //     fmt.Fprintf(w, `{"message": "API response from HandleFunc"}`)
    // })   //和上面等效
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
	//用server:= &http.Server创建地址和handler，然后server.ListenAndServe也是一种表现方式
	
}
```
mux和http.HandleFunc()的区别：
mux 可以创建多个路由器实例，用于不同的目的。同时可以为不同的路由器配置不同的中间件（用着先）


第三方有库httprouter，比如可以解决url不能是变量代表的问题，了解就行
更多都是使用restful API进行开发的～目前的了解有个概念就行

### 处理器函数

#### Handle
注册处理器过程中调用的函数：Handle
``` GO
type username struct {
	name string
}

func (this *username) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", this.name)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/jack", &username{name: "jack"}) //会调用对应的serveHTTP方法
	mux.Handle("/lily", &username{name: "lily"})
	//可以为不同的路径使用相同的处理器结构体，但传入不同的参数
	//这就是比用handleFunc()更灵活的地方
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil { //防止错误
		panic(err)
	}
}
```

#### HandlleFunc
处理器函数：HandleFunc
（注意不是HandlerFunc，没有r ！！ HandleFunc 是处理器函数）
之前已经学习过，定义就是
`func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`

深入源代码会发现内部也是借助serveMux对象，从而实现了Handler的ServeHTTP()方法的

#### Handler
Handler就是处理器接口，实现ServeHTTP方法的，之前展示过
``` GO
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

#### HandlerFunc
HandlerFunc是结构体，用于实现接口的
定义：
``` GO
type HandlerFunc func(ResponseWriter, *Request)
```
用于连接处理器Handle和处理器函数HandleFunc，它实现了 Handler 接口，使得函数可以直接当作处理器使用：

- 理解“连接连接处理器Handle和处理器函数HandleFunc”：（之前也学过）
``` GO
// 方式一：普通函数
func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}
// 注册方式一：将函数转换为 HandlerFunc
http.Handle("/hello", http.HandlerFunc(hello))

// 方式二：直接使用 HandleFunc
http.HandleFunc("/hello", hello)
```

### 处理请求
请求分为：请求头，请求URL，请求体等

##### html表单的enctype属性
在postman的body部分可以查看
- application/x-www-form-urlencode
	url方式编码，较为通用，get和post都可以用
- multipart/form-data
	通常适配post方法提交
- text/plain
	适合传递大量数据

### ResponseWriter接口涉及方法
- 补充：
fmt.Fprint和fmt.Fprintln能够写入ResponseWriter是因为ResponseWriter实现了`io.Writer`接口，fmt.Fprint/Fprintln将数据按格式转换为字节流（如字符串、数字等），最终调用io.Writer的Write方法
##### Writeheader
`curl -i localhost:8080/noAuth`  或者使用postman进行验证
``` GO
func noAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	fmt.Fprint(w, "没有授权，你需要认证后访问")
}

func main() {
	http.HandleFunc("/noAuth", noAuth)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

```

##### Header
调用了Writeheader后的话就不能对响应头进行修改了
`curl -i http://localhost:8081/redirect` （可以直接看到301）或者postman验证
- 重定向代码
``` GO
func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://localhost:8080/hello")
	// 必须使用包括http的完整的URL！
	w.WriteHeader(301)
}
func main() {
	http.HandleFunc("/redirect", Redirect)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
```
- 主服务代码
``` GO
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello !!"))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8080", nil) //创建基本服务
}

```

##### write
之前都有demo，就是写入返回，注意需要是\[]byte()这样的表示形式
如果不知道content-type格式可以通过数据的前512 比特进行确认

除了一般的文本字符串之外，还可以返回html和json，下面给出json的示范
``` Go
type language struct {
	Language string `json:"language"` //反引号
	 // 字段名首字母需要大写才能被 JSON 序列化！！！！
}

func uselanguage(w http.ResponseWriter, r *http.Request) {
	uselanguageis := language{Language: "en"}
	message, err := json.Marshal(uselanguageis)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json") //通过json形式传递
	w.Write(message)
}

func main() {
	http.HandleFunc("/lan", uselanguage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```
注意一些格式上的细节，比如字段名首字母需要大写才能被 JSON 序列化






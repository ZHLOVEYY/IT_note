电商系统迅速发展后，有很多服务需要拆分，比如用户服务，商品服务，支付服务，订单服务，售后服务等  这些服务之间的相互调用就需要使用rpc

rpc框架：
![[../../../attachments/Pasted image 20250427144445.png]]

服务注册中心：负责本地服务发布为远程服务，管理提供给消费者
消费者：通过远程代理对象调用远程服务
服务提供者：提供接口和相关服务

rpc框架可以跨语言，协议私密，传输效率也很高

## GO的rpc应用
官方的net/rpc包对于方法有定义要求符合：
``` GO
func (t *T) MethodName (argType T1, replyType *T2) error
```
其中T，T1，T2需要被encoding/gob包序列化（就是说不同编码器都需要适用）
第一个参数是调用者需要传递的参数
第二个是会以字符串方式返回的结果

### 简单编写

服务端：  go run server.go运行就行
``` Go
package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct { //定义传入参数结构
	X, Y int
}

type Algorithm int //就是一个辅助的定义

func (t *Algorithm) Sum(args *Args, reply *int) error { //定义方法
	*reply = args.X + args.Y
	fmt.Println("Exec Sum ", reply)
	return nil
}

func main() {
	algorithm := new(Algorithm) //分配内存空间指针同时分配内存初始化为0值
	fmt.Println("algorithm start", algorithm)
	//注册服务
	rpc.Register(algorithm)
	rpc.HandleHTTP() //将 RPC 服务挂载到 HTTP 服务器
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("err=====", err.Error())
	}
}


```
- 定义传入的结构体参数
- 为了调用方法，定义相关结构体承载
- 实例化对应的承载结构体
- 注册服务，并注册到HTTP上
- 启动端口监听


客户端： 终端开启运行 `go run xxx.go 1 2`  (后面的1和2是制定的参数传入)
``` go
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type Args struct { //定义传入参数结构
	X, Y int
}

type Algorithm int  //就是一个辅助的定义

func (t *Algorithm) Sum(args *Args, reply *int) error {
	*reply = args.X + args.Y
	fmt.Println("Exec Sum ", reply)
	return nil
}


func main() {
	//连接服务器
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dailHTTP发生错误", err)
	}
	//获取第一个输入
	i1, _ := strconv.Atoi(os.Args[1])
	//获取第二个输入
	i2, _ := strconv.Atoi(os.Args[2])
	args := Args{i1, i2}
	var reply int
	//调用远程方法
	err = client.Call("Algorithm.Sum", args, &reply) //调用对应的方法
	if err!= nil {
		log.Fatal("调用远程方法发生错误", err)
	}
	fmt.Printf("Arith: %d+%d=%d\n", args.X, args.Y, reply)
}

```
- 连接对应的服务
- 输入，组成结构体
- 通过Call 调用对应的方法并进行传递

## json编写rpc
- 标准 RPC：主要用于 Go 程序之间通信，用Gob进行二进制编码
- JSON-RPC：适合跨语言通信（因为 JSON 是通用格式），用json进行编码
jsonrpc更加通用，，二者的实际使用其实就是启动方式时函数调用上的区别，先不重复展开

# 微服务
特点：
- 单一职责：一个服务用来解决一个业务问题
- 面向服务：一个服务封装并对外提供服务，也可以调用别的服务

微服务治理
微服务如何发现别的微服务：服务注册中心
客户端或外部服务调用的处理：通过统一的网关进行验证授权
此外还有熔断限流保证高可用，负载均衡，分布式事务等等方法概念

## grpc
grpc是什么：跨平台高性能的rpc框架，多语言互通，可以比如GO创建一个服务端然后PHP/android服务端调用


proto文件： （programmer.proto）

``` Go
syntax = "proto3";
package proto; //通过 package 区分不同的命名空间
//proto.ProgramRequest和other.ProgramRequest是不同的命名空间下的相同消息名
option go_package = "./protooo"; //指定 Go 包路径（生成对应文件夹）
//比如 ./protoo;server 就是创建 protoo 文件夹，里面 package 是 server

service Program{
    rpc Getinfo(ProgramRequest) returns(ProgramResponse){} //定义服务端处理函数
}

message ProgramRequest{
    string name = 1; //[修饰符]类型 字段名=标识号
}

message ProgramResponse{  //定义服务端响应数据格式
    int32 uid = 1;
    string username = 2;
    string job = 3;  
    repeated string hobbies = 4; //repeated是修饰符，表示为可变数组
}
```
当前文件夹终端下执行`protoc --go_out=. --go-grpc_out=. ./programmer.proto`
前面的会根据option 的包路径放置，最后的是指定proto文件
会生成对应的两个pb.go 文件    （知道和protobuf有关就行，脚手架）

server部分代码 （`go run server.go`）
``` GO
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "practice/proto" //这里practice是我自己的go mod的名字

	"google.golang.org/grpc"
)

type ProgramServer struct {
	pb.UnimplementedProgramServer //向前兼容性保护，如果添加了新的方法
	//这里是嵌入的结构体
}

func (s *ProgramServer) Getinfo(ctx context.Context, req *pb.ProgramRequest) (*pb.ProgramResponse, error) {
	// 模拟业务逻辑
	if req.Name == "张三" {
		return &pb.ProgramResponse{
			Uid:      1001,
			Username: req.Name,
			Job:      "软件工程师",
			Hobbies:  []string{"编程", "读书", "运动"},
		}, nil
	}else{
		return &pb.ProgramResponse{
			Uid:      -1,
			Username: req.Name,
			Job:      "嘿嘿嘿，不知道",
			Hobbies:  []string{"没有东西"},
		}, nil
	}
	
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("监听失败:", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterProgramServer(s, &ProgramServer{})
	fmt.Println("gRPC 服务器启动在 :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal("服务失败: ", err)
	}
}


```
- 导入（pb）生成好的代码中的对应的结构体和返回变量名
- 写方法
- 监听窗口并创建RPC服务器，注册服务

客户端代码 (go run client.go)
``` GO
package main

import (
	"context"
	"log"
	"time"

	pb "practice/proto" //这里practice是我自己的go mod的名字

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接服务器
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewProgramClient(conn)

	// 设置超时上下文，context 是必需的，用于调控grpc的生命周期
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用远程方法
	resp, err := client.Getinfo(ctx, &pb.ProgramRequest{Name: "张三"})
	if err != nil {
		log.Fatal("调用失败", err)
	}

	log.Printf("响应: %+v", resp)
}

```
- 也是导入pb包
- 连接服务器并创建客户端
- 设置上下文传递
- 调用远程方法











## 简单知识了解
Socket位于应用层和传输层之间，属于传输的媒介
tcp三次握手和四次挥手是基于socket实现的
三次： A：syn（连接么） B：syn+ack（连，你也发个连） A：ack（连）
四次：A：fin（断吧）  B：ack （ok） B：fin（断吧） A：ack（ok） 

Go中用net包操作socket
- 客户端net.Dial无论什么形式都能连接
- 服务器端使用net.listen (进一步net.accept进行监听)

Dial支持TCP，UDP，ICMP，等


## 实现一个TCP 服务器与客户端（聊天室）

分为server和client  可以一个serer对多个client

server代码：
``` Go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("监听失败:", err)
	}
	defer listener.Close()
	fmt.Println("TCP 服务器启动，监听 :8080")

	// 客户端连接管理
	var clients []net.Conn
	var mutex sync.Mutex

	// 添加服务器发送消息的 goroutine
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				log.Println("读取控制台输入失败:", err)
				continue
			}

			// 广播服务器消息
			broadcast := fmt.Sprintf("[Server]: %s", msg)
			mutex.Lock()
			for _, c := range clients {
				c.Write([]byte(broadcast))
			}
			mutex.Unlock()
		}
	}()

	// 继续处理客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("接受连接失败:", err)
			continue
		}
		fmt.Printf("新客户端连接: %s\n", conn.RemoteAddr().String())

		// 添加到客户端列表
		mutex.Lock()
		clients = append(clients, conn)
		mutex.Unlock()

		// 启动 goroutine 处理客户端
		go handleClient(conn, &clients, &mutex)
	}

}
func handleClient(conn net.Conn, clients *[]net.Conn, mutex *sync.Mutex) {
	defer conn.Close()

	// 读取客户端消息
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("客户端 %s 断开: %v\n", conn.RemoteAddr().String(), err)
			// 从客户端列表移除
			mutex.Lock()
			for i, c := range *clients {
				if c == conn {
					*clients = append((*clients)[:i], (*clients)[i+1:]...)
					break
				}
			}
			mutex.Unlock()
			return // 退出这个协程
		}

		// 广播消息
		broadcast := fmt.Sprintf("[%s]: %s", conn.RemoteAddr().String(), msg) //这里是整合起来
		fmt.Print(broadcast)
		mutex.Lock()
		for _, c := range *clients {
			if c != conn {
				c.Write([]byte(broadcast))
			}
		}
		mutex.Unlock()
	}

}

```
这一类包名一般都是记不住的，需要掌握构建逻辑：
- 建立listen
- 因为可能有多个client连接，所以需要list，所以为了解决并发问题所以才引入的lock
- 服务器发送消息一个协程  （广播 ）
- listen accept建立conn，每个conn对应客户端
- conn中处理读取客户端的信息，通过bufio
- 如果客户端断开，通过msg能不能读取的到进行判断



client代码：
``` GO
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()
	fmt.Println("已连接到服务器")

	// 启动 goroutine 读取服务器消息
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				log.Println("服务器断开:", err)
				return
			}
			fmt.Print(msg)
		}
	}()

	// 从标准输入读取并发送消息
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text() + "\n"
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("发送失败:", err)
			return
		}
	}
}
```
client部分思路：
- 通过dial进行连接，生成conn
- 通过bufio读取conn，从服务器传递的信息
- 设置发送给服务器的信息，通过write写入

## UDP客户端和服务端
多开几个终端跑，会发现一个client发送的信息，所有的client都能收到，是UDP的群发，server在中间充当转接和管理的作用

serer.go:
``` GO
package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	// 监听 UDP 端口
	addr, err := net.ResolveUDPAddr("udp", ":8081") //进行地址转换
	if err != nil {
		log.Fatal("解析地址失败:", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("监听端口失败:", err)
	}
	defer conn.Close()
	fmt.Println("已监听 UDP 端口:", addr)

	// 客户端地址列表
	var clients []net.Addr
	var mutex sync.Mutex

	// 读取和广播
	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("读取失败:", err)
			continue
		}
		msg := string(buffer[:n])
		fmt.Printf("收到 %s: %s", clientAddr.String(), msg)

		// 添加新客户端
		mutex.Lock()
		if !contains(clients, clientAddr) {
			clients = append(clients, clientAddr)
			fmt.Printf("新客户端: %s\n", clientAddr.String())
		}
		mutex.Unlock()

		// 广播消息
		mutex.Lock()
		for _, addr := range clients {
			if addr.String() != clientAddr.String() { //避免消息回显给发送者，实现群体发送的效果
				conn.WriteToUDP([]byte(fmt.Sprintf("[%s]: %s", clientAddr.String(), msg)), addr.(*net.UDPAddr))
			}
		}
		mutex.Unlock()
	}

}

func contains(clients []net.Addr, addr net.Addr) bool {
	for _, c := range clients {
		if c.String() == addr.String() {
			return true
		}
	}
	return false
}
```
- 进行UDP地址转换
- 建立conn连接
- 建立客户列表
- 通过buffer读取信息
- 通过list进行广播，防止发送给发送者

client.go
``` GO

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 连接服务器
	addr, err := net.ResolveUDPAddr("udp", "localhost:8081")
	if err != nil {
		log.Fatal("解析地址失败:", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()
	fmt.Println("已连接到 UDP 服务器")

	// 启动 goroutine 读取消息
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				log.Println("读取失败:", err)
				return
			}
			fmt.Print(string(buffer[:n]))
		}
	}()

	// 发送消息
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text() + "\n"
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("发送失败:", err)
			return
		}
	}
}
```
- udp地址转换
- 通过conn连接
- 启动一个协程一直监听，这样能输出收到的消息
- 同时利用scan，自己也能write发送消息


## WebSocket 服务器与客户端（实时聊天）

WebSocket 提供全双工通信，适合实时应用。以下使用 gorilla/websocket 实现聊天室
多启动启动几个前端文件，然后可以实现一个发消息，别的接收到

server.go:
``` Go

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源的连接
	},
}

func main() {
	// 客户端连接管理
	var clients = make(map[*websocket.Conn]bool) //定义函数在外部，每次访问都调用函数，修改 
	var mutex sync.Mutex

	//WebSocket路由
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) //upgrader.Upgrade 的作用是将 HTTP 连接升级为 WebSocket 连接
		if err != nil {
			log.Println("升级失败:", err)
			return
		}
		fmt.Printf("新客户端连接:%s\n", conn.RemoteAddr().String())
		// 添加客户端
		mutex.Lock()
		clients[conn] = true
		mutex.Unlock()

		//处理消息
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("客户端断开连接:%s\n", conn.RemoteAddr().String())
				mutex.Lock()
				delete(clients, conn)
				mutex.Unlock()
				conn.Close()
				return
			}
			broadcast := fmt.Sprintf("[%s]: %s\n", conn.RemoteAddr().String(), msg)
			fmt.Print(broadcast)

			//广播消息
			mutex.Lock()
			for c := range clients {
				if c != conn {
					c.WriteMessage(websocket.TextMessage, []byte(broadcast))
				}
			}
			mutex.Unlock()
		}
	})
	// 启动服务器
	fmt.Println("WebSocket 服务器启动，监听 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) //使用 log.Fatal 可以立即记录错误并终止程序
}

```
- 构建conn，每个客户端访问都会调用函数修改clients
- 通过conn读取信息，并判断客户端是否还连接着
- 通过clients列表广播消息
- 正常到http包，启动服务


前端html文件： （了解就行）
``` html
<!DOCTYPE html>
<html>
<head>
	<title>WebSocket 聊天室</title>
</head>
<body>
	<input id="message" type="text" placeholder="输入消息">
	<button onclick="sendMessage()">发送</button>
	<div id="output"></div>
	<script>
		const ws = new WebSocket("ws://localhost:8080/ws"); //建立 WebSocket 连接
		ws.onmessage = function(event) { //处理接收消息（事件监听）
			const output = document.getElementById("output");
			output.innerHTML += "<p>" + event.data + "</p>";
		};
		ws.onclose = function() { //处理连接断开
			alert("连接断开");
		};
		function sendMessage() { //发送消息函数：
			const input = document.getElementById("message"); //获取输入
			ws.send(input.value);
			input.value = "";
		}
	</script>
</body>
  </html>
```
- 简单的输入框设计
- js部分见注释，先建立连接，然后处理消息接收，断开和发送 三个组成部分




终端中打开html文件：
``` bash
# Windows
start file.html  # 或 start chrome file.html

# macOS
open file.html  # 或 open -a "Google Chrome" file.html

# Linux
xdg-open file.html
```

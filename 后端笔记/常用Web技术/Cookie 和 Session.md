## Cookie
- 在http.cookie包中
- 用于登陆系统，解决http无状态问题
- 服务器端会设置set cookie（一堆键值）通过http发送给客户端存储下来，访问的时候带上
	- 在服务器端设置会方便HttpOnly 安全一些
	- set cookie位置，在登陆成功后设置
- 最大问题是在客户端上是明文展示的
	- 设置secure部分，这样就可以只让https访问
	- path是访问路径，domain

## Session
和cookie类似平替的操作
- 设置在服务器中
- 认session_id不认人
	- cookie/head/query中携带过去
- session是放在了ctx中传递的 GO中

目前都是基于Gin框架的cookie和Session模块
- 为什么关闭程序后重新启动，直接访问index_c还可以直接访问？但是如果是session的话似乎就不行？
	cookie是存在客户端的浏览器中（发送时携带），而session是存储在服务器端！！！！


## 跨域问题
preflight 请求是浏览器在发送实际的跨域请求之前，先发送的一个 OPTIONS 请求，用于检查实际请求是否安全
[[Pasted image 20250329182555.png]]

- 请求头中Headers
	- 比如content-type，accept，origin，这样服务器才知道数据是什么格式（Content-Type）客户端想要什么格式的响应（Accept） 请求从哪里来（Origin）
	- Authorization 可以用于携带cookie等，用于验证







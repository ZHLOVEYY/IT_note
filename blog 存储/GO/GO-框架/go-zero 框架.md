## 简单了解
主流的 GO-rpc 框架除了 go-zero 外还有：gRPC-GO，go-micro，Kitex，Trift（GO）
性能优先​​：KiteX 或 gRPC。
​​跨语言需求​​：gRPC 或 Thrift。
​​微服务生态​​：go-zero（集成治理工具）或 Go-Micro

微服务就是会拆分为很多个小的服务，这样不会因为一个有问题就需要关闭所有的
![[../../../attachments/Pasted image 20250611145703.png]]

gozero 官方参考的项目以及官网文档： https://go-zero.dev/docs/reference/examples
很多比如：并发控制、自适应熔断、自适应降载、自动缓存控制这些功能都是在用 go-zero-gen 生成。
像限流功能就是在.yaml文件下增加相关的配置就行（加载的时候会自动进行配置），`CpuThreshold`、`MaxConns` 等这些也都是可以配置的

官网的测试文档所在的界面： portal
```bash
git clone https://github.com/zeromicro/portal.git
cd portal
git checkout feat/v3  # 切换到指定分支（这个好像才是对应的是网页）

# 启动部分
npm install --legacy-peer-deps
export NODE_OPTIONS=--openssl-legacy-provider # 设置node 17以上版本的安全问题
# 然后启动项目 
npm run start

```

部分演示文档所在的界面：zero-doc
之下的 website 文件夹也可以用上面的方法同样跑起来，估计是以前的官网界面
[快速构建高并发微服务](https://github.com/zeromicro/zero-doc/blob/main/doc/shorturl.md)通过查看 url 来看对应的文档路径进一步查看
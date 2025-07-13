[腾讯云文档](https://console.cloud.tencent.com/lighthouse/instance/detail?searchParams=rid%3D1&rid=1&id=lhins-nd8hu8yt&tab=firewall)
#### 本地测试
docker 中的 exec 检查mysql -u user -p ，用于检查本地的是否 ok  （数据库情况）
注意终端中操作需要 docker exec it  加在指令前面
`docker exec -it <容器名称或ID> /bin/bash`   结合 sql 自己查

本地的docker-compose up --build 可以运行起来说明就是完全 ok 的，就是迁移上云
#### 安装docker-compose
``` bash
sudo curl -L "https://github.com/docker/compose/releases/download/v2.23.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose  # 增加权限查看版本
docker --version   # centos9 腾讯云上自带的
docker-compose --version
```

#### 配置密钥
- 在控制台设置密钥，会自动下载文件到本地
- cp ~/Downloads/1234.pem ~/.ssh/ 复制密钥到专门的 ssh 文件夹（部分人可能没创建过）

https://cloud.tencent.com/document/product/1207/44643    ssh 连接
- chmod 600 ~/.ssh/1234.pem      记得添加权限，不然会认定为不安全

#### 传输和解压
`tar -czvf gin-mysql-demo.tar.gz .`  本地进行文件压缩  在目标的文件夹下面（下面！）解压！同时这个名字是自己定义的文件的名字，需要注意


- scp 上传似乎有问题，`scp gin-mysql-demo.tar.gz root@<服务器公网IP>:/root/`
好像没有开启，不过腾讯云的界面中，直接上传是可以的，更加方便 （一个电脑的图标）

- root 下有一个mydemo_ginsql.tar.gz 压缩文件，该怎么新建一个文件夹然后解压进去：
（千万别直接解压不然都跑出来了）

``` bash
mkdir -p mydemo_ginsql
ls -al  #查看
tar -zxvf ./mydemo_ginsql.tar.gz -C ./mydemo_ginsql
```

#### 安装 GO
为了 docker 中 go mod 下载更快 -> 设置 goproxy->需要安装 go

- 为 dnf 配置源
``` bash
sudo sed -i 's|mirrorlist=|#mirrorlist=|g' /etc/yum.repos.d/CentOS-*
sudo sed -i 's|#baseurl=http://mirror.centos.org|baseurl=https://mirrors.tencent.com/centos|g' /etc/yum.repos.d/CentOS-*
sudo dnf clean all
sudo dnf makecache
```
- 下载 golang
``` bash
sudo dnf install golang
```
- 设置相关环境
``` bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc # 添加环境变量
source ~/.bashrc  
go env -w GOPROXY=https://proxy.golang.com.cn,direct # 添加代理
```

- 发现原来需要在 docker 中配置相关的 goproxy！  因为是在容器中进行的，而不是本地，本地自己配置过终端代理所以 docker 中也可以很快
	- 这里可以看我的代码中

#### 启动服务
dockerfile 中添加 env！！
``` text
# 设置 GOPROXY 环境变量!!!
ARG GOPROXY
ENV GOPROXY=${GOPROXY:-https://mirrors.tencent.com/go/,direct}
ENV GO111MODULE=on
```
看看代码，我是改好了的
接着 `docker-compose up --build`

#### 连接测试
``` bash
curl -v http://localhost:8080/users # 测试连接
sudo iptables -L -n | grep 8080 | grep 8080 # 测试防火墙
```

- 注意，腾讯云的服务器需要你自己设置端口开放，自己添加 8080 端口设置规则！！！！ （除了本地的防火墙系统中的关闭以外）

- 接着就可以愉快的测试了！
http://你的公网ip:8080/users   GET和 POST 轮流测试

- 测试完记得及时关闭端口！保证安全性 


#### 补充（不需要看）：
- 更新软件包
```
sudo dnf update -y
sudo dnf upgrade -y
```
```
sudo dnf install -y vim wget curl git zip unzip
```
- 防火墙和端口相关
```
# 安装防火墙
sudo dnf install -y firewalld

# 启动防火墙并设置开机自启
sudo systemctl start firewalld
sudo systemctl enable firewalld

# 开放常用端口（根据需要调整）
sudo firewall-cmd --permanent --add-port=22/tcp    # SSH
sudo firewall-cmd --permanent --add-port=80/tcp    # HTTP
sudo firewall-cmd --permanent --add-port=443/tcp   # HTTPS
sudo firewall-cmd --reload
```

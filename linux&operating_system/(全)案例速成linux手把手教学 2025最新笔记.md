- 个人笔记仅供参考，速成完后会在个人github仓库[linux-operating_system](https://github.com/ZHLOVEYY/linux-operating_system)慢慢会完善补充操作系统等知识。（真的，CSDN更新博客，传照片太麻烦了）
- 基于CentOS学习，这个听说企业用的比ubuntu多
- 大伙要是有看不懂的多问AI，一开始有看不懂很正常
- 链接为优质博文如有侵权联系删除～
## 学习资源
（参考）导航：[鱼皮的总结](https://zhuanlan.zhihu.com/p/420311740 ) （还没用上）
用书：Linux运维从入门到精通 （还没用上）
一开始看不懂就对了，因为不知道知识，也不知道重点


AI生成目标大纲：[2天](https://link.csdn.net/?from_id=147258373&target=https%3A%2F%2Fi-blog.csdnimg.cn%2Fdirect%2F8e9799f102114a7f8be541eb2a4c2b09.png)  （grok3生成）
AI生成一周学习目标大纲：[一周](https://i-blog.csdnimg.cn/direct/7e4ba32b229d4c2287d84ad8348306f9.png)
那么我们就按照大纲来学习吧！（也是靠AI指引）

## 学习注意事项

- 注意目标操作的功能实现，实现了什么功能，而不是死背代码！（之前自己踩过的坑）
- 20%的操作可以解决80%的问题，适用于80%的场景，别钻牛角尖。一些什么-n，-s，-m敲多了就知道通用性含义了，不用记（很多书给人的误导真的比较大）
- 网络类一定想到两个问题：校园网，防火墙  检查检查
## 学前知识了解（科普）
- 运维分类：
应用运维：负责上线服务，服务变更回滚等
系统运维（初期只有）：IDC，网络，CDN，负载均衡建设
数据库运维
安全运维
运维研发：语法运维平台

- 95%都是线上服务器了，现在一般不会搞本地传统的

## 安装和配置

### 安装
- 参考链接这里面的配置过程是没问题的，不过里面是CentOS7没更新了，自己换源折腾了很久，最后还是在官网（自己搜，下载ISO镜像）下载了CentOS stream 10：[安装centOS和配置教程](https://blog.csdn.net/weixin_42364929/article/details/143487396)
- 注意安装的时候注意要把自己的用添加为管理组，不然进去后sudo命令执行不了。以及记得开启wifi
- 我为了方便操作先安装的带有图形化的，以及选了中文
### 配置
- 了解到dnf已经替代了yum～
#### 换源（必看必看必看）
- 换源么？我没换，搞CentOS7麻木了。但是我science net了，关于虚拟机science net以及终端设置代理，可以参考我的这两篇文章：
[终端代理science net](https://blog.csdn.net/Carlos5en/article/details/147233730)
[虚拟机连物理机代理端口](https://blog.csdn.net/Carlos5en/article/details/147152656)

#### 配置物理机和虚拟机共享剪切板 （必看必看必看）
方便复制和粘贴指令操作
`sudo dnf install -y open-vm-tools open-vm-tools-desktop`
mac是Crtl+shift+v在终端中粘贴（因为虚拟机中还是用ctrl的），自己一直笨笨用cmd，直到我右键复制粘贴发现是已经可以的了

#### 输入法切换
- 首先需要添加源，设置中找到“输入源”添加就行（注意“汉语”不是拼音，需要加入带有“智能拼音”的）
- 在“键盘”自己定义 我设置为的Ctrl+w （因为mac键盘似乎别的识别不太灵）

## 文件和权限操作
- 关于基本的文件目录结构，大家根据自己具体的虚拟机型号搜索查询就行，主要知道个大概就行，一般一开始我们都在/home/username目录下，代表你自己的个人用户文件夹，linux中一切都是文件（夹），可以自己先cd指令了解下到处走走看


大致纲领：
- 掌握文件操作命令：ls, cd, mkdir, cp, mv, rm, touch, find, ln。
- 学会查看文件内容：cat, less, head, tail, wc。
- 管理权限：chmod, chown, chgrp, umask。
- 管理用户：useradd, passwd, sudo, id。
- 完成练习任务：创建目录和文件、设置权限、创建用户
### 文件操作
目标：学会创建、复制、移动、删除文件和目录，创建链接。

- /tmp 并不是用户文件夹（如 /home/username）下的目录，而是位于​​根目录（/）​​下的系统级临时目录，用于存放系统和用户程序运行时产生的临时文件重启的时候会自动清理
``` bash
mkdir /tmp/test   # 创建目录和文件
ls /tmp  # 可以看到test目录
cd /tmp/test  # 切换到test
pwd # 确认当前路径，应该输出 /tmp/test
touch file1.txt file2.txt # 创建文件
ls # 应该看到file1.txt和file2.txt
cp file1.txt file1_copy.txt # 复制 
ls # 应该看到file1.txt, file2.txt, file1_copy.txt
mv file2.txt /tmp/  # 移动文件，tmp后的/有和没有区别不大
ls /tmp # 能看到file2.txt（应该在最上方，别漏看）
ls #当前还在test目录下，应该没有file2.txt。
ln -s file1.txt file1_link.txt # 创建软连接

```
- 软连接：可以理解成类似快捷方式，源文件没了就没了
- 硬链接：为原文件（file1）的 ​​inode​​（文件在磁盘上的唯一标识）创建一个新的目录项（file1_hard.txt）。两者共享相同的 inode 和数据块。所以修改 file1 或 file1_hard.txt 都会同步更新同一份数据，和复制不一样。inode通过ls -l就能查看的
我们继续～
``` bash
ls -l # 应该看到file1_link.txt -> file1.txt（箭头表示软链接）
ln file1.txt file1_hard.txt # 创建file1.txt的硬链接file1_hard.txt
ls -l # file1_hard.txt和file1.txt的inode号相同（用ls -li查看第一列）
find /tmp -type f -name "*.txt" # 查找/tmp下所有.txt文件，可能包括/tmp/file2.txt, /tmp/test/file1.txt
rm file1_copy.txt # 删除file1_copy
ls # 应该没有file1_copy.txt
```

收获：创建（mkdir, touch）、复制（cp）、移动（mv）、删除（rm）、查找（find）和链接（ln）操作
### 文本查看
目标：学会查看和统计文件内容。
``` bash
vim file1.txt # vim是文本操作器，补充学习放下面了
#按i进入插入模式，输入以下内容
#Line 1: Hello, this is a test file.
#Line 2: CentOS Stream 10 is great.
#Line 3: Learning Linux for backend.
#按Esc，输入:wq，保存退出。
cat file1.txt # 查看全文
less file1.txt # 分页查看 space翻页，q退出
head -n 2 file1.txt # 查看前两行
tail -n 1 file1.txt # 查看最后一行
wc -l file1.txt # 统计行数
head -n 10 /etc/passwd # 查看系统文件
head -n 10 /etc/passwd # 统计行数

```
收获：学会了初级vim，查看（cat, less, head, tail）和统计（wc）文件内容。

##### vim常用学习了解
命令模式：
	命令模式下按I进入编辑，按esc退出编辑模式，输入:写指令比如:wq
	常用   :w      # 保存
		  :q      # 退出（如有修改会提示）
		  :q!     # 强制退出，不保存修改
		  :wq     # 保存并退出
		dd   # 删除当前行
		yy   # 复制当前行
		u    # 撤销
		u：撤销 ！！！！
		ctrl+r # 重做（取消撤销）!!!
		v         # 进入可视模式     按上下左右进行移动，gg到头，shift g到尾
	查找功能：
		/word     # 向下查找
		?word     # 向上查找
		n         # 继续查找下一个
		N         # 继续查找上一个
	编辑模式快捷键：（常用！！

### 权限管理
目标：学会设置文件权限。
- 关于权限的解释：
权限 755 是一个三位数的数字表示法，分别对应​​所有者（User）​​、​​所属组（Group）​​和​​其他用户（Others）​​的权限组合。每个数字由​​读（r）、写（w）、执行（x）​​三种权限的数值相加而成，具体如下：
![[../attachments/Pasted image 20250415205624.png]]
读（r）​​ = 4
对文件：可查看内容；对目录：可列出文件列表。
​​写（w）​​ = 2
对文件：可修改内容；对目录：可创建/删除文件。
​​执行（x）​​ = 1
对文件：可运行（如脚本）；对目录：可进入（cd）

``` bash
ls -l file1.txt # -rw-r--r-- 1 user user ...。
chmod 750 file1.txt # 设置权限为rwxr-x---（750）
ls -l file1.txt # 检查
umask # 常见默认值：0022，新文件的默认权限是666（文件）或777（目录），减去umask值。
# 文件：666 - 022 = 644（rw-r--r--）

# 更改文件属性和组
ls -l file1.txt # 查看发现是rwxr-x--- 1 user user ...（用户和组都是当前用户）。
sudo chown root file1.txt # 更改属主为root
sudo chgrp root file1.txt # 更改属组为root
ls -l file1.txt # 验证查看


```
收获：已经学会了权限管理（chmod, chown, chgrp）和umask

### 用户管理
目标：学会创建用户并配置权限。
``` bash
sudo useradd -m -s /bin/bash testuser # 创建用户testuser 
#-m：创建家目录。
#-s /bin/bash：设置默认Shell为bash。
#若不使用 -m，用户将没有独立的主目录，可能导致某些程序无法正常运行 -s 选项​ 指定用户的默认登录 Shell

sudo passwd testuser #输入新密码（比如123456） 说密码不对正常输就行
id testuser #显示uid, gid等信息。
grep testuser /etc/passwd # 看到testuser一行
sudo usermod -aG wheel testuser #添加到wheel组
su - testuser #切换用户
sudo ls /root #检验权限，输入testuser密码，预期：能列出/root内容
exit #退出
 
```

收获：已经学会了用户管理（useradd, passwd, sudo）


### 综合练习
看懂先
``` bash
mkdir /tmp/scripts
cp /etc/*.conf /tmp/scripts/ #将/etc下所有.conf文件复制到/tmp/scripts
ls /tmp/scripts
chmod 755 /tmp/scripts
ls -ld /tmp/scripts # ls -ld查看文件夹的权限
ln -s /tmp/scripts /tmp/scripts_link
ls -l /tmp #scripts_link -> /tmp/scripts
```



## 进程、网络与系统管理

可以在之前mkdir创建的/tmp/test下继续进行练习～
## 进程管理
### 查看进程
``` bash
ps aux  # 查看所有进程
```
- USER：进程所属用户。
- PID：进程ID。（重点记忆）
- %CPU：CPU使用率。
- COMMAND：运行的命令。
``` bash
ps aux | grep bash  # 筛选bash进程
```
|：管道的概念，将左边命令输出作为右边输入
``` bash
top #实时监控，观察CPU、内存使用率，按q退出
```
htop 需要在EPEL下载，可以提供比top更清晰的视图查看
``` bash
sudo dnf install -y epel-release
sudo dnf repolist | grep epel  检查是否已经启用
```
EPEL（Extra Packages for Enterprise Linux）是一个由 ​​Fedora 社区​​维护的第三方软件仓库，专门为 ​​Red Hat Enterprise Linux (RHEL)​​ 及其衍生发行版（如 ​​CentOS、Rocky Linux、AlmaLinux、Oracle Linux​​ 等）提供额外的开源软件包，相当于扩展官方的软件包
``` bash
htop #按F5显示进程树，q退出。
```

### 启动和终止进程
``` bash
sleep 1000 & 
```
sleep 1000 & 命令的作用是​​启动一个后台进程​​，该进程会暂停执行（即“睡眠”）1000 秒，会挂起 1000 秒后退出。&的作用是放到后台，防止阻塞当前shell，适合异步任务
``` bash
jobs #查看当前终端后台进程，显示sleep 1000的任务。
jobs -l  # 显示任务列表及其 PID
ps aux | grep sleep #找到PID（假设是1234）
kill 1234
ps aux | grep sleep # 显示的是查找指令本身的一个进程，涉及sleep

sleep 1000 & sleep 2000 &#再启动几个进程
killall sleep # 批量终止
ps aux | grep sleep # 验证

# 前后台切换
sleep 1000 #前台进程
按Ctrl+Z，会暂停并放入后台。
jobs # 查看后台任务 ，显示sleep 1000（可能标记为[1]+ Stopped）
fg %1 # 放到前台，%1：任务编号（jobs输出的编号）
按Ctrl+C终止
sleep 1000 # 前台进程
按Ctrl+Z
bg %1
jobs  # 进程在后台运行
```

- 这里引申到：进程和后台？
前台进程：直接与终端关联，占据终端输入/输出的进程，如vim，top，按ctrl+c停止
后台进程：脱离终端控制，在后台中运行，但是若未使用 nohup（后面会继续涉及到），终端关闭时可能被终止！后台适合长时间任务比如wget
守护进程：后台进程的一种特殊，完全脱离终端，由系统控制
- 这也就能理解ps aux 和 jobs 的区别：
ps aux：所有正在运行的进程（包括其他用户、终端、守护进程）	
jobs：​	仅当前终端启动的后台/暂停任务

- 出一道题：如果有一个后台进程sleep，能想到多少种方法终止？
1. 可以jobs 查看序列号，然后fg %1 ，Ctrl+c
2. 1基础上看完jobs后直接kill %1
3. jobs -l 然后根据 pid号删除比如 kill 1234
4. 执行 ps aux | grep sleep 然后kill 1234 (注意这个会看到有两个结果，因为调用ps aux找sleep本身也是一个进程，聪明的你肯定可以马上看出不一样 （说明的指令为grep --color=auto sleep）)


收获：已经学会了查看（ps, top, htop）、终止（kill, killall）和管理（jobs, fg, bg）进程


### 管道与重定向
目标：学会使用管道和重定向处理命令输出
``` bash
ps aux | grep bash > proc.txt # 会保存在当前目录下
ls -l proc.txt  查看文件是否存在
# ps aux | grep bash > /path/to/your/dir/proc.txt 如果指定路径
cat proc.txt # 查看文件
ps aux | grep bash >> proc.txt # 追加输入 >> :追加不会覆盖
cat proc.txt #文件内容应该变长

# 重定向
ls /notexist 2> error.txt
```
解释：
- ls /notexist​​：尝试列出名为 /notexist 的目录或文件内容。若该路径不存在，ls指令会输出错误信息到​​标准错误（stderr）​​。
- 2> error.txt​​：将标准错误（对应文件描述符 2）的输出重定向到文件 error.txt 中，而非默认的终端屏幕。（Linux 中命令的错误输出默认关联到文件描述符 2，与标准输出（stdout，文件描述符 1）分开处理）

``` bash
cat error.txt # 查看错误信息的文档
#应该包含类似ls: cannot access '/notexist': No such file or directory

#重定向输入
echo "hello" > input.txt #自动创建input.txt
cat < input.txt # 用 cat读取输出，输出hello。


# 后台任务持久化
nohup sleep 1000 &
ps aux | grep sleep # 能看到sleep 1000进程
cat nohup.out # nohup会生成nohup.out文件，sleep没有out，所以为空
ls -l nohup.out # 验证是有这个文件的
```


- nohup虽然持久化，但是另外一个新终端中jobs是不行的，因为jobs只是查看当前终端的后台，需要ps aux ｜grep

收获：已经学会了管道（|）和重定向（>, >>, <, 2>）。
小记录：
- |：管道，左边输出给右边输入
- >:覆盖重定向。
- >>:追加重定向。
- nohup command &：后台持久运行

## 网络和系统管理

### 查看网络信息

``` bash
ip addr #查看ip地址
# 找到你的网卡（如eth0或ens33），记录IP（可能是192.168.73.x）
ss -tuln # 检查端口，列出监听端口
sudo dnf install -y net-tools # 以便使用netstat
netstat -tulnp # netsat功能类似ss，显示监听端口和进程

## 测试网络
ping -c 4 baidu.com #-c 4：发送4次请求，预期收到4次
# 不行的话就看看网络，NAT桥接什么的打开没，以及校园网！！
curl https://baidu.com # 之前前面换源教程的那里是curl google～
#预期返回html

```
curl 是一个用于传输数据的命令行工具（支持 HTTP、FTP 等协议），核心功能包括：
- 发送请求​​：获取网页内容、API 数据等。
- 调试网络​​：查看响应头、状态码（如 -I 显示头部）。
- 下载文件​​：通过 -o 保存到本地。
但是发现，回答很短小精悍：
![[../attachments/Pasted image 20250415205703.png]]
这是因为：​​
- 302 状态码​​：表示临时重定向，浏览器或 curl 会收到新地址（需手动跟随跳转）。
- 服务器标识​​：bfe/1.0.8.18 是百度自研的负载均衡服务器（Baidu Front End）的版本信息，所以是到这里就返回了，没有继续深入
所以需要结合使用curl具体操作参数：
-L：自动跟随重定向 （可以查看完整的）
-v：显示详细请求和响应过程  （可以观察整个过程）
-o 文件名：输出保存为文件，如curl -o baidu.html https://baidu.com
``` bash
curl -L https://baidu.com # 能看到html

# 检查80端口
sudo dnf install -y nc # 安装nc，用于端口测试
nc -zv localhost 80 #测试80端口是否开放，现在是没开放的，因为你没有设置服务监听，下面还会讲到
```

收获：学会了网络操作（ip, ss, netstat, ping, curl）

### 包管理与服务管理
其实之前的dnf就一直是“包管理”的概念
``` bash
sudo dnf install -y nginx #安装nginx，用于提供监听服务
nginx -v # 显示版本
sudo systemctl start nginx # 启动服务
sudo systemctl status nginx # 检查状态
sudo systemctl enable nginx # 设置开机自动启动
sudo systemctl stop nginx # 关闭服务
sudo systemctl status nginx # 验证

sudo systemctl start nginx
ss -tuln | grep 80 # 这时候nginx就在监听了
curl http://localhost #返回nginx欢迎页面（HTML内容）


```

- ss -tuln | grep 80 显示 0.0.0.0:80 的监听状态，是因为 Nginx ​​默认监听 HTTP 协议的 80 端口​​（HTTPS 默认监听 443 端口）
- Nginx可以理解为把你的电脑高成服务器，别人要从你这拿东西就要从你的80端口来拿，所以开启nginx自然就会监听了
- 如果想修改监听端口可以修改Nginx 的配置文件

收获：已经学会了包管理（dnf）和服务管理（systemctl）

### 综合任务

``` bash
ss -tuln > network_status.txt # 保存网络状态
cat network_status.txt # 查看
ps aux | grep httpd > httpd_proc.txt # 保存httpd进程（如果有）
cat httpd_proc.txt # 查看，发现应该是没有的
```

![[Pasted image 20250414195544.png]]
结合前面学到的，这个其实就说明是只有ps aux这个查询命令
- httpd 是 ​​Apache HTTP Server​​（Apache 超文本传输协议服务器）的主进程名称，它是用于提供 Web 服务的核心程序。具体而言（全称 HyperText Transfer Protocol Daemon）是一个 ​​Web 服务器守护进程​​，负责监听 HTTP 请求（默认端口 80/443），处理客户端（如浏览器）的请求并返回网页内容，用于返回静态文件（HTML、图片等），或者调用脚本（PHP/Python）生成动态内容
- Apache 最常用～Nginx 使用事件驱动模型，资源占用更低，适合高并发
- Apache 功能模块更丰富，适合复杂需求 这个再拓展就涉及到LAMP还有LNMP的设计和组成啦～（个人浅浅了解过一点）
``` bash
sudo dnf install -y httpd
```

需要sudo systemctl stop nginx 不然会和httpd启动冲突（都监听8080端口）
``` bash
sudo systemctl start httpd # 启动
sudo systemctl status httpd
curl http://localhost # 预计返回Apache欢迎页面
ps aux  | grep httpd >> httpd_proc.txt  #追加查看
cat httpd_proc.txt
```

总结与记录：
- **命令掌握**：
    - 进程管理：ps, top, htop, kill, killall, jobs, fg, bg。
    - 管道重定向：|, >, >>, <, nohup。
    - 网络操作：ip, ss, netstat, ping, curl。
    - 包管理：dnf。
    - 服务管理：systemctl。


## 后续内容
如果感觉有用的话～后面的内容更新在个人github仓库[linux-operating_system](https://github.com/ZHLOVEYY/linux-operating_system)
包括shell脚本，部署服务器等等，个人也在同步整理和学习！
后面仓库还会补充各种linux以及操作系统相关的知识！


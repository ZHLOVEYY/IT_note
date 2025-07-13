20 % operation can solve 80 % problems 
# learning resources
guide: [鱼皮的总结](https://zhuanlan.zhihu.com/p/420311740 )

# knowledge 
- 95% of Operation and maintenance are online servers

## centos installation 
- centos 10 is preferred[安装centOS和配置教程](https://blog.csdn.net/weixin_42364929/article/details/143487396)
when installiing:
make sure to add your own app as a management group, and turn on  wifi option

dnf have taken the place of yum
#### change the net source
use proxy 
[终端代理science net](https://blog.csdn.net/Carlos5en/article/details/147233730)
[虚拟机连物理机代理端口](https://blog.csdn.net/Carlos5en/article/details/147152656)
#### Physical machines and virtual machines share the clipboard
`sudo dnf install -y open-vm-tools open-vm-tools-desktop`
easy for your test and copy the instruction
if you use mac ,remember that you need to use Crtl+shift+v to paste in virtual machine, and cmd +c in mac to copy

#### Input method switch
- in setting , choose “智能拼音”, add to source
- i set Ctrl + W to change the input method

# concept note
## crontab 
- explanation
	``` text
	* * * * *
	│ │ │ │ │
	│ │ │ │ └─ What day of the week (0-7) (0 or 7 is Sunday)
	│ │ │ └─── Month (1-12)
	│ │ └───── Date (1-31)
	│ └─────── Hour (0-23)
	└───────── Minute (0-59)
	```
- example
	``` bash
	# Executed at midnight every day
	0 0 * * * /root/xxxx/cleanshell/clean-es-posts.sh
	
	# every Monday morning
	0 3 * * 1 /root/xxxxx/cleanshell/clean-es-posts.sh
	
	# first day of each month.
	0 3 1 * * /root/xxxxx/cleanshell/clean-es-posts.sh
	
	# three o 'clock every afternoo
	0 15 * * * /root/xxxxx/cleanshell/clean-es-posts.sh
	
	# every hour
	0 * * * * /root/ xxxxx/cleanshell/clean-es-posts.sh
	```

## Files 
#### Create： mkdir touch
``` bash
mkdir /tmp/test   # Create directories and files
ls /tmp  
cd /tmp/test  # change to test
pwd # where are you now 
touch file1.txt file2.txt # Create files
ls # 应该看到file1.txt和file2.txt
```
#### copy：cp
``` bash
cp file1.txt file1_copy.txt # copy 
```
#### move： mv 
``` bash
mv file2.txt /tmp/  # 移动文件，tmp后的/有和没有区别不大 
```
#### delete files and directories : rm
``` bash
rm
rm -f # 
rm -rf # Recursively confirm deletion
```
#### create chains
``` bash
ln -s file1.txt file1_link.txt # Create a soft link
ln file1.txt file1_hard.txt # create a hard link of  file1_hard.txt
ls -l # check the soft/hard link (-> / inode )
```
soft link: like Shortcut, if the source file is lost, it's gone
hard link: Create a new directory entry on the file's inode (The unique identifier of the file on disk ), two share the same inode and data blocks. so modify will change at the same time
#### search :find 
``` bash
find /tmp -type f -name "*.txt" # search all the .txt file under /tmp
```
#### see content: cat 
``` bash
cat proc.txt # 查看文件
```
#### vim
``` bash
yy   # Copy the current line
u # undo 撤销
Ctrl + r # cancel undo
also search mode
```

#### scp to translate files
``` bash
scp -P 22 /Users/mac/Downloads/myfile.txt root@142.171.225.226:/root/
```
## permission
755 ： read write ex
![[../attachments/Pasted image 20250415205624.png]]
``` bash
chmod 750 file1.txt # set the permission as rwxr-x---（750）
umask -> set to count the permission

use chown / chrgp to change the owner / group of a file
```



## User Management
``` bash
sudo useradd -m -s /bin/bash testuser # create user : testuser 
#-m：Create a home directory
#-s /bin/bash：Set the default shell to bash。
#If -m is not used, the user will not have an independent home directory, which may cause some programs to fail to run normally. The -s option specifies the default login Shell for the user

sudo passwd testuser #Enter a new password (e.g. 123456) and say that the password is incorrect, just enter it normally
id testuser #show uid / gid
grep testuser /etc/passwd # 看到testuser一行
sudo usermod -aG wheel testuser #Add to the wheel group -> have the root permission 
su - testuser #change the user 
sudo ls /root #To verify permissions, enter testuser password, expected: can list the /root content
exit 
 
```

## Process management
``` bash
ps aux  # View all processes
ps aux | grep bash  # Filter the 'bash' process
# PS:Calling "ps aux" to find "xxx" itself is also a process (show in result)
top #observe CPU and memory usage rates, and press q to exit
```
 PID：Process ID (Key Points to Remember
 
 htop do better than top

#### background (terminal)
``` bash
sleep 1000 & 
# set a Background process, sleep 1000s, &: put to the background ,Prevent blocking of the current shell
jobs 
jobs -l  # check the background process in the current terminal
killall sleep # stop all the sleep progress 

# Front-end and back-end switching
sleep 1000 #fg frontgroud 
Crtl+Z -> put to the background 
jobs 
fg %1 # put to fg, mission %1
Crtl+C to stop
```
ps aux：All running processes (including other users, endpoints, daemons)
jobs：​	Only the background/pause tasks started by the current terminal

 nohup :  taks may be terminated when the terminal is closed, use nohup to solve
``` bash
nohup sleep 1000 &
ps aux | grep sleep 
cat nohup.out # nohup will generate the nohup.out file, sleep has no output, so it is empty
```
 

## Pipelines and Redirection
``` bash
ps aux | grep bash > proc.txt # will save as a txt 
ps aux | grep bash >> proc.txt # Append input >> : Append will not overwrite

# Redirection
ls /notexist 2> error.txt
stderr (File Descriptor 2) output to error.txt if cannot ls /notexist 
```

## net check/ port check

#### ss netstat
``` bash 
ip addr #Check the ip address, have different net card 
ss -tuln  # check the listening ports
ss -tuln | grep 80 # 这时候nginx就在监听了
netstat -tulnp # like ss
```

``` bash
sudo netstat/ss -tuln | grep -E ':80|:443'
```
#### nmap
``` bash
nmap -p 80,443 142.171.225.226
```
#### ufw
``` bash
sudo ufw status
sudo ufw allow 80 
sudo ufw allow 443 
sudo ufw reload
```
#### curl
Check the response header and status code
![[../attachments/Pasted image 20250415205703.png]]
bfe: baidu's load balancing server -> 302 means redirection
-L：Auto-follow redirection- > will show the true result
-v: Displays detailed request and response procedures
-s : no output
-o : output file  like : `curl -o baidu.html https://baidu.com`

check mywebsite:
``` bash
curl -I http://231114.top
```

## Service Management
 #### systemctl

``` bash
nginx -v # 显示版本
sudo systemctl start nginx # 启动服务
sudo systemctl status nginx # 检查状态
sudo systemctl enable nginx # 设置开机自动启动
sudo systemctl stop nginx # 关闭服务
sudo systemctl status nginx # 验证
```
# operation note
## File path problem
When executed in different folders, new files will be output in different folders(if you use ./ in your file). 
like:
``` bash
bash -x /root/xxxxx/cleanshell/clean-es-posts.sh （in shell you need absolute path）
```

## package management
a nodejs / npm install :

``` bash
sudo dpkg --configure -a
sudo apt update
sudo apt install -f # -f try to fix 

sudo apt remove --purge libnode-dev # purge will remove the setting files 
sudo apt autoremove 

curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y --reinstall nodejs # reinstall to cover the files 
sudo apt install -y npm
```
when you need to restart service, can use N to pass 

## nginx
concerning deploy webserer
if you buildup nginx -> show you can set things 
``` bash
sudo apt install -y nginx
```

######  upload a dist file 
``` bash
sudo cp /etc/nginx/sites-available/default /etc/nginx/sites-available/default.bak
sudo vim /etc/nginx/sites-available/myvue
```

``` nginx
server {
    listen 80;
    server_name 231114.top 142.171.225.226;

    location / {
        root /var/www/vuetest; # change to the dist route
        index index.html index.htm;
        try_files $uri $uri/ /index.html; # Support Vue Router
    }

    # 可选：错误页面
    error_page 404 /index.html;
}
```

```bash 
# set up soft link
sudo ln -s /etc/nginx/sites-available/myvue /etc/nginx/sites-enabled/
# (you can choose)
sudo rm /etc/nginx/sites-enabled/default

sudo nginx -t # check the grammar 
sudo systemctl restart nginx


```
i put in /var/www is preferred

``` bash
sudo mkdir -p /var/www/vuetest 
sudo mv /root/vuetest/dist /var/www/vuetest/ # move the file

# change the permisson
sudo chown -R www-data:www-data /var/www/vuetest
sudo chmod -R 755 /var/www/vuetest
```

check the log of nginx -> to know certain problems
``` bash
sudo tail -f /var/log/nginx/error.log
```
更多个人笔记：（仅供参考，非盈利）
gitee： https://gitee.com/harryhack/it_note
github： https://github.com/ZHLOVEYY/IT_note
（里面也有我的git教学）

本文基于mac，linux和win可以参考
个人同时配置gitee和github的ssh密钥过程，也算是又复习了一次。SSH一劳永逸
需要注意二者的覆盖问题

参考资源：[gitte的官方教学](https://gitee.com/help/articles/4181#article-header0)（但是也是不太全）


## 本地生成密钥
密钥都存在 `~/.ssh/`  下，`ls -al ~/.ssh/` 可以先简单查看
那么接下来先生成gitee的：
``` bash
ssh-keygen -t ed25519 -C "your_location@.com"
```
后面就是方便命名的，随便都可以的
接着就是按三次回车确认，不要覆盖了
生成github的密钥的时候做一个区分，不然就覆盖文件了
``` bash
ssh-keygen -t ed25519 -C "your_github_email@example.com" -f ~/.ssh/github_id_ed25519
```
接着：
``` bash
ssh-add ~/.ssh/ed25519   
ssh-add ~/.ssh/github_id_ed25519  #两个都添加
```

配置config文件：（可以配一下，不过不配置似乎好像也可以）
`vim ~/.ssh/config` 增加config文件
```bash
# Gitee配置
Host gitee.com
  HostName gitee.com
  User git
  IdentityFile ~/.ssh/id_ed25519  # 已覆盖的Gitee密钥

# GitHub配置
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/github_id_ed25519  # 新生成的GitHub密钥
```


## 网站添加公钥
github和gitee都是打开个人的设置，然后看到SSH添加，点击添加公钥就可以了
注意`ls -al ~/.ssh/` 中，pub就是公钥所在的文件夹

```bash
cat ~/.ssh/ed25519.pub
cat ~/.ssh/github_id_ed25519.pub
```

注意不要复制下面自己的邮箱了，那不是密钥，是上面说的，一开始自己命名的

## 测试
`ssh -T git@github.com`
`ssh -T git@gitee.com`
（如果还有问题有可能是文件权限问题）
分别测试链接，大功告成

PS：注意git中set orgin的地址要变，不能http
修改的话如：`git remote set-url origin2 git@gitee.com:xxx/xxx.git`
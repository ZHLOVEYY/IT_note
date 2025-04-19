更多个人笔记和后续持续更新可以查看本人Github仓库：
（本文参加CSDN活动，过几天再发到我的github上）

## 凝练理解
git凝练而言：方便你管理和记录自己各个版本的代码（理解成“修改历史”），新建不同的分支搞不同方向，以及与别人合作。
- 部分参考资料：（不如跟着我练，我都只存了都没看）
	[菜鸟教程git指令]( https://www.runoob.com/git/git-basic-operations.html)
	[字节的git飞书笔记](https://bytedance.larkoffice.com/file/boxcnqXgX9cP9uX5CVNGNeDZiLd)

### 知识扩展
常见代码托管平台：github，gitlab，Gerrit
## 相关软件推荐
可以先不安装，基本的学习通过我下面的案例就可以的，软件只是锦上添花
- github desktop  可以查看自己仓库各个版本的文件变化
- Sourcetree 可以以文件目录形式展示版本变化（这个功能似乎github desktop没有），同时以工作时间流的方式展示不同分支之间的关系
- VScdoe插件 gitlen  可以对比查看不同分支和版本文件的不同  （不过效果一般，别的一些IDE比如Goland自带的对比也很不错）
## 核心概念和操作
那么，我们开始吧！
- 案例是具有连续性的，方便大家快速上手，跟着操作就行
### git 安装
- git是需要安装的，mac自带，linux系的sudo install自己搜索下命令，win这里我提供一个我之前用的安装参考链接：[win上安装git](https://blog.csdn.net/weixin_42242910/article/details/136297201)
- 网络问题记得检查下是不是校园网等
### git基本功能入门
#### git工作区基本概念
- 目标：理解怎么知道创建git，理解工作区，暂存区，提交

先开启终端并cd 到一个文件夹下（这部分不懂的话就需要补补知识力）
建议在编程IDE中打开比如Vscode
``` bash
mkdir my-test
cd my-test

#初始化git仓库,相当于建立一个自己的“记录档案”，记录这个文件夹里文件的变化
git init 
#git是针对文件夹/工作区的概念而言的，新的大工程文件都需建立，但是工程文件下的文件夹中不用再建立

# 配置用户信息（可以不配也行，和别人合作的时候需要配置，方便别人联系你）
git config --global user.name "输入你的名字"
git config --global user.email "输入你的邮箱"

# 创建项目文件
echo "123" > 123.txt
echo "this is my first git" > README.md
# README.md 一般会被默认读取用来解释项目

# 查看自己的git 状态
git status

git add 123.txt
git status # 可以发现123.txt已经被添加

```

- 这里解释下git 的操作流程：
工作区 → git add → 暂存区 → git commit → 版本库
所以说你修改改动的文件，都需要git add 后才能真正的提交
- 这个过程中的add操作也可以可视化点击执行，比如VScode左侧边栏，大家可以自行探索

``` bash
git add . # 将所有文件添加到暂存区

# 提交到版本库，记得-m
git commit -m "第一个commit"

# 查看提交历史
git log 
# git log --oneline 可以查看简化后的
```
- git log可以查看自己commit的历史，非常有用，展示的hash数值是对应commit特有的，可以用于后面的回溯
- 对比文件等功能也有对应的指令，比如查看特定提交`git show commit-hash -- a`  以及对比两次提交不同：`git diff commit-hash1 commit-hash2 -- a `等但是这部分其实通过前面推荐的软件或者IDE以及插件，可视化看会更好，所以不用记。
commit后面的就是hash值 ，后面讲到回滚等会有涉及（主要记这个hash值就ok力）
![[attachments/Pasted image 20250416095454.png]]

#### 分支操作入门
- 目的：了解分支的概念以及基本操作
这时候你心想，不行，我仅仅只写了123，但是我可能需要给一些人的版本需要456，怎么办？
你可能马上想：
``` bash
vim 123.txt
```
然后改123为456，接着一顿git add，git commit。但是这时候你会发现，你不能同时给别人提供123了，因为456覆盖了123。所以当涉及到多版本，探索试错，合作的时候，分支的概念就非常重要

``` bash
git branch main2  # 一般会根据该分支的特点命名，我这图方便用main2
git branch #查看分支
git status # 发现自己还在Main分支上，工作区也干净着
# 切换到新分支
git checkout main2
#git branch -m main3 给分支改名
#git checkout -b main2 创建并同时切换
git status 
```
这时候你再修改你的123.txt文件，把之中的123改为456
用vim或者vscode直接编辑都行
``` bash
git status # 发现有未提交的123.txt
git add .
git commit -m "我的main2第一次commit"
git log
```
可以发现自己有新的提交，在分支main2上～
``` bash
git checkout main
```
这时候切换回主分支，发现123.txt文件的内容又变回123了～大功告成

- 分支其实是一门大学问，比如分支合并（如果你想让你的123和456合并一下），分支冲突（123和456该用谁？）等，后续再继续探讨

### 远程仓库的上传
- 目的：了解remote，创建仓库，配置ssh，push上传
远程仓库涉及到remote的概念，remote中存储了远程的仓库
我们还是通过例子来学习
``` bash
git remote -v #查看当前git下的remote库
```
- 你会发现什么都没有，因为一开始git是没有配置remote的，需要你自己配置
	那么在配置之前，你需要先会自己创建仓库：
	这里我采取github进行讲解，需魔法
- 常见的配置是：`git remote add origin https://github.com/your-username/myproject.git`
- origin是配置的名字，一般一个本地工程文件对应一个仓库所以就是origin，如果你喜欢多个可以origin1，2等等 但是很快后面执行 `git pull origin main` 的时候会有HTTP的问题，这是因为终端没有配置代理以及各种问题导致的，（同时Http后期使用也会有各种问题比如输密码）如果感兴趣一定一定看我的文章：[终端配置代理，这下不用换源了！](https://blog.csdn.net/Carlos5en/article/details/147233730)。git其实也可以配置代理，`git config --global --get http.proxy`可以查看代理端口，后面如果遇到一些clone等的问题可以自己在这方面搜索了解下，类似的方法进行设置一劳永逸。

上面都是小科普，出于安全性和方便性考虑，建议大家配置SSH，一劳永逸（只用在一开始配置），参考站友教程：[给github配置SSH]( https://blog.csdn.net/PleaseBeStrong/article/details/139378481)
这是我当时配置的笔记：
- 不选择RSA，自己：ssh-keygen -t ed25519 -C "your_github_email@example.com"（自己的邮箱）  （之中yes，enter就行）
- 然后cat ~/.ssh/id_ed25519.pub  查看完整的公钥（私钥也生成） 然后复制到github上的setting中增加SSH认证就可以了
- ssh -T git@github.com 尝试链接
- git remote set-url origin git@github.com:yourname/project.git 修改origin
	- git remote add origin git@github.com:yournameproject.git

好的那么我们学会创建仓库，配置SSH后继续案例学习：
``` bash
git remote add orgin git@github.com:yourname/project.git
# 填入自己创建好的仓库,可以github的code处直接复制ssh
# git remote rm origin 可以删除一些remote，如果你配置错了。或者set-url重新设置。
git pull origin main 
# 一般不会有冲突，空仓库最多licence和readme
git pull origin main --rebase #如果有矛盾的话可以尝试
# 可以自己再修改以及add，commit  拉取这样的行为也记录下
git push origin main  # 将mian分支推送上去，别的分支不会变
# 需要先拉取～记得

```
到这里，你就可以个人构建自己的并进行简单的开发管理啦

### 远程仓库的拉取
- 目标：学会clone别人的仓库学习
拉取是学习别人代码等的第一步～ 
github 中点击右上角的code，复制别人的http/ssh链接，就可以clone
``` bash
git clone https://github.com/user/project.git #自己找一个，或者拉我的笔记也行：https://github.com/ZHLOVEYY/IT_note.git
```
git clone自动初始化本地仓库，无需手动执行 `git init`

具体还有fetch，pull等，由于涉及到冲突（冲突的种类太多了，不好举例，最好实际问题实际解决），请看最后的动态语法库了解概念

### git 事务回滚
- 目标：学会revert，reset，文件回滚以及之间区别，回滚主要用于解决bug等


好的我们现在继续刚才的demo
``` bash
git checkout main
echo "456" >> 123.txt  # 追加修改内容
git add .
git commit -m "增加了456"
git log 
echo "789" >> 123.txt
git add .
git commit -m "增加了789"
cat 123.txt #查看文件，记得在当前文件夹终端下
git log # 查看
echo "bug" >> 123.txt 
git add .
git commit -m "好像有bug"
git log --oneline  #看自己的记录，简介版
```

#### revert回滚
接下来我们利用revert进行回滚
``` bash
git revert HEAD # HEAD可以替换成具体commit的hash值，由于这里是最近的一次提交有bug，就直接用HEAD

```
这会打开一个编辑器，让你确认提交信息（可以直接保存退出 ，输:wq）。
git revert HEAD 会创建一个新的提交，撤销对应的修改
```bash
cat server.js # 或者直接看也行
```
发现文件恢复到`git commit -m "增加了789"`的时候
``` bash
git log --oneline
```
发现多了一个commit：Revert "好像有bug"

- revert适合需要保留历史记录的场景（如团队协作中已推送到远程仓库）。会保留之前的commit！！！基于之前有bug的commit重新commit放在
理解：
A---B（bug）---C---D
A---B（bug）---C---D --- B‘（修改后）
#### reset回滚
##### 硬reset回滚
先了解学习硬reset回滚的概念（千万别执行，不然上面的部分demo代码可能需要再敲一下）等看完回滚这章内容再尝试可以

硬回退会直接到对应的commit，并丢弃后续修改！！非常暴力
``` bash
git reset --hard 你想回到的hash
cat 123.txt  # 或者直接查看内容
```

- 强行覆盖远程`git push -f origin main` 
- 强制回到某个commit：`git reset --hard commit-hash`
- 强制回到上一个commit：`git reset --hard HEAD~1`
- 为什么已推送的分支避免用 `reset`？（后面push的话需 `git push -f`强制推送，会破坏协作!!!）自己不合作的项目本地开发用就行

##### 软reset回滚
先来学习软reset回退
``` bash
git reset --soft 增加了456的commit-hash
```
这会将 HEAD（可以理解成你的“工作指针”） 指向 对应增加了456的commit-hash，但修改内容仍保留在工作区和暂存区，简单说就是代码内容是不变的

这时候我们执行 `git log` 我们会发现只有增加了456的commit和之前的commit，后面的commit看不到了
``` bash
git reflog  # 可以找到对应Revert "好像有bug"commit的hash数值
# 按q退出查看
git reset --hard 对应的hash # 这里直接硬回滚回去
```

#### 回滚特定文件
只回滚特定的文件内容
``` bash
git log # 查看你的commit hash
git checkout 增加了456的commit-hash -- 123.txt
或者 git restore --source=增加了456的commit-hash -- 123.txt
cat 123.txt #查看变化，发现到当时的123 第二行456了

```
- 如果想保存这个更改，就`git add .` 然后提交 `git commit -m "Restore server.js to 增加456"`  
- 如果不想就一样的方法改回去就行  git log 或者reflog等查看hash

回滚这部分就学习完了，你可以尝试硬回滚力

#### 回滚问题巩固

- git reset 和 git revert 有什么区别？
	git reset 会修改提交历史，将 HEAD 回退到指定提交，后续提交会被删除（可以用 --soft 或 --hard 控制是否保留修改）。适合本地操作。git revert 则是创建一个新提交来撤销指定提交的内容，不改变历史，适合已推送远程仓库的场景。

- 如果误删了一个提交，如何恢复？
	如果用了 git reset，可以用 git reflog 查看历史操作,输出会显示所有 HEAD 的移动记录,找到误删的提交，如（c2f3a12）。`git reset --hard c2f3a12`

- 如何安全地回滚已推送到远程的代码？
	使用 git revert 创建一个新提交来撤销修改，然后正常推送。避免使用 git reset --hard 后强制推送（git push --force），因为这会改变历史，可能导致其他协作者出现冲突。

初步学习到这里结束力
## 动态案例语法库
这部分主要针对一些特定的问题，如果单独融入案例中可能比较麻烦（也受限于笔者的知识需要扩展）定期整理到核心概念和操作中变为具体案例

### 本地工作流

- 一开始建库后的一些问题 使用`rm -rf .git`  删除当前仓库的git，重新init
### 分支操作
基于origin/week创建my-week1用于学习别人的代码  ：`git checkout -b my-week1 origin/week1`  
删除远程分支：  `git push origin --delete branch_name` 
本地分支推送到远程 ：`git push origin my - week2`
显示所有分支：`git branch -a`
显示远程分支：`git branch -r`
舍弃未提交的更改：`git restore .`
### 合作修改
这部分主要问题是“解决冲突”
#### clone，fetch，pull区别

| ​**​命令​**​  | ​**​适用场景​**​ | ​**​是否修改工作目录​**​ | ​**​是否自动合并​**​ | ​**​历史记录影响​**​ |
| ----------- | ------------ | ---------------- | -------------- | -------------- |
| `git clone` | 首次下载远程仓库     | 是（创建新目录）         | 不适用            | 完整复制远程历史       |
| `git fetch` | 安全获取远程更新     | 否                | 否              | 仅更新远程分支引用      |
| `git pull`  | 快速同步并合并远程变更  | 是                | 是              | 可能生成合并提交       |
pull相当于fetch后merge  （记得切换到你想要修改的分支上） 如果没有冲突重叠的文件就不会报错

#### rebase
一般合并类主要就是：个人的分支合并，和别人合作时候和远程合并
个人分支整理可使用 `rebase` 替代 `merge` 以保持线性历
```bash
# 1. 开始变基（例如将 feature 分支变基到 main）
git checkout feature
git rebase main

# 2. 遇到冲突时，手动解决冲突文件
git add <冲突文件>  # 标记冲突已解决

# 3. 继续变基（仅处理当前分支的剩余提交）
git rebase --continue

# 4. 重复步骤2-3直到所有提交应用完毕
```
终止rebase：`git rebase abort`


可视化理解rebase：
原始状态：
```
远程: A --- B --- C
本地: A --- B --- D --- E
```
执行 `git pull --rebase origin main` 后：
```
远程: A --- B --- C
本地: A --- B --- C --- D' --- E'
```

和merge对比：​

|​**​对比项​**​|`git merge` (合并)|`git pull --rebase` (变基)|
|---|---|---|
|​**​冲突触发时机​**​|一次性合并所有冲突（`C` 和 `ED` 的冲突一起解决）|逐个提交解决冲突（先解决 `D` 和 `C` 的冲突，再解决 `E` 和 `C` 的冲突）|
|​**​解决后操作​**​|`git add` + `git commit`（生成合并提交）|`git add` + `git rebase --continue`（不生成新提交）|
|​**​历史记录​**​|保留分叉和合并提交（非线性历史）|线性历史，本地提交哈希值改变|


#### merge
生成一个​**​新的合并提交​**​（Merge Commit），将远程分支（也可以是本地分支）和本地分支的修改合并到一起。
- 在main上git merge main2 和在main2上git merge main 是不是没有什么区别：
	git merge 的执行顺序（即在 main 上合并 main2，或在 main2 上合并 main）会导致不同的结果
特性：
- `main` 通常是稳定分支，合并到它意味着发布新版本。
- `main2` 是开发分支，合并主分支是为了同步更新
结论：
- 若要将开发代码发布到主分支，用 `git checkout main && git merge main2`。
- 若自己开发，需同步主分支到开发分支，用 `git checkout main2 && git merge main`

先合并再删除
```bash
# 1. 切换到 main2 分支的合并目标分支（如 main）
git checkout main
# 2. 将 main2 合并到当前分支
git merge main2
# 3. 确认合并后，安全删除 main2
git branch -d main2
```
强制删除
```bash
git branch -D main2  # 注意是大写的 -D
```

取消merge：`git merge --abort`




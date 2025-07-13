个人的python爬虫学些笔记记录～
仅供参考，非盈利

参考：
[[../../../资源管理和软件使用/记录工具和用法/IT工具使用#anaconda 虚拟环境管理]]
比较有名的网站资源（Python3Web）：
个人按照学了部分，不过主要是学的b站（算是简单入门）
[崔庆才](https://scrape.center/)    [练习网站]( https://cuiqingcai.com/9522.html  )  [书本源代码]( https://github.com/Python3WebSpider)

b站教程：
[一个还可以的爬虫](https://www.bilibili.com/video/BV1Yh411o7Sz/?p=76&spm_id_from=333.1007.top_right_bar_window_history.content.click)该笔记也是基于这个学习过程中记录的～
## 基础预备知识
- 爬虫：编写程序模拟浏览器上网去互联网抓取数据
- 爬虫主要原则:
	不爬取公民个人隐私
	不断了别人财路
	不爬过多流量（导致别人宕机，侵入系统等） 利用time.sleep可以
- 是否违法：
	不违法但是窃取数据违法，爬虫干扰正常网页访问以及抓取受到法律保护的特定信息就不行力
	所以：优化自己的程序防止干扰网站     如果爬取发现涉及隐私和商业机密：停止
- 爬虫种类：
	通用型爬虫：不针对特定网站或特定主题的爬虫/抓取一整张网页内容
	聚焦爬虫：抓取页面局部内容/一整张页面中部分内容
	增量式爬虫：获取网站数据更新情况并爬取更新出来的数据
- robots.txt协议：
	君子协议。规定网站中哪些能爬哪些不能爬
	在后面加/robots.txt可以查看但不是所有网站都有都让 看allow和disallow

建议用annaconda环境管理，pycharm中的话File -> Settings -> Project -> Python Interpreter 选择对应解释器
## request模块
request基本取代了urllib，后者比较古老。
作用：模拟浏览器发请求
使用过程： （request编码流程）
	指定url（输入网址）
	发起请求
	获取响应数据（看到的东西）
	持续化存储响应数据

案例：简单的搜狗网站爬取
``` python
import requests

if __name__ == "__main__":
    #指定url
    url = 'https://www.sogou.com'
    #发送请求,同时接收get方法返回的数据
    response = requests.get(url = url)
    page_text = response.text
    #text表示以字符串的形式！response是get的实例化
    print(page_text)
    #保存，持久化
    with open('./sogou.html','w',encoding = 'utf-8')as fp:
        fp.write(page_text)
    print("任务结束")
```
右键html可以open in browser
open函数解释：
	open()是一个内置函数，用于打开一个文件。
	它接受多个参数（以下是主要参数）：
	第一个参数（'./sogou.html'）：这是要打开的文件的路径和名称。这里的./表示当前目录。sogou.html是你想要创建和写入的文件名称。
	第二个参数（'w'）：这是文件打开模式。'w'表示以写入模式打开文件。如果文件已存在，它将被清空；如果文件不存在，将创建一个新文件。
	第三个参数（encoding='utf-8'）：指定文件的编码格式。在这里，选择utf-8编码，以便可以正确处理中文和其他字符。

案例：网页采集器，读取器网页的搜索结果页面
查看UA方法：
	1、控制台中输入navigator.userAgent，直接显示
	2、选择网络选项，看query包的最下面
``` python
import requests

if __name__ == "__main__":
    #指定url
    url = 'https://www.sogou.com/web'
    #https://www.sogou.com/web?query=%E6%AC%B8%E5%98%BFd
    #网页本身后面的很多可以去掉，中文会编码，‘欸嘿’是一样的 ，可以去掉问号也

    kw = input('输出搜索关键词：')
    param = {
        'query' : kw
    }
    header = {
        'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0'
    }
    #params用于指定查询字符串参数，也就是URL中问号?后面的部分
    #其中还有filter等参数，和heder区分
    response = requests.get(url = url,params= param,headers= header)
    page_text = response.text
    fileName = kw+'.html'
    with open(fileName,'w',encoding='utf-8') as fp:
        fp.write(page_text)
    print(fileName,'保存成功')
```
爬取对应的代码
利用UA检测和UA伪装

案例：百度翻译结果实时展示
在网络中查看对应的结果
	[[Pasted image 20241223081528.png|F12]] 可以发现是一步步实时更新的d，do，dog  返回时json文件
	网页的动态更新通过ajax实现
```python
import json
import requests

if __name__ == "__main__":
    url = 'https://fanyi.baidu.com/sug'

    kw = input('输出搜索关键词：')
    data = {
        'kw':kw
    }
    #‘kw’其实是特定的键名要求，在网站的api文档中可以查看
    header = {
        'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0'
    }
    #利用requests模块中的post
    # response = requests.post(url=url,data=data,headers= header)
        # get = response.text   检查结果
        # print(get)
    response = requests.post(url=url, data=data, headers=header)
    dic_obj = response.json()
    #确认是json文件所以这么存
    fileName = kw+'.json'
    with open(fileName,'w',encoding='utf-8') as fp:
        json.dump(dic_obj,fp=fp,ensure_ascii  =False)
        #因为我们拿到的这个dic_obj是中文的，中文是不能用ascii的
        # json.dump 比 fp.write 更适合在写入 JSON 对象时使用，因为它自动处理序列化，确保数据格式正确。而 fp.write 只能写入字符串，不支持直接将字典等复杂数据结构保存为预期的 JSON 格式。
        #以原格式写入，相当于
    print(fileName,'保存成功')

``` 
代码
存在问题：
	1、api文档不知道怎么找和看（比如kw就是api中限定的修饰词）
	2、在https://fanyi.baidu.com/mtpe-individual/multimodal?query=dog&lang=en2zh
	中似乎是以query方式查询的，返回的结果在“tranlate”中格式是text，不知道怎么同时读取多结果
	3、改的url网址虽然有结果，但是访问时候error？

案例：豆瓣电影 排行榜 喜剧
	数据解析方法可以但是不用，尝试请求json形式数据
	滚轮拖到中间就会多加载一批新的数据，说明有ajax请求，还有就是看地址有没有变！
		是get类型，看query parameters（或者说“负载”）中有需要上传到服务器的参数
	“？”后面的参数需要封装，后面部分在url中去掉
	网上的json解释器可以看下载下来的json文件格式
``` python
import json
import requests

if __name__ == "__main__":
    url = 'https://movie.douban.com/j/chart/top_list'

    header = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0'
    }
    param = {
        #每个键值对之间需要用逗号分隔
        'type': '24',
        'interval_id': '100:90',
        'action':'',
        'start':'0', #从库中第几部电影取
        'limit':'20'  #一次取出多少 个
        }

    response = requests.get(url=url, params=param, headers=header)
    #ajax，通过查看content-type确定用json
    list_obj = response.json()
    #返回的确实也是个列表
    fileName = '豆瓣电影' + '.json'
    with open(fileName, 'w', encoding='utf-8') as fp:
        json.dump(list_obj, fp=fp, ensure_ascii=False)
        #fp表示将对象写入fp中
    print(fileName, '保存成功')

```
代码
运用了json文件（XHR）ajax

案例：化妆品生成许可证（药监局）
	判断页面中的信息是不是通过get方法直接.text就可以获取的（也可以在网络页面中抓包all直接刷新然后查看response部分的结果中有没有想要的信息内容，Ctrl+F）
	如果想要的信息不是就考虑下动态刷新出来的（别的请求方式）比如ajax的可能性  所以查看XHR部分看抓包结果
	ajax中发现有数据，json文件形式，只有企业的名字，和id，没有点开企业链接后的更多内容。
	对企业数据详情页的url页面观察：
		每个前面的网址是一样的，只是请求的id不一样->域名和id拼接而成
	详情页的企业数据信息也是动态获得的（直接看response没有查到结果）
		同样通过看XHR，刷新获取ajax动态更新的数据包分析
	对比不同id号的备案信息的网址，发现是一样的，就是id不一样 （也看到是post请求类型）
list[{ID:},{ID:},{ID:}]这样的结构：
	for dic in list_json['list'] list本身也是键名 然后 dic['ID']
``` python
import json
import requests

if __name__ == "__main__":
    url = 'xxx'#网站首页，为获取id
    header = {
        'User-Agent': 'F12的request中查看看自己的'
    }
    param = {
        }
    id_list = []
    all_data = []
    #post方法对应的参数是data
    json_ids = requests.post(url = url,headers= header,data=data).json()
    for dic in json_ids['list']:
        id_list.append(dic['ID'])
    #选取特定的id列
    #进一步获取企业的详细数据
    post_url = '' #共有的url头
    #还是以参数的形式传入，因为是补充在后面的
    for id in id_list:
        data2 = {
            'id':id
        }
        detail_json = requests.post(url = post_url,headers= header,data = data2).json()
        all_data.append(detail_json)
    #持久化存储
    with open('./alldata.json','w',encoding='uft-8') as fp:
        #./指出在当前目录下
        json.dump(all_data, fp=fp, ensure_ascii=False)
    print( '保存成功')
```

进一步：获取更多页面下的备案信息，需要做到页面刷新：
	在get的参数data中，有个键：page
		for page in range（1，6）：
			page = str(page)   因为传入的是字符串类型

#### 辨析request保存形式
text：
	返回的是响应内容的字符串类型，是对响应内容进行解码后的结果。requests会根据响应头中的Content-Type自动选择合适的编码方式进行解码（默认是 UTF-8）。
	使用response.text获取到的内容通常用于处理网页的HTML等文本数据，方便后续的文本操作（如使用 lxml、BeautifulSoup 进行解析）。
content：
	返回的是响应内容的原始字节格式（bytes），不进行任何解码。这对于处理二进制数据（如图片、PDF、RAR文件等）非常有用。
	使用response.content获取到的内容适用于保存文件或处理非文本数据。
.json:
	作用: 获取响应内容并将其转换为 Python 字典或列表（根据响应的JSON结构）。前提是响应的 Content-Type 必须为 application/json。
	使用场景: 当你请求一个返回 JSON 格式的数据的API时，使用 .json() 方法会更加方便，因为它会自动处理JSON解析的细节。


## 数据解析
方式：正则，bs4，xpath（学习重点）
解析原理：
	需要解析的局部文本内容在标签之间或者标签对应属性中存储
		1、进行标签的定位
		2、标签或者标签属性中存储的数据值进行提取分析
#### 正则表达式
python 正则表达式语法总结：
正则表达式常用语法总结
- 基本符号
	\[aeo\]: 匹配字符集中任意一个字符。
	\d: 匹配数字 [0-9]。
	\D: 匹配非数字。
	\w: 匹配字母、数字和下划线。
	\W: 匹配非字母、数字和下划线。
	\s: 匹配任意空白字符（如空格、换行、制表符等）。
	\S: 匹配非空白字符。
	.  :用于匹配除换行符（如 \n）以外的所有单个字符
- 量词:（必须跟在特定的字符后面！！！）
	\*: 匹配前一个表达式零次或多次。
	+: 匹配前一个表达式一次或多次。
	?: 匹配前一个表达式零次或一次。
	{m}: 匹配前一个表达式 m 次。
	{m,n}: 匹配前一个表达式至少 m 次，至多 n 次。
- 边界:
	$: 匹配行的结束。
	^: 匹配行的开头。
- 分组:
	(ab): 匹配 ab，并创建一个捕获组。
- 常用方法：
	import re
	e.I: 忽略大小写
	re.M: 多行模式
	re.S: 点匹配模式
	# 替换
	re.sub(正则表达式, 替换内容, 字符串)

``` python
import re
key = 'javapythonc++php'
result = re.findall('python',key)[0]
print(result)

key = '<html><h1>hello<h1><html/>'
re.findall('<h1>(.*)<h1>',key)[0]

string = '我喜欢身高170的女3孩4'
result =re.findall('\d+',string)
print(result)
#如果是/d*那么汉字其实也会匹配进去，因为是0次可以

key = 'http://www.baidu.com and https://boob.com'
re.findall('http?://',key)

key = 'lalala<hTml>hello</HtMl>hahah'
result =re.findall('<[Hh][Tt][mM][Ll]>(.*)</[Hh][Tt][mM][Ll]>',key)
print(result)
#正则表达式返回()中的所有内容

key = 'bobo@hit.edu.cn.hit'
result=re.findall('h.*?\.',key)
#如果没有？，就会为hit.edu.cn.
#\.是转义符
print(result)

key = 'saas and asa and saaas and sas'
result = re.findall('sa{1,2}s',key)
#限定次数
print(result)

```
正则表达式训练代码

案例：图片爬取
``` python
import requests

if __name__ == "__main__":
    url = 'https://tiebapic.baidu.com/forum/crop%3D0%2C21%2C300%2C210%3Bwh%3D150%2C105%3B/sign=5932f51b786d55fbd1892c6650126378/a144ad345982b2b7af07a00174adcbef77099b5f.jpg?tbpicau=2025-01-06-05_333eb3b354b5c81849a2c6f0f7b0b120'
    image_data = requests.get(url = url).content
    #图片返回二进制对象,利用get请求页面其实就是图片
    with open('./dog.jpg','wb') as fp:
        #二进制文件不用encoding = 'utf-8'
        #wb专门针对二进制文件的写入（不单纯0和1）
        fp.write(image_data)
    print('爬取完毕')
```
代码
右键“检查”，源码中找到链接右键打开

案例：爬取一个页面中多张图片（以及多个页面）

``` python
import requests
import re
import os

if __name__ == "__main__":
    #创建一个文件夹，保存图片
    if not os.path.exists('./贴吧图片'):
        os.mkdir('./贴吧图片')
        
    url = 'https://tieba.baidu.com/p/7681741431'
    header = {
        'User-Agent': '放自己的'
    }
    page_text = requests.get(url = url,headers = header).text
    #先爬取页面  # print(page_text)
    #使用聚焦爬虫
    # ex = '<img class="BDE_Image".*?src="(.*?)"'
    ex = '<img class="BDE_Image" src="(.*?)" size='
    image_src_list = re.findall(ex,page_text,re.S)
    #re.S:使得*也可以匹配/n   (原本不可以)       #非贪婪一次匹配一个
    print(image_src_list)
    for src in image_src_list:
        image_data = requests.get(url= src,headers= header).content
        #生成图片名称-》从图片链接特征中找或者自定义
        filename = src.split('/')[-1].split('?')[0]   #按照顺序来split理解
        imgpath = './贴吧图片/' + filename
        print(imgpath)
        with open(imgpath,'wb') as fp:
            fp.write(image_data)
            print(filename,"打印成功")
            #利用，分开
#复制网页中的源码
   # <div id="post_content_142720568915" class="d_post_content j_d_post_content " style="display:;">                    查看原图→保存到相册<br><img class="BDE_Image" src="http://tiebapic.baidu.com/forum/w%3D580/sign=34a2f8033d59252da3171d0c049b032c/7bd7ad6eddc451da6c9016f2ebfd5266d0163243.jpg?tbpicau=2025-01-07-05_cc05ba8ca3fa45e5453d28f3a9910a6d" size="250342" changedsize="true" width="560" height="560"></div>


```  
%% ENDREGION %%代码
F12后选择元素，再选方框鼠标（在浏览器中选择元素进行检查）可以识别浏览器中图片自动对应源代码
怎么批量发送img中链接请求：
	寻找图片所在位置的规律，比如都在\<div>的
%% REGION %% 
``` python
import requests
import re
import os

if __name__ == "__main__":
    #创建一个文件夹，保存图片
    if not os.path.exists('./贴吧图片'):
        os.mkdir('./贴吧图片')
    #header用一次就行，放外面
    header = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0'
    }
    #不同页面url：
        #https://tieba.baidu.com/p/7681741431?pn=2
        #https://tieba.baidu.com/p/7681741431?pn=3   然后确实？pn=1也可以到第一张页面
    #设置一个通用url模板
    url = 'https://tieba.baidu.com/p/7681741431?pn=%d'
    for pagenum in range (1,10):
        new_url = format(url%pagenum)
        page_text = requests.get(url = new_url,headers = header).text
        # ex = '<img class="BDE_Image".*?src="(.*?)"'
        ex = '<img class="BDE_Image" src="(.*?)" size='
        image_src_list = re.findall(ex,page_text,re.S)
        # print(image_src_list)
        if not os.path.exists(f'./贴吧图片/第{pagenum}页图片'):
            os.mkdir(f'./贴吧图片/第{pagenum}页图片')
        for src in image_src_list:
            image_data = requests.get(url= src,headers= header).content
            #生成图片名称-》从图片链接特征中找或者自定义
            filename = src.split('/')[-1].split('?')[0]   #按照顺序来
            imgpath = f'./贴吧图片/第{pagenum}页图片/' + filename
            #图片后面的/一定要注意！！到这个文件夹下面寻找的意思
            #imgpath = os.path.join(page_dir, filename) 也可以，
            with open(imgpath,'wb') as fp:
                fp.write(image_data)
                print(filename,"打印成功")
                #利用，分开
    print('打印完成！！')
``` 
%% ENDREGION %%代码：循环打印页面找不同图片
发现其实贴吧上的page1包含的page1和page2的内容，所以之前才会打印不全

#### bs4（python中独有的数据解析）
原理：
	1.实例化一个BeautifulSoup对象，并且将页面源码加载到该对象中
	2.通过BeautifulSoup对象中相关属性或者方法进行标签定位和数据提取
安装库：
	bs4，lmxl
from bs4 import BeautifulSoup   对象实例化
	将本地html文档加载到该对象中 
	将互联网上获取的页面源码加载到对象中
	
%% REGION %% 
``` python
import requests
from bs4 import BeautifulSoup
import lxml
if __name__ == "__main__":
    #本地实例化
        #fp  = open('./xxx.html','r',encoding='utf-8')
        #soup = BeautifulSoup(fp,'lxml')
        #lxml是字符串参数，所以需要''传入
    #互联网实例化
        #page_text = response（请求）.text
        #soup = BeautifulSoup(page_text,'lxml')
    #soup.a: (soup.tagName 返回文档中第一次出现tagName的标签) = soup.find('a')
    #soup.find('div',class_='song'))
    #soup.find_all('a')  返回符合要求的所有标签
    #soup.select('.tang')   可以定位id或者class或者标签选择器为tang  返回为列表
    #层级选择器；
        #soup.select('.tang > ul > li > a')[0] 如果有多个a会都返回了,再取第一个标签 
        #soup.select() 返回的列表中的每个元素都是 BeautifulSoup 的 Tag 对象（是 BeautifulSoup 的一种实例化对象）。 这意味着你可以对返回列表中的每个元素继续使用 BeautifulSoup 的方法和属性
        #soup.select('.tang > ul > a')[0]  中间空一格表示多个层级！！！也可以
    #获取标签之间文本数据
        #soup.a.text/string/get_text()   两个属性，一个方法 获取文本内容 text和get_text()可以获取所有文本内容（包括子系），string只能是该标签下的
        #比如当一个div下很多p的时候，获取很多内容，就用div 的text
    #获取标签中属性值：
        #soup.a['href']
```
%% ENDREGION %%代码，解释bs4使用

%% REGION %% 
```python
import requests
from bs4 import BeautifulSoup
import lxml

if __name__ == "__main__":
    header = {
        'User-Agent': '查看自己的'
    }
    url ='https://www.bqzw789.org/569/569088/'
    chapter_page = requests.get(url = url,headers=header).text
    ##其实处理text或者content似乎都行,传入soup
    #第一次遇到了反爬虫的，会返回一堆javascript代码
    # with open('./chuyin.html','w',encoding='utf-8') as fp:
    #     fp.write(chapter_page)
    # print('ok')
    soup = BeautifulSoup(chapter_page,'lxml')
    # a_list = soup.select('#list > dl > dd > a#list ')
    a_list= soup.find_all('a',id = 'list')  #两种效果确实一样
    #出现了多id的情况，这样指引好一些！   .是class  #才是id
    print(a_list)
    fp = open('./神仙小说.txt','w',encoding='utf-8')
    for a in a_list:
        title = a.string
        # detail_site = 'https://www.bqzw789.org/' + a['href']
        detail_site  = 'https://www.bqzw789.org' + a['href']
        detail_page = requests.get(url = detail_site,headers=header).text
        detail_soup = BeautifulSoup(detail_page,'lxml')
        div_tag = detail_soup.find('div',id = 'content')  #只用第一个就可以
        #find返回第一个对象，findall返回列表，select返回列表
        #div_tags = detail_soup.find_all('div', id='content')和 div_tags = detail_soup.select('div#content') 效果一样！！
        content = div_tag.text
        #里面全部文字都搞下来
        fp.write(title+':'+'\n'+content+'\n')
        print(title,' 保存完毕')
    #4000多章节，赶紧不敢爬了
```
案例：利用bs4爬取网络小说
熟练find，find_all,和select 的相关语法，容易混~注意返回的形式

#### Xpath
- 原理：
	实例化一个etree对象，将被解析的页面源码加载
	调用etree对象中xpath方法结合xpath表达式实现内容定位等
	也需要lxml
- 实例化：
	本地实例化：
		etree.parse(filePath)
	互联网实例化
		etree.HTML('page_text')
		-xpath('xpath表达式')

%% REGION %% 
``` python
from lxml import etree
#包导入，不然lxml.etree了

if __name__ == "__main__":
    tree = etree.parse('./test.html')
    r = tree.xpath('//head//text()')

#gpt给出改进：
    # if __name__ == "__main__":
    #     # 使用 lxml.html 解析 HTML 文件
    #     with open('./test.html', 'r', encoding='utf-8') as f:
    #         content = f.read()  # 读取文件内容
    #     tree = html.fromstring(content)  # 解析内容
    #     r = tree.xpath('//head//text()')  # 执行 XPath 查询
    #     print(r)  # 打印结果

# /:从根开始遍历定位 //：多个层次中间  不要忘记./的. !!!!!!!
#属性定位： //div[@class = 'song']  表示匹配对应属性值
#多个相同标签，定位从1开始，下面li为例
#r = tree.xpath('//div[@class = 'tang']//li[5]/a/text()')[0]
#因为返回为字典，记得取位置
#r = tree.xpath('//div[@class = 'tang']/img/@src') 这样是取出属性值
#text()也会返回很多比如\n\t等，需要识别去除
```
讲解例子代码

案例：爬取58的信息 （遇到了需要动态访问）

``` python
import requests
from bs4 import BeautifulSoup
from lxml import etree
#包导入，不然lxml.etree了

if __name__ == "__main__":
    header = {
        'User-Agent': '查看自己的'
    }
    url ='https://cn.58.com/ershoufang/k1'
    page_text = requests.get(url = url,headers= header).text
    # print(page_text)
    tree = etree.HTML(page_text)
    title_list = tree.xpath('//div[@class = "property-content-title"]/h3/text() ')
    # print(title_list) #会出现很多元素
    # content_list = tree.xpath('//div[@class = "property-content-info"]/text()')
    # content_class_list = tree.xpath('//div[@class = "property-content-info property-content-info-comm"]/text()')
    # print(len(title_list))    #检查对应长度
    # print(len(content_list))
    # print(len(content_class_list))
    # print(content_class_list) #发现结果是空的，可能和动态获取有关！！！！！后面再学
    fp = open('./title地址信息.txt','w',encoding='utf-8')
    # #若需要二层聚焦，需要在上面列表中的“元素” 后比如：title = li.xpath('./div[2]/h2/text()')[0]
    # 解释：在li的xpath基础上添加，再取出第一个h2的内容

    # for title, content, content_class in zip(title_list, content_list, content_class_list)
    #这样遍历会导致问题因为如果content同时增加2的话，步调不一样，改进：用一个i去取元素索引
    #     for i in range(len(title_list)):
    #         title = title_list[i].strip()
    #         # 计算对应的 content 索引
    #         content_index = i * 2
    #         # 确保不存在越界
    #         if content_index < len(content_list):
    #             content1 = content_list[content_index].strip()
    #         else:
    #             content1 = '无详细信息'
    #         if content_index + 1 < len(content_list):
    #             content2 = content_list[content_index + 1].strip()
    #         else:
    #             content2 = '无详细信息'
    #         # 获取对应的 content_class, 假设每个 content_class 与第一个 content 对应
    #         class_index = content_index // 2
    #         content_class = content_class_list[class_index].strip() if class_index < len(content_class_list) else '无附加信息'
    #         # 打印信息
    #         print(f"地址信息: {title}")
    #         print(f"详细信息1: {content1}")
    #         print(f"详细信息2: {content2}")
    #         print(f"附加信息: {content_class}")
    #         print("---")  # 用于分隔每项
    for title in title_list:
    # #别和字典的遍历混淆了
        fp.write('地址信息' + ':' + title.strip() + '\n' )
    #     fp.write('详细信息' + ':' + content.strip() + '\n')
    #     fp.write('\n')
    print('保存房产title完毕')

```
代码

```python
import requests
from lxml import etree
import os
if __name__ == "__main__":
    header = {
        'User-Agent': '查看自己的'
    }
    url ='https://wallspic.com/cn/tag/chu_yin_wei_lai/3840x2160'
    response = requests.get(url=url, headers=header)
    print(response.status_code)  # 输出状态码
    if response.status_code == 200:
        page_text = response.text
    else:
        print("请求失败，状态码:", response.status_code)
    tree = etree.HTML(page_text)
    div_list = tree.xpath('//div[@class = "gallery_fluid"]')
    if not os.path.exists('./piclib'):
        os.mkdir('./piclib')
    #假设是由很多个class一样的div中有内容
    for div in div_list:
        img_src = 'https://xxx' + div.xpath('./a/img/@src')[0]
        img_name =  div.xpath('./a/img/@alt')[0] +'.jpg'
        #一般alt中存放名字
        img_name = img_name.encode('iso-8859-1').decode('gbk')
        #解决中文乱码的问题
        image_data = requests.get(url = img_src,headers=header).content
        img_path = './piclib/'+img_name
        with open (img_path,'wb') as fp:
            fp.write(image_data)
            print(img_name,'下载成功')

```
案例：爬取图片
存在比如动态渲染使用reque捕捉不到等问题！！图片网站还是难找

案例：爬取城市信息
```
#//div[@class ='a']/ul/li/a
#//div[@class ='a']/ul/div[2]/li/a
#可以xpath('第一个 | 第二个')   这样同时锁定两个
#使用//div/ul//li/a也可能可以但是怕引入别的地方还有的问题
```
理论相同，处理同时爬取两个路径信息方法

https://sc.chinaz.com/  站长素材网，用于练习似乎相当不错！免费因为

``` python
import requests
from lxml import etree
import os
if __name__ == "__main__":
    header = {
        'User-Agent': '查看自己的'
    }
    base_url ='https://sc.chinaz.com/jianli/daxuesheng'
    #https://sc.chinaz.com/jianli/daxuesheng_3.html
    # 多个网网的网站形式不统一,_1不能访问，同时没有?page=1这样的params形式能输入，所以get不能params
    count = 0
    if not os.path.exists('./简历'):
        os.mkdir('./简历')
    for page_number in range(1, 4):  # 假设你想爬取1到3页
        if page_number == 1:
            full_url = base_url + '.html'  # 第一页
        else:
            full_url = f'{base_url}_{page_number}.html'  # 从第二页开始
        response = requests.get(url=full_url, headers=header)
        response.encoding ='utf-8'  #尝试解决中文编码错误问题,成功
        print(response.status_code)  # 输出状态码
        #状态码检查！！！！！！！！！！！！！！！
        if response.status_code == 200:
            page_text = response.text
        else:
            print("请求失败，状态码:", response.status_code)
            
        tree = etree.HTML(page_text)
        # print(page_text)
        a_list = tree.xpath('//a[@class = "title_wl"]')
        # print(len(a_list))  #存在分页和实际展示不一样问题,确实，在谷歌浏览器中展示的更好，40个5*8
        # print(a_list[0])   #用于检测是否存好了
        if not os.path.exists(f'./简历/第{page_number}页简历'):
            os.mkdir(f'./简历/第{page_number}页简历')
        #这里只打开文件夹就行应该
        for a in a_list:       #已经是完整的url了
            a_src = a.xpath('./@href')[0]   #记得[0]选出元素！！
            # print(a_src + '\n')
            a_name = a.xpath('./text()')[0]
            # a_name = a_name.encode('iso-8859-1').decode('gbk')    
            #也尝试这样调整下文字，不太行，这个和网页编码有关，不细了解先
            detail_page = requests.get(url= a_src,headers=header).text
            download_tree = etree.HTML(detail_page)
            download_url = download_tree.xpath('//div[@class = "down_wrap"]//ul[@class = "clearfix"]/li/a/@href')[0]
            print(download_url)  #检查下地址
            #直接从福建电信处下载,对于rar类型的，直接获取response就行
            rar_response = requests.get(url=download_url, headers=header)
            fp = open(f'./简历/第{page_number}页简历/{a_name}.rar','wb')
            fp.write(rar_response.content)
            print(f'{a_name} 保存完毕')
            count += 1
            if count == 3:
                count =0
                break #跳出当前循环，不能continue！
```
爬取多页简历代码！！！！集大成回顾
注意对于中文的编码解决在代码中的体现，两种方法！
状态码检查！！！！！200，查看检查结果



## 验证码识别
- 反爬机制：识别验证码中的数据
- 使用第三方识别图片！！这个过程方法就比较多，不唯一，不钻
	tesserocr，超级鹰，图鉴，tesseract
	注意导入库，以及对应的函数修改。
- 获取网站上的验证码的图片的链接地址一般需要，保存到本地（地址加xpath）

#### 模拟登录思路
- 思路：
	输错几次后可能弹出验证码直接输入的
	打开F12看正确登录的结果的抓包，比如login？相关的看具体数据（url，方式post，参数）
- 流程：
	1.验证码识别，获取验证码文字数据（结合xpath，每次更新每次获取）
		//\*\[@id = "verify xxx"] 这样直接匹配也可以的，用\*
		一般用.content 接收  
	2.对post发送请求（携带参数）
		登录得到的是网站响应数据
	3.响应数据持久化处理
	可以通过响应状态码来看，是200就是成功，上一个案例的代码中有涉及！！

#### 爬取个人信息（cookie）：
登录之后操作
	接着通过打开用户具体的个人信息网站（一般是不会变的）爬取网页信息。但是发现又会重新需要登录（因为不知道这是基于登陆状态的请求（http/https协议的”无状态性“）），这时就和cookie有关了
cookie：
	让浏览器记住客户端状态的！！
	可以通过截取网络包分析对应在信息页的post请求包（看名字找一般，‘profile’）的响应头中的cookie
自动处理cookie：（手动式通过抓包工具进行分析，不太行不适用）！！！！！！
	cookie值来源是？（具体分析网络包）
		在服务端登录（login数据包）的响应头的set-cookie中携带，由服务器创建
		session会话对象：
			作用：可以进行请求发送，如果过程中产生了cookie就会自动放到session中
	**所以使用session对象（session = requests.Session(),接着后面就是session.get(...)）进行模拟登录的请求发送，cookie机会携带在session中，再使用session对个人主页对应发送get请求**

#### 代理
某个ip单位时间内访问过多可能被封，所以通过代理去破解封IP的反爬机制
- 通过代理服务器突破访问限制
	可以隐藏自己真实的ip
- 代理相关网站：
	快代理  https://www.kuaidaili.com/free/intr    更新不及时
	西祠代理 
	www.goubanjia.com  被关闭了
- 在百度中输入ip就能查到自己在公网上的ip地址了（不过静态似乎不太行）
	和本地的Wlan 的ip地址不一样，Wlan中的时内网（局域网 IP）
	内网ping不了公网可能是和网络运营商会禁止 ping 公网 IP以及NAT相关有关，可以ping域名
- 这个网站有个功能可以检测代理ip是否能链接，很多情况下应该是因为失效了所以自己访问的还是自己的ip
	https://www.89ip.cn/check.html 
- 这个网站能看自己的ip
	https://myexternalip.com/ 
- 代理ip的匿名度：
	透明：服务器知道使用了代理并知道真ip
	匿名：知道用了代理，不知道真实ip
	高匿：不知道代理和ip
```python
import requests
from lxml import etree
import os
if __name__ == "__main__":
    header = {
        'User-Agent': '自己的'
    }
    # url ='https://www.baidu.com/s?wd=ip'
    url = 'https://myexternalip.com/'
    proxy = {
        'http':'http://60.188.5.232:80'
    }    #这样也行
    response = requests.get(url=url,proxies={"https":'114.231.41.101:8089'}).text
    with open ('./ip.html','w',encoding='utf-8') as fp:
        fp.write(response)
``` 
代理代码


## 高性能异步爬虫
目的：在爬虫中使用异步进行高性能数据爬取
	比如多个网址线性get访问的时候容易出现get堵塞（耗费时间很多）的问题
		不过也不建议为所有的都开一个线程。
多线（进）程：会一直创建，销毁  （不太推荐）
线程池：高效复用，可以短期执行大量任务

``` python
import time
from multiprocessing import Pool

start_time = time.time()

def get_page(item):
#不要str，和定义重名
    print('正在下载', item)
    time.sleep(2)
    print('下载完毕', item)

if __name__ == '__main__':
	#一定要放在main函数下不然会有冲突！！导致报错(保护机制相关)
    data_list = ['a', 'b', 'c', 'd']
    pool = Pool(4)
    pool.map(get_page, data_list)
    #map传入参数需要符合格式，比如如果是dic那么上面函数定义就是要dic/dictionary

    # pool.close()  # 关闭进程池，不再接受新的任务
    # pool.join()   # 等待所有子进程结束

    end_time = time.time()
    print(end_time - start_time)
```
案例：理解Pool

response中查不到\<vedio但是能有mp4的结果但是可能在\<script>中，这时用xpath和bs4就不能查询了（不是html） 可能用到selenium   这时就需要正则来提取
同时存名字和链接？可以创建一个字典，存放两个，再列表.append

案例：梨视频，使用了伪装的url！！
学习参考：
	https://www.cnblogs.com/industrial-fd-2019/p/14413568.html （主要）
	https://www.cnblogs.com/atangaba/articles/15851270.html
	down_url=dic_obj\['videoInfo']\['videos']\['srcUrl']   当多个键值嵌套的时候，一层层深入
	自己一步步发现的确实也是一样的
	findall可以和xpath结合    replace函数去替换伪地址中

``` python
import requests
from lxml import etree
import os
from multiprocessing import Pool

if __name__ == "__main__":
    header = {
        'User-Agent': '填入自己的'
    }
    url = 'https://www.pearvideo.com/'
    response = requests.get(url=url,headers=header).text
    # print(response.status_code)  # 输出状态码
    # if response.status_code == 200:
    #     page_text = response.text
    # else:
    #     print("请求失败，状态码:", response.status_code)
    tree = etree.HTML(response)
    li_list = tree.xpath('//div[@class = "vervideo-tlist-bd recommend-btbg clearfix"]//div[@class = "vervideo-tbd"]')  #别忘记“”
    #print(len(video_list))
    for li in li_list:
        detail_url = 'https://www.pearvideo.com/' + li.xpath('./a/@href')[0]
            # print(detail_url)
        name = li.xpath('./a//div[@class="vervideo-name"]//text()')[0] +'.mp4'
            # print(name)
        detail_page = requests.get(url = detail_url,headers=header).
            # print(detail_page.status_code)  # 输出状态码
            # if detail_page.status_code == 200:
            #     page_text = detail_page.text
            # else:
            #     print("请求失败，状态码:", detail_page.status_code)
        #https://video.pearvideo.com/mp4/short/20241106/cont-1797087-16040060-hd.mp4 直接在cont中有个包
``` 
未完成代码：

**单线程和异步协程（推荐）**

``` python
import asyncio
async def request(url):  #定义异步函数
    print("访问",url)
    # await asyncio.sleep(1)
    print("访问成功",url)
    return url
c = request("www.baidu.com")

def callback_func(task):
    print("Callback Result:", task.result())  # 打印任务结果

if __name__ == "__main__":
    #futrue和task没有什么区别，利用async定义一个协程，利用await挂起阻塞方法的执行
    loop = asyncio.get_event_loop()
    #需要在创建任务对象,不能直接loop去跑
    task = loop.create_task(c)
    # task = asyncio.ensure_future(c) #一样的
    task.add_done_callback(callback_func) 
    # print(task) #看task的执行状态
    loop.run_until_complete(task)
    # print(task)
```
协程理解示例代码（只在循环中注册了一个任务对象）


``` python
import asyncio
import time
async def request(url):  #定义异步函数
    print("访问",url)
    await asyncio.sleep(1)
    # time.sleep(1)
    # 后者相当于阻塞，两者不一样，await别的协程还可以同时执行（挂起）
    #异步模块中出现同步代码就无法实现异步
    print("访问成功",url)

if __name__ == "__main__":
    start = time.time()
    urls = [
        'www.baidu.com',
        'www.4399.com',
        'www.7k7k.com'
    ]
    #任务列表
    stacks = []
    for url in urls:
        c = request(url)
        task = asyncio.ensure_future(c)
        #就不需要loop
        stacks.append(task)
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.wait(stacks))
    #传入协程对象的参数就行，不用await （应该）
#loop.run_until_complete() 只能接受“一个”可等待对象,，两种方式
    # done, pending = await asyncio.wait(tasks)
    #results = await asyncio.gather(*tasks)
    print(time.time() - start)
    ``` 
案例：多任务异步协程实现代码**

#### aiohttp模块

``` python
import asyncio
import time
import re
import os
import requests
async def get_page(url):
    header = {
        'User-Agent': '填入自己的'
    }
    print("访问",url)
    await asyncio.sleep(1)
    response = requests.get(url = url,headers=header)
    if(response.status_code == 200):
        print("下载成功",response.text)
    else:
        print('fail')

if __name__ == "__main__":
    start = time.time()
    urls = [
        'https://www.baidu.com',
        'https://www.4399.com',
        'https://www.7k7k.com'
        #否则request没办法解析
    ]
    #任务列表
    stacks = []
    for url in urls:
        task = asyncio.ensure_future(get_page(url))
        #就不需要loop
        stacks.append(task)
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.wait(stacks))
    print("总耗时",time.time() - start)
    #如果用+后面要str(),字符串的拼接
```
 **request不行的示例，需要aiohttp**
其实应该搭建flask服务器自己设置time，在服务器延时，这里的request访问还是快的
需要安装pip aiohttp

``` python
import asyncio
import time
import aiohttp

async def get_page(url):
    async with aiohttp.ClientSession() as session:
    #创建ClientSession()的实例化对象
        async with session.get(url) as response:
            #这里不用await，没必要！
            #text()方法返回字符串形式响应数据
            #read()方法返回二进制形式响应数据
            #json()返回的就是json对象
            try:
                # 先尝试 utf-8
                page_text = await response.text(encoding='utf-8')
            except UnicodeDecodeError:
                # 如果失败，尝试 gb2312
                page_text = await response.text(encoding='gb2312')
                #4399使用了不同的编码方式，这个应该是在响应头看
            print(page_text)
            status_code = response.status  # 使用 status 属性
            # 打印状态码
            print("状态码:", status_code)
            if status_code == 200:
                print("下载成功")
            else:
                print('请求失败，状态码:', status_code)

if __name__ == "__main__":
    start = time.time()
    urls = [
        'https://www.baidu.com',
        'https://www.4399.com',
        'https://www.7k7k.com'
        #否则request没办法解析!!
    ]
    #任务列表
    stacks = []
    for url in urls:
        task = asyncio.ensure_future(get_page(url))
        #就不需要loop
        stacks.append(task)
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.wait(stacks))
    print("总耗时",time.time() - start)
    #如果用+后面要str(),字符串的拼接
``` 
**案例，多任务异步爬虫**

``` python
class Dog:
    def __init__(self, name):
        self.name = name  # 这是一个属性，用于存储狗的名字
    def speak(self):  # 这是一个方法
        return f"{self.name} says Woof!"
# 创建一个 Dog 对象
my_dog = Dog("Buddy")
# 访问属性
print(my_dog.name)  # 输出: Buddy
# 调用方法
print(my_dog.speak())  # 输出: Buddy says Woof!
```
理解.text和.text(): (在requests和aihttp中，分别是属性和方法)不过知道用法就行

- async 和 await 是 Python 语言本身的关键字，用于支持异步编程！！
- 理解await：
	在调用 async 定义的函数时需要使用 await
	在调用返回 coroutine （协程）对象的方法时需要使用 await，比如：await response.text()
	普通函数/方法不需要 await
	类似就是看返回的对象是不是需要协程，进一步取出，但是前面取出（对resposne，await）不代表后面用的response就不是协程了！
	
**Coroutine（协程）是 Python 异步编程的核心概念！！**
```python
# 普通函数 - 直接返回结果
def normal_function():
    return "Hello"

# 协程函数 - 返回一个协程对象
async def coroutine_function():
    return "Hello"

# 使用方式
result1 = normal_function()     # 直接得到 "Hello"
result2 = coroutine_function()  # 得到一个协程对象，需要 await
result3 = await coroutine_function()  # 得到 "Hello"
```
理解协程和await关系
- 打个比喻：
	普通函数像是直接给你结果
	协程像是给你一张"提货单"（协程对象）
	await 就是用"提货单"去取实际的结果

- 协程和线程的区别：
	# 线程池：每个线程都需要系统资源
	pool = ThreadPoolExecutor(max_workers=100)  # 创建100个线程
	# 协程：单线程内即可运行大量协程
	tasks = \[download_file(url) for url in range(100)]  # 创建100个协程任务

``` python
from concurrent.futures import ThreadPoolExecutor
import time

def download_file(url):
    time.sleep(2)  # 模拟IO操作     
    return f"Downloaded {url}"

# 使用线程池
def main_thread():
    urls = ["url1", "url2", "url3", "url4"]
    with ThreadPoolExecutor(max_workers=3) as pool:
        results = pool.map(download_file, urls)
    return list(results)
```
线程操作  （适合CPU密集）

``` python
import asyncio
import aiohttp

async def download_file(url):
    await asyncio.sleep(2)  # 模拟IO操作，适合大量IO操作的时候
    return f"Downloaded {url}"

# 使用协程
async def main_coroutine():
    urls = ["url1", "url2", "url3", "url4"]
    tasks = [download_file(url) for url in urls]
    results = await asyncio.gather(*tasks)
    return results
``` 
%% ENDREGION %%协程操作  （适合IO密集）

## selenium 模块
搜索所有包：在抓包的所有中任意点一个，然后Ctrl+F  是可以搜索的  （这么来找动态加载麻烦）
selenium作用：
	便捷模拟登录，便捷捕获动态响应。
基于浏览器自动化的模块
	让python代表对应的动作和行为，映射到浏览器中 （自动化操作）
安装selenium模块 
	pip install selenium

安装谷歌浏览器驱动：
	参考网站：
		https://blog.csdn.net/NiJiMingCheng/article/details/144231155# （老的）
		https://blog.csdn.net/Z_Lisa/article/details/133307151 （新的）
		主要是和自己的谷歌浏览器版本匹配   chrome://version/ 浏览器中
		https://googlechromelabs.github.io/chrome-for-testing/#stable
	自己放在了：C:\Program Files\Google\Chrome\Application\chromedriver.exe 
		配置了环境路径（其实放同一个文件夹应该也ok）
由于selenium版本更新，测试方法也不一样了，有变化

- 测试代码，包含使用流程

``` python
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
chromedriver_path = r"C:\Program Files\Google\Chrome\Application\chromedriver.exe"
#chromedriver_path = r"./chromedriver.exe" 也行，复制进来
#r保证传递原始的路径，不用\\
# 为了使用 ChromeDriver 正确实例化浏览器
service = Service(executable_path=chromedriver_path)  # 创建 Service 对象
driver = webdriver.Chrome(service=service)  # 使用 Service 对象来实例化浏览器
#新版中推荐使用Service()
# 登录百度
def main():
    driver.get("https://baidu.com/")

if __name__ == '__main__':
    main()
    time.sleep(5)  #延长一下再关闭
```
代码
page_text = driver.page_source 这时候里面就包括了动态加载的数据，不然之前的时候有些需要post，参数，json这样才能返回的，而且很多很杂

- 案例：iframe中读取和移动

	``` python
	import time
	from selenium import webdriver
	from selenium.webdriver.chrome.service import Service
	from selenium.webdriver.common.by import By
	from selenium.webdriver.common.action_chains import ActionChains
	from selenium.webdriver.support.ui import WebDriverWait
	from selenium.webdriver.support import expected_conditions as EC
	
	if __name__ == '__main__':
	    chromedriver_path = r"./chromedriver.exe"
	    service = Service(executable_path=chromedriver_path)
	    driver = webdriver.Chrome(service=service)
	    driver.get("https://www.runoob.com/try/try.php?filename=jqueryui-api-droppable")
	
	    # 显式等待 iframe 加载并切换到它,
	    # WebDriverWait(driver, 2).until(EC.frame_to_be_available_and_switch_to_it((By.ID, "iframeResult")))
	    try:
	        driver.switch_to.frame("iframeResult")
	    except Exception as e:
	        print("无法切换到 iframe:", e)
	
	    target = driver.find_element(By.XPATH, '//*[@id="draggable"]')
	    action = ActionChains(driver)
	    action.click_and_hold(target).perform()
	    # 按住鼠标不一定.perform()不过规范
	
	    for i in range(10):
	        action.move_by_offset(20, 0).perform()  # 移动鼠标
	        time.sleep(0.2)  # 等待
	
	    action.release().perform()  # 释放鼠标
	    driver.quit()  # 关闭浏览器
	
	```
代码
	创建动作链，对鼠标进行操作
	 iframe中的标签需要switch_to.frame进行操作


- 案例：模拟登录qq
	%% REGION %% 
	``` python
	from time import sleep
	from selenium import webdriver
	from selenium.webdriver.chrome.service import Service
	from selenium.webdriver.common.by import By
	from selenium.webdriver.common.action_chains import ActionChains
	
	if __name__ == '__main__':
	    chromedriver_path = r"./chromedriver.exe"
	    service = Service(executable_path=chromedriver_path)
	    driver = webdriver.Chrome(service=service)
	    driver.get("https://qzone.qq.com/")
	    try:
	        driver.switch_to.frame("login_frame")
	    except Exception as e:
	        print("无法切换到 iframe:", e)
	    #在frame中开始寻找
	    target = driver.find_element(By.XPATH, '//*[@id="switcher_plogin"]')
	    target.click()
	
	    username_tag = driver.find_element(By.ID, "u")
	    password_tag = driver.find_element(By.ID,"p")
	    sleep(1)
	    username_tag.send_keys('2311714905')
	    sleep(1)
	    password_tag.send_keys('zhy20040518')
	    sleep(1)
	    btn = driver.find_element(By.ID,'login_button')
	    sleep(1)
	    btn.click()
	    sleep(4)
	    #接着会遇到安全验证，后面解决
	    driver.quit()  # 关闭浏览器
	
	```


- 示例：防止selenium被检测出来
	
	``` python
	from time import sleep
	from selenium import webdriver
	from selenium.webdriver.chrome.service import Service
	from selenium.webdriver.chrome.options import Options
	
	if __name__ == '__main__':
	    #无头访问（无可视化界面）
	    chrome_options = Options()
	    chrome_options.add_argument('--headless')  # 启用无头模式
	    chrome_options.add_argument('--no-sandbox')  # 不使用沙盒
	    chrome_options.add_argument('--disable-dev-shm-usage')  # 解决某些共享内存问题
	    chrome_options.add_argument('--disable-gpu')  # 禁用 GPU 硬件加速
	    chrome_options.add_argument('--window-size=1920,1080')  # 设置窗口大小
	    chrome_options.add_experimental_option('excludeSwitches', ['enable-automation'])  # 禁用自动化，防止被检测出来是selenium
	
	    chromedriver_path = r"./chromedriver.exe"
	    service = Service(executable_path=chromedriver_path)
	    driver = webdriver.Chrome(service=service, options=chrome_options)
	    driver.get("https://www.baidu.com")
	    print(driver.page_source)
	    sleep(5)
	    driver.quit()
	
	``` 
	
复制对应部分就行



- 案例：模拟登录12306
	- 超级鹰实现图片识别等 https://www.chaojiying.com/api-14.html
	   账号hacktoh  密码megumimegumi应该 绑定谷歌邮箱
	   文件夹，python使用文档 
	   
	使用selenium请求页面后应该截图然后截取验证码部分识别，因为对验证码的src发起请求获取图片得到的会是另外的验证码！！（和登陆页面对应的不一样了，内部应该有tag之类的保证对应）
	- 需要使用from PIL import Image，进行屏幕截图操作，选择特定的位置截图，再进一步改变x,y位置坐标 
	- 通过传递给超级鹰的内容可以返回需要的数值进行修改  （很多不同类型的）
	- 直接在F12中可以选中对应元素可以复制xpath！！！！！！！！！！
	- 对于打开的时候刷新

## scrapy框架
框架：继承了很多功能并且有很强通用性的项目模板
功能；： 高性能持久化存储，异步数据下载（twisted），高性能数据解析，分布式
下载：[[../../../资源管理和软件使用/记录工具和用法/IT工具使用#anaconda]]

- 使用：pycharm中右键文件夹，open in termina
	scrapy startproject scrapy_learning     创建文件夹（工程文件夹！！）    
	- 之中的spiders 爬虫文件夹中有很多文件，之中创建一个爬虫文件：
		- cd  scrapy_learning 
		- scrapy genspider  （spidername） www.xxx.com  
	- 执行工程！（不是pycharm运行）：scrapy crawl spidername （--nolog） 这样没有输出日志
	- 需要将settings.py中的robot 改为 false！！！！！！！！！！！！！！！！
	- 同时修改setting中的USER_AGENT方便UA伪装
	- settings.py中添加LOG_LEVEL = ‘ERROR’ 这样就只输出错误的log也行！！
（clear 清理下控制面板指令）

- 数据解析
	- 直接使用xx = response.xpath('')就能获取相关内容了，然后就是前面的xpath方法
		- （复习//text()取之中全部  /text()就取一个  ， xpath返回 都是列表记得\[0]等）
		- xpath返回的是的selector对象的列表，包括xpath和data
			- 所以需要用.extract()    
			- content = ''.join(content)  就能得到都是字符串类型的内容

- 持久化存储
	- 基于终端指令
		- 将parse方法的返回值（需要传返回值！！）存储到本地文件中
			- 比如author和content，通过字典存放然后append到列表中，返回列表
		- 指令：scrapy crawl spidername -o filepath （xx.csv）  只能存储特定类型
	- 基于管道（在items.py和pipelines.py中）
		- 在item类中定义相关的属性
		- 解析数据会封装到item类型的对象中
		- 然后传递给管道（也对应一个函数）做持久化处理 对应函数：process_item中进行编辑
			- （对于管道中的class，暂时理解成按照顺序执行吧）
		- 配置文件中开启管道 （有个ITEM_PIPELINES 对应数字是优先级，小的高）

- 存储案例：一份存储到本地，一份存储到数据库
	- 利用配置文件管道设置中的优先级，同时在管道中新定义一个class 管道 sqlpipeline（object）
		- 需要import pymysql用于数据库连接
		- cursor是用于连接后方便和数据库进行互动操作的，比如fetchall，或者execute
		- commit() 用于保存对数据库的所有更改，使其永久化。
		  rollback() 则用于撤销所有未提交的更改，用于维护数据一致性和完整性。
		- 编码问题，在连接的时候补充charset = utf-8
		- 在优先级高的管道中的process_item一定记得写return item 这样就会自动传递给下一级的继续处理（养成习惯都写上）


- 五大核心模块包括：引擎（scrapy），调度器（scheduler），下载器，爬虫（spiders），管道管理

- 请求传参：
	- 场景：请求解析的数据不在同一页面中（深度爬取）
	- 需求：爬取boss直聘岗位名称和详情页中的描述
	- 类中别的方法需要yield取调用一下，所以有detail_parse  (parse 方法是一个特殊的方法，Scrapy 框架会在收到 start_urls 的响应时自动调用这个方法！！！)
	- 管道，item，分页操作等

- 图片爬取之ImagesPipeline
	-  字符串：基于xpath解析并给管道持久化就行   但是图片：（不适合item）所以使用别的方法，xpath解析处图片的src属性值，单独对图片地址发起请求获取图片二进制类型数据就行
	- 滑动的时候发现会从src2动态变成src，和显示区域有关  （这种确实就是自己探索！！） 所以记得@伪属性  （主代码部分）
		- yield 在这个管道中的作用是产生请求并将其发送到 Scrapy 的调度器，调度器会管理请求的发送和响应的接收。
		- file_path 和 item_completed 方法是系统定义的，并具有特定的功能和用法！（截取url看出）
	- setting中设置IMAGES_STORE = './imgs' 表示图片存储的目录！！

- 中间件初始
	引擎下中间件可以拦截引擎给下载器的请求和下载器的响应
	- 拦截请求
		- UA伪装  (和一开始写不一样，是中间修改)
		- 代理IP更换
	- 拦截响应
		- 篡改响应数据和响应对象 

- 请求处理：
	-    middleware.py对应中间件
		- spider爬虫部分以及downloader下载部分
		- downloader中的
			- process_request     UA伪装
			- process_response
			- process_exception     代理IP 并且记得return request
	- 需要在setting中设置DOWNLOADER_MIDDLEWARES 解封
	- 和selenium类似，获取page的时是.text / .body(二进制) 不是text() (xpaht中)
		- user-agent也是一个列表，使用random随机选取！

- 拦截响应
	- 需求：爬取网易新闻中的新闻数据   可以查看对应的源代码
	- 观察：发现网页存在“正在加载”，说明有动态加载的过程
		-  直接抓包，网易新闻首页解析出五大板块对应的url是可以的，说明没有动态加载！！
		-  进一步分析新闻的详情页的源码，可以得到url，可以检查不是动态加载的
			- 虽然是动态加载，但是实际的F12是可以看到对应的html的，就是爬虫代码获取html数据的时候自己需要在代码中调整方法！！！
	- middle部分对于response修改： （结合五大板块的工作过程理解）
		- from scrapy.http import HtmlResponse   from time import sleep
		- 如果请求的 URL 在 models_urls 中，代码会创建一个新的 HtmlResponse 对象，并将其返回。这表明该响应可能需要经过特定的处理，如解析动态内容或其他逻辑
			- 感觉像是同时读取了request和response的数值进行判断，这样就不会修改request中别的请求了，应该也算一种规范  修改特定请求的特定响应数据
		- 如何获取动态加载的数据？->基于selenium进行操作]]
	- 主体部分
		- callback等等一些调用列表/函数都要记得 self.xxx因为是在class中
		- from wangyipro import  wangyiproItem
		- item部分补充： 记得第二部分title然后传给下去的yield需要meta = {'item:item}传递
		- close应该也是scrapy的lass中自带专有的方法，没有调用就结束
	- pipeline部分为了方便显示就直接print看结果可以

- CrawSpider 类  ： Spider 的一个子类   对应源代码sunpro文件夹
	- 全站数据爬取->所有页码的所有爬取
		- 基于Spider进行手动请求
		- 使用CrawSpider进行全站爬取
	-  scrapy genspider  -t crawl（spidername） www.xxx.com 
	- rule ->规则解析器
	- link ->链接提取器  （通过正则（allow = “内容‘）取筛start）

后面就没有学完力，学完还是要多用项目练手，不然忘的真的快
鉴于通过 Heap 看出 GO 还是有太多需要自己手动实现的了，所以 python 自己也需要搞一些
主要是简化复杂操作，关键知道对应的思维
以后主要还是利用 python 进行面试等，经典的题目也使用 GO

- 最后冲刺企业的时候看下Leetcode 会员的最近算法库

## 常见需要导入的库
自己本地测试等
from typing import List

## 本地测试书写注意
r`es = Solution().longestConsecutive(nums)` # 注意加括号创建实例 对于类的方法


## 语法规则

#### if
- if c == '\['

#### 除
python 中 ：
- / 得到浮点数
- // 得到向下取整
- % 得到余数

#### Bool
True  False  都需要大写

#### 字典加循环初始化
last = {c:i for i,c in enumerate(s)} # 遍历所以会覆盖,标志最后出现的下标

#### 构建列表list（）
`list()` 的构造函数​**​只接受一个可迭代对象​**​
比如 `list(dict.values())`  注意加 s！！！  
或者 list(\[1,2,3]) 其实就是 a = \[1,2,3]

#### 遍历
主要利用 for 进行遍历   如果注重条件可以用 while
######  for i,v in enumerate(a) 才能同时有键和数值
``` python
fruits = ["apple", "banana", "cherry"]
for index, fruit in enumerate(fruits):
    print(f"Index {index}: {fruit}")
```
######  for 变量 in 可迭代对象:
``` python
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(fruit
    
text = "Hello" 
for char in text:
	print(char)
```
######  range 遍历
``` python
# 遍历 0 到 4
for i in range(5):
    print(i)

# 遍历 2 到 5（不包含 6）
for i in range(2, 6):
    print(i)

# 遍历 1 到 9（同样也不包括 9 的），步长为 2   
for i in range(1, 9, 2):  
    print(i)  # 1 3 5 7

```
######  遍历字典
``` python
# 遍历键
person = {"name": "Alice", "age": 25}
for key in person:
    print(key)
# 遍历数值
for value in person.values():
    print(value)
# 遍历键值对
for key, value in person.items():
    print(f"{key}: {value}")
```
###### 嵌套循环 
``` python
matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
for row in matrix:
    for num in row:
        print(num, end=" ")
    print()  # 换行

```

######  zip 同时遍历 两个
``` python
names = ["Alice", "Bob", "Charlie"]
ages = [25, 30, 35]
for name, age in zip(names, ages):
    print(f"{name} is {age} years old")
```

######  列表推导式

``` python
numbers = [1, 2, 3]
squared = [x**2 for x in numbers]
print(squared)  # 输出 [1, 4, 9]
```

## 常用语法糖

#### sort
- sort会在原来的基础上直接排序
	- nums.sort()
- sorted 会返回一个新排序的地址   
	- sorted(nums, reverse=True)  降序排序
- sort（reverse=True） 调整方向

######  sort 列表转换为字符串
`sorted(s)` 返回的是一个​**​列表​**​（`list`），而 `"".join(sorted(s))` 返回的是一个​**​字符串​**​（`str`） 相当于是进行了一次拼接的过程

#### 最大最小值基准
math.inf  代表正无穷大
可以 - math.ind 代表负无穷大

#### \[]相关函数
stack = \[]
stack.append()     stack.append(\[multi,res])  可以添加“组”
stack.pop()


#### collections 包
- 使用`defaultdict`简化统计类题目（如字符频率统计）。
- 使用`deque`处理BFS或滑动窗口问题（如二叉树层序遍历

######  defaultdict(lsit)
defaultdict(list) 会在访问不存在的键时​​自动调用 list()​​ 初始化一个空列表，因此可以直接 `d[sorted_s].append(s)`
而如果是正常创建的字典就需要： `hashmap[sorted_s] = []`

###### Counter 
用于统计频次
- count = Counter(nums)  # 复杂度是 O(n)
	- 接着可以 for num, freq in count.items():
- collection.Counter.most_common(k) 直接统计频次最多的
	- \[item[0] for item in count.most_common(k)]  直接将对应字母取出来[[
]]

#### heapq 堆的库
######  常用函数
- heapq.heappush(heap,(freq,num)) 推入东西    
- heapq.heappop(heap)   推出东西   加符号本质是转换为负的方式存储进去，-heappop 本质还是小顶堆，是不是小顶堆弹出后转换为负数！  负负得正
	- heappop 和 heappush 复杂度都是 O（logn ）   -> 二叉树的上浮和下沉  
heap 本身一般是\[] 就是

###### 统计最大频次
``` python
def topKFrequent(self, nums: List[int], k: int) -> List[int]:
	count = collections.Counter(nums)
	heap = [(val, key) for key, val in count.items()]
	return [item[1] for item in heapq.nlargest(k, heap)]  # 统计频次最大的k个
```



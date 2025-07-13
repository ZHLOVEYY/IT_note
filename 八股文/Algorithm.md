
# 代码随想录
跟着carl学算法   https://www.programmercarl.com/
hello 算法： https://www.hello-algo.com/
算法就是解决特定方法的问题，知道就 ok


#### 二分查找
There is no fixed pattern, left closed right closed (left closed right open is artificial plus one minus one, when indexing) 
- Consider the actual situation
- \[1,1] and \[1,1) to think of it 

###### [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)
一个简单的二分是做不了的，关键是找出上下界   关键注意 5,6,8 而目标是 7 这样的情况，要判断是不是等于
一个找下界的函数就足够，对于上界传入 target+1就可以

[162. 寻找峰值](https://leetcode.cn/problems/find-peak-element/)
为什么要规定 nums[−1]=nums[n]=−∞  -> 说明一定有峰值！

#### 链表
###### [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)
双指针，通过 temp 存储 cur.Next
翻转是基于双指针的方法的

#### 删除链表类
如果需要删除头节点就需要 dummyhead 的～


#### 树
###### [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
从此题可以看出有自顶向下和自底向上两种遍历方法



#### 回溯 
子集型

#### 动态规划

###### [494. 目标和](https://leetcode.cn/problems/target-sum/)
0,1 背包问题
核心思路在于从一些中“选出”，以满足某个条件
改成递归： 记忆化搜索可以改成递归，dfs 改为数组，递归改为循环



###### [322. 零钱兑换](https://leetcode.cn/problems/coin-change/)
完全背包问题 -> 可以选多次

###### 从前往后还是从后往前遍历问题：
关键在于看**每个物品能不能多次用**。如果是不能多次，就需要依赖于“之前一次的状态”，从前往后会更新状态掉状态，所以不行。 



#### 移除元素
- the key point is that you can't just remove the target because the storage space is continuous 
- if use two loop ,one to examine the value and another to move other value forword
- use double pointer
	- fast pointer to check whether the value is the taeget
	- slow pointer to put the non-tatget value into the same array
	- return the slow pointer  (last ++ remember!  then the length of it is the length of the array)

#### 有序数组的平方
- also use two pointer from two side 
	- remember that the array is sorted first

#### 长度最小数组
- use the method of sliding Window
- use another for loop inside to make sure the move of the left pointer
	- l and sum can be a reference

#### 螺旋矩阵
- use the same rule to possess each line


## 链表
#### basic knowledge
- linked list ‘s storage is not continuous , it's random on the memory
- [[Pasted image 20250304103109.png|comparison with array]]

inked list basic：
``` Go
%% REGION %% package main

import "fmt"

type ListNode struct {
	Val int
	//mind the type!
	Next *ListNode
}

// init the first
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// Insert node P after node n0 in the linked list
func insertNode(n0 *ListNode, P *ListNode) {
	n1 := n0.Next
	P.Next = n1
	n0.Next = P
}

// Delete the first node after node n0 in the linked list
func deleteNode(n0 *ListNode) {
	if n0.Next == nil {
		return
	}
	P := n0.Next
	n1 := P.Next
	n0.Next = n1
	//n0.Next = n0.Next.Next
}

// Access the node at index in the linked list
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

// Find the first node with the value target in the linked list.
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		index += 1
		head = head.Next
	}
	return -1
}

func main() {
	//create nodes
	n0 := NewListNode(1)
	n1 := NewListNode(3)
	n2 := NewListNode(2)
	n3 := NewListNode(5)
	n4 := NewListNode(4)
	//create next for nodes
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	num := findNode(n0, 5)
	fmt.Println(num)
} %% ENDREGION %% 
```

list (not linked list) in GO / there is source code for list：
``` GO
%% REGION %% package main

import (
	"fmt"
	"sort"
)

func main() {
	//initial wtth none
	//nums1 := []int{}
	nums := []int{1, 2, 3, 4, 5}
	num := nums[2]
	fmt.Println(num)
	//insert and delete
	nums = nil
	nums = append(nums, 1)
	fmt.Println(nums)
	nums = append(nums, 2)
	nums = append(nums, 4)
	//will not contain index 1
	nums = append(nums[:1], append([]int{6}, nums[1:]...)...)
	fmt.Println(nums)
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)
	//Iterate through the list by index.
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	fmt.Println(count)
	count = 0
	//not in range!
	for _, num := range nums {
		count += num
	}
	fmt.Println(count)
	//Concatenate lists.
	nums1 := []int{2, 3, 9}
	//use...
	nums = append(nums, nums1...)
	//sort  remember typr
	fmt.Println("before sorting:",nums)
	sort.Ints(nums)
	fmt.Println("after sorting",nums)
} %% ENDREGION %%
```

#### problems

###### leetcode 203
https://leetcode.cn/problems/remove-linked-list-elements/description/
 remove head: head = head.next  / use dummy head
- use current.next and current.next.next
	- current = current.next to iterate through
###### leetcode 707: list design
https://leetcode.cn/problems/design-linked-list/submissions/606242491/
 - insert at end
	- condition: cur.next ==  ==nil
- the head (and dummyhead )is important !   mind the index 


###### leetcode 19
dummyNode point to the head!!! solve the problem of deleting the head node

###### leetcode 142
if fast/fast.Next is nil, you can't access to it ！！  so fast == slow to define the location is ok

使用哈希表的方法，复杂度：n，n 

使用快慢指针，这时候slow和fast就不像环形链表1一样初始化（head，head.Next）以及相等作为判断条件 而是具体要求举例，所以可以以fast和fast.Next不是nil为条件循环
（如果是n个（b+c）也是一样还是在进入点相遇的，因为会在里面绕很多圈，等式仍然成立）

## stack
based on array class:  (when write programming problem, use it)
``` GO
%% REGION %% package main

import "fmt"

func main() {
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 4)
	stack = append(stack, 5)
	stack = append(stack, 6)
	//the top of the stack
	peek := stack[len(stack)-1]
	fmt.Println(peek)
	//Ele out
	pop := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(pop)
	//size
	size := len(stack)
	fmt.Println(size)

} %% ENDREGION %%
```
based on linked list: (use \*list.List in go)
``` GO
%% REGION %% package main

import "container/list"

type linkedlistStack struct {
	data *list.List
}

// init
func newLinkedListStack() *linkedlistStack {
	return &linkedlistStack{
		data: list.New(),
	}
}

// Enstack
func (s *linkedlistStack) push(value int) {
	//push back to push in
	s.data.PushBack(value)
}

//check Empty
func (s *linkedlistStack) isEmpty () bool{
	return s.data.Len() ==0
}

//get length
func (s *linkedlistStack) size() int {
	return s.data.Len()
}

// out stack
func (s *linkedlistStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	//compare to queue，just Back/Front！
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

//get top
func (s *linkedlistStack) peek() any {
	if s.isEmpty(){
		return nil
	}
	e := s.data.Back()
	return e.Value
}

//get list
func (s *linkedListStack) toList() *list.List {
    return s.data
}

func main() {

} %% ENDREGION %%
```
can also based on array
## Quene
use the Queue class.
``` GO
%% REGION %% package main

import (
	"container/list"
	"fmt"
)

func main() {
	//make use the "list"pkg in GO 
	queue := list.New()

	queue.PushBack(1)
	queue.PushBack(3)
	queue.PushBack(6)
	queue.PushBack(4)
	queue.PushBack(5)
	//Access the front element of the queue
	peek := queue.Front()
	fmt.Println((peek))
	//Dequeue an element.
	pop := queue.Front()
	queue.Remove(pop)
	fmt.Println(pop)
	//get length
	size := queue.Len()
	fmt.Println(size)
	//isEmpty
	isEmpty := queue.Len() == 0
	fmt.Println(isEmpty)
	//Iterate through the elements
	for e := queue.Front() ; e!= nil ; e=e.Next(){
		fmt.Println(e.Value)
	}
} %% ENDREGION %%

```
can based on array / list
llinked list:
refer to Queue,just Back() or Front()
array:
- We can use a variable front to point to the index of the front element and maintain a variable size to record the length of the queue. Define rear = front + size, this formula calculates rear to point to the position right after the rear element
	- Enqueue operation: Assign the input element to the position at the rear index and increment size by 1.
	- dequeue oper: front index increase by 1 and size index decrease by 1 (rear unchanged)
``` GO
%% REGION %% package main

import "fmt"

type arrayQueue struct {
	nums []int
	//The front pointer points to the front element.
	front int
	//true size
	queSize int
	//can contain size
	queCapacity int
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

// get length
func (q *arrayQueue) size() int {
	return q.size()
}

// isEmpty
func (q *arrayQueue) isEmpty() bool {
	return q.queSize == 0
}

// enqueue
func (q *arrayQueue) push(num int) {
	if q.queSize == q.queCapacity {
		fmt.Println("full")
		return
	}
	// use %
	rear := (q.front + q.queSize) % q.queCapacity
	q.nums[rear] = num
	q.queSize++
}

// get the front
func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

// deque
func (q *arrayQueue) pop() any {
	num := q.peek()
	if num == nil {
		return nil
	}
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return num
}

// print
func (q *arrayQueue) toSlice() []int {
	rear := (q.front + q.queSize) % q.queCapacity
	//remember ...
	return append(q.nums[q.front:], q.nums[:rear]...)

}

// func (q *arrayQueue) toSlice() []int {
//     rear := (q.front + q.queSize)
//     if rear >= q.queCapacity {
//         rear %= q.queCapacity
//         return append(q.nums[q.front:], q.nums[:rear]...)
//     }
//     return q.nums[q.front:rear]
// }

func main() {

} %% ENDREGION %%
```
[[Pasted image 20250304102304.png| use list.List / \[]int is both ok]]

## Double-ended queue 
https://www.hello-algo.com/chapter_stack_and_queue/deque/


## Hash table
 high efficiency in add/delete/search  （search for an element）
Sacrificing space for time ！ 
when we need to check for existence , use hash, (especially set


basic oper and set up
GO version：
``` GO 
%% REGION %% package main

import (
	"fmt"
)

type pair struct {
	key   int
	value string
}

type arrayHashMap struct {
	//use *
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	return &arrayHashMap{
		buckets: make([]*pair, 100),
	}
}

func (a *arrayHashMap) hashFunc(key int) int {
	index := key % 100
	return index
}

// search
func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return "not found"
	}
	return pair.value
}

// insert
func (a *arrayHashMap) set(key int, value string) {
	pair := &pair{key: key, value: value}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

// delete
func (a *arrayHashMap) delete(key int) {
	index := a.hashFunc(key)
	//set nil to delete the pair
	a.buckets[index] = nil
}

// get all keys/values
func (a *arrayHashMap) getAll() []*pair {
	//use *
	//array is set [5]int like this, slice is []int
	result := make([]*pair, 0)
	for _, pair := range a.buckets {
		if pair != nil {
			result = append(result, pair)
		}
	}
	return result
}

//print the hashmap
func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.value)
		}
	}
}

// original oper
func main() {
	hmap := make(map[int]string)

	hmap[12836] = "小"
	hmap[15976] = "算"
	hmap[16700] = "法"
	hmap[17408] = "科"

	name := hmap[17408]
	fmt.Println(name)

	// update value
	delete(hmap, 16700)

	//iterate over map
	//for key / for _,value
	for key, value := range hmap {
		fmt.Println(key, "->", value)
	}

} %% ENDREGION %%

```

bucket will contain the same results of calculate by % 

hash conflict and solution -> list/different storage
##### hash algorithm design
basic：
``` GO
%% REGION %% package main

import "fmt"

// addHash
func addHash(key string) int {
	var hash int64
	var modulus int64

	modulus = 1000000007
	//int + byte will convert string to ASCII
	for _, b := range []byte(key) {
		hash = (hash + int64(b)) % modulus
	}
	return int(hash)
}

// 乘法哈希
func mulHash(key string) int {
	var hash int64
	var modulus int64

	modulus = 1000000007
	for _, b := range []byte(key) {
		hash = (31*hash + int64(b)) % modulus
	}
	return int(hash)
}

func xorHash(key string) int {
	hash := 0
	modulus := 1000000007
	for _, b := range []byte(key) {
		fmt.Println(int(b))
		hash ^= int(b)
		hash = (31*hash + int(b)) % modulus
	}
	return hash & modulus
}

func main() {
} %% ENDREGION %%

```
（hash salt）

#### Examples：
usually in problems hash comeup with： array / set / map  to solve
- small amount - >array  , large amount ->set , key to value -> map
###### leetcode 242
can use the array created by a, then use -- to test b (no need to create to array)

###### leetcode 202 
when we need to check for existence , use set!!



###### leetcode 383
for each word in a string, use for to iterate, it's type "rune"

## double pointer 双指针

#### Examples：
###### leetcode 15 三数之和

错误代码：for \_,v := range res     如果一开始res就是空，就会直接返回了。（去重复的逻辑中需要注意）
一开始需要进行排序方便去除重复，不然会比较麻烦，三元组去重复
sort.Ints()和sort.Slice(ans,func(i,j int)bool{return nums[i]<nums[j]})  其实差不多的

一开始也需要对i进行去除重复（这个能节约很多）
在向中间夹的过程中的写法需要注意   可能还会有比如  -2 -1 1 2 这样的所以还需要往中间靠拢，注意判断l++的次数，理解对应的位置

######  [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)
- method 1 
	use premax  and sufmax to count the one bucket contatin capacity
	space：O（n） time：O（n）
- method 2
	find the max length as a standard and move the other side -> short one minus the height
	the left and right will eventually meat at the tallest 
	space：O(1)  time: O(n)
	


## Sliding windows 滑动窗口
###### [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)
利用单调性，第一个 for 直接对准数值和右下标，注意利用 for 再移动左下标，而不是一直依靠一个 for
如果用 right 会超过下标（自己的方法）就需要做一个额外判断维持住

###### [LCR 009. 乘积小于 K 的子数组](https://leetcode.cn/problems/ZVAVXX/)
子数组的问题，从左边 l 开始到后边 r，一个个加 l，l 和 l+1，l 和 l+1 和 l+2 这样，同时因为 l
###### [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
利用 hash 表/数组 （GO 的 string 中的都是数字）

## string

#### Examples
###### leetcode 541
difference between string and \[]byte [[Pasted image 20250306152216.png]]
	s string can't use   s\[xx] !! need to convert to byte 

###### kama fill
https://kamacoder.com/problempage.php?pid=1064
[[Pasted image 20250306155918.png]] fill in from behind！

###### leetcode 151
' ' refer to byte space and " " refer to string space
\[variable]byte  is unaccepted! should be a const
use two pointer and the reverse function concept!!!

##### KMP algorithm
concept:
- Prefix：contain firt but not contain last!
- Suffix：contain last but not contain first!   But read them in order! not from back
find the max length of same Prefix and Suffix    core idea is to use it to find the next start place

The prefix list tells you where the pattern string(string you want to search for (not in)) should jump to in the next match

Find the position that does not match, then at this time we want to see what the value of its **previous** character prefix list is



problems:
###### leetcode 28
debug many many times (self write)  remember the border(for rec) and restart place for index(v)! to iterate all!
a   a   b   c   a   a   f（a）
0   1   0   0   1   2   0   2


## queue and stack

problems：

###### leetcode 347
维护xx个元素->大顶堆   类似滑动窗口
运用了heap的思想
快速排序利用GO函数：[[Pasted image 20250310194909.png]]
堆也是利用heap  两个都有less函数（封装的）[[Pasted image 20250310195133.png]]
- less都是大于就是大到小或者大顶堆 小于就是。。。

## tree

#### binary tree
https://www.hello-algo.com/chapter_tree/binary_tree_traversal/#__tabbed_1_1
``` GO
%% REGION %% package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{Val: v, Left: nil, Right: nil}
}

func main() {
	//init
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	//insert p
	p:= NewTreeNode(6)
	n1.Left = p
	p.Left = n2
	//delete p
	n1.Left = n2
} %% ENDREGION %%
```

Pre-order traversal (root->left->right),In-order traversal(left ->root->right)，Post-order traversal(left->right->root)
Pre-order example:
``` GO
%% REGION %% func preOrder(node *TreeNode) {
    if node == nil {
        return
    }
    // 访问优先级：根节点 -> 左子树 -> 右子树
    nums = append(nums, node.Val)
    preOrder(node.Left)
    preOrder(node.Right)
} %% ENDREGION %%
```
###### binary search tree
no same val，left is smaller than right
when insert -> will insert into the lowest part(so cursor will be nil)
when delete , focusing on the degree is 0,1,2 ( 2 should use In-order traversal)

完全二叉树：就是最后一排
满二叉树：节点的子节点都是0/2  （叶子节点可能处于不同的层级）
完美二叉树：完美的满
#### 递归和遍历
非递归方式基本都能用栈来模拟  可以使用迭代的方式实现遍历    （了解先）
中序不太一样      栈的出入顺序需要注意
层序遍历对应广度优先搜索   递归对应深度优先搜索
[递归遍历（前中后序）](https://www.programmercarl.com/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E9%80%92%E5%BD%92%E9%81%8D%E5%8E%86.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC)

###### leetcode 102
层序遍历的关键在于，记录size方便栈弹出。层序遍历的模版是很多题目的模版结构！！
[对应两种方法，代码随想录](https://www.programmercarl.com/0102.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E5%BA%8F%E9%81%8D%E5%8E%86.html#_102-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E5%BA%8F%E9%81%8D%E5%8E%86)

###### leetcode 226 and101
翻转二叉树注意遍历时不能用中序，因为会翻转后处理了左边，结果又处理左边
对称二叉树就是从一个节点出发判断左和右  类似后序遍历 （从信息传递上去的角度理解）

###### leetcode 104 and 111 最大深度和最小深度
最大深度采用后序遍历  [[Pasted image 20250313084507.png]]   后序就是最好的 
高度和深度的区别：高度是到叶子结点的距离，反而根的高度最大
	根结点的高度其实就是二叉树的最大深度

一个节点的高度（深度）需要由叶子节点返回

###### leetcode 429
灵活变通，传递的是数组 
每次用GO重新声明，var，也会清空变量

###### leetcode 222 完全二叉树的节点数
可以使用层序遍历或者递归（后序遍历）就是结合left和right的sum加起来，可以结合前面 104，111的深度的代码理解（那个是取最大值）    其实都是比较模版化的！！
满二叉树思路：
- 如果是满二叉树，向左遍历和向右遍历深度相同 （结合因为是完全二叉树） 接着通过2的深度次方减1计算出对应树的节点 
- 如果中间缺了，就不是完全二叉树了

###### leetcode 257 二叉树所有路径
自己思路： 我需要知道两个子节点都是nil，才能判断确实这个节点是叶子结点，可能后序遍历
思路：
- 使用前序遍历！！，才能使父节点指向子节点 中左右  确实后序写不出来 
- 弹出去，弹出去，再加进来的过程涉及到回溯
- 注意最后叶子节点也要收集进来，所以“中”需要在空的返回之前
- 回溯就是push进去后，在递归回来的时候有pop  （如果是递归，那么针对传入的s处理就可以了似乎）


###### leetcode 404 左叶子之和
自己：层序遍历，判断
思路：后序遍历好判断并传递数值～写出判断为左叶子的条件。 这里就是需要收集累计的数值之和，中就是左右节点收集的左叶子之和

###### leetcode 513
存在回溯的过（如果是传入depth+1其实就是隐藏了回溯的过程）   没有中的处理逻辑  
只用if root.Left== nil && root.Right == nil 会出现空指针问题！（也需要条件～） 因为比如一个节点有右无左，就会遍历到左，就有问题了
一开始max应该是0！不然第一行的读入不了

###### leetcode 112  
理解回溯：因该是传入当层的参数（如513中d+1），而不是某个常数！（此题的sum不能在传递中sum+ndoe.Val） 应该当层先加，后传入sum 不然会造成累加
其实利用-=是更快的

###### leetcode 106 通过中序和后序构建树
通过后序确定根，然后在中序中进行切割分出左，右，然后再回到后序中进行切割出左右
（暂时不知道用处是什么，没刷）

###### leetcode 617 合并二叉树
同步进行遍历的过程

###### leetcode 98 验证二叉搜索树
注意在右边的子树中是“所有”都要大于节点的！不是只判断和上一个节点之间的关系（这样才能搜索）
思路：利用二叉搜索树的特性：中序遍历（左中右）肯定是从小到大的！ 所以可以push到一个数组中
如果不使用数组，利用递归法
如果是左中右中序遍历，一开始会到最左下角，这个思路get
利用pre其实就是利用类似数组中两个指针一个个移动，二叉树中的双指针

###### leetcode 530 
类似学习到的指针，数组的逻辑，但是注意赋值！！！！ 还是化简判断会好些

###### leectcode 536 最近公共祖先
思路：需要从下往上手机信息->回溯以及后序遍历，最后在根节点处理 （如果有p/q就传上去对应节点，不然就是NULL）  每个数数值不同同时必然有p和q

###### leetcode 235 二叉搜索树的最近公共祖先
从上到下遍历，如果遇到夹在两个中间的就是了，因为如果再往左或者往右就会有问题

###### leetcode 701 搜索二叉树的插入
插入在叶子结点就行  （别和大顶堆小顶堆搞混）

###### leetcode 450 搜索二叉树删除节点
讨论情况：没有需要删除的节点，左右空，左空右不空（父节点直接指向下一个就行），左不空右空（同上），左右都不空
左右都不空：如果让右孩子继承位置，需要找稍微比要删除的节点大一点
的，所以就是右边子树的最左下角，然后指向原本左孩子  
return root.left的过程其实就是用上一个节点来接收当前的root.left ，相当于跳过当前节点～

###### leetcode 669 修剪二叉树
自己：类似删除节点，但是其实也有大小的传递关系。比如如果node.Val小于low，那么左边后面的肯定都是小于low
确实是这样，注意遍历逻辑就行
###### leetcode 108 有序数组转换为平衡二叉搜索树
每次都取中间节点，这样能保证平衡。    可以直接利用数组索引进行实现

###### leetcode 538 二叉搜索树转换为累加树
左中右   右中左



## heap
是一种特定的完全二叉树
Min heap: the value of any node is less than that of a subnode ，max heap 相反
常用于实现优先队列
[[Pasted image 20250310185520.png|存储和实现]]

## 回溯算法
回溯的本质是穷举所有可能性   可以配合剪枝提升效率
问题种类：
	组合问题（无顺序）
	切割问题
	子集合问题
	排列问题（有顺序）
	棋盘问题
回溯法解决的问题都可以抽象为树形结构。合的大小就构成了树的宽度，递归的深度就构成了树的深度。
注意理解递归的“深度”的概念  以及宽度的对应概念
模版：[[Pasted image 20250316094249.png]]   理解撤销的过程
需要确定：
	递归函数参数
	确定终止条件
	单层递归逻辑


problems：
###### leetcode 77 组合问题
到叶子节点就是我们希望的结果
剪枝操作在for循环中进行，根据条件进行筛选 （注意树的构建过程，选出xx，剩余xx，清楚逻辑）

###### leetcode 216 组合总和
k控制递归的深度（深度终止条件）！！！  宽度由1-9控制
注意树构建，取出1，剩下\[2-9] 取出2 ，剩下\[3-9] [[Pasted image 20250316110757.png]]
注意：
- path 添加到 res 时，你直接使用了 res = append(res, path)。由于 path 是一个切片，后续对 path 的修改会影响 res 中已存储的结果。需要深拷贝 path！！！ 需要tmp进行copy处理
- 如果用全局变量的sum   局部的targetsum可以传递给下一层而在当前层不变  （很重要的想法）
- 需要是i+1而不是startindex+1，因为i会基于startindex变化！（1，2，4，5） 遍历了124后下一个125

###### leetcode 17 电话号码组合
这题的index及时对应深度（有几个数进来了） 也就是返回条件
如果用形式参数的形式传入 letter + x 这样就不用pop的   [[Pasted image 20250317104958.png|结构图]]

###### leetcode 39 组合之和
自己：没有搜索深度k的限制了，感觉应该先给数组进行排序？
思路：还是需要startindex来进行。其实不用排序，关键还是回溯的逻辑，选中的和剩下能选的，形成树
[[Pasted image 20250317141230.png]]    关键在于理解回溯的逻辑
注意浅拷贝，本质是共享底层数组 （之前17题转换成string后类型就不变了的）

###### leetcode 40 和2
注意i+1而不是startindex+1 (这样会有很多重复的了)
同时需要注意去重复，因为同一个数字不止出现一次，比如 1.。。7.。。1.  这样就会有17和71 （target为8的话）
采用used来标记的话：
used可以工作的原因是第一层回溯的时候其实就都给赋上了flase（第一次遍历的时候是没有阻拦的！！）

###### leetcode 131 分割回文串
注意采用stratindex到i 这样来理解分割！[[Pasted image 20250318082427.png]]
注意copy的时候对应的temp需要初始化对应的长度才能copy   copy的前一个是copy进的对象

###### leetcode 93 复原ip地址
思路：应该是先分组，然后进行加.处理后，放入res，最后再对re
[[Pasted image 20250318085417.png|利用strconv和strings.join转换]]题目中是字符串转字符串更好转
num, err := strconv.Atoi(s) 是转换成数字     Atoi：​A​（ASCII） ​to** ​I​（Integer）Itoa 同理  首大写就行
自己思考过程中也是运用了多次剪枝和判断

###### leetcode 78 子集问题
之前都是在叶子节点进行收获，如果采用之前的循环就会漏掉比如单独的元素\[1]这样
所以就是进入一层递归就放进去（一开始的空也会放进去） 理解之前的startindex其实就是判断是不是到了叶子节点
排序sort.Ints(nums)
i>startindex 也是可以的，因为“那一层中”（比如第一层）i向前进，也会让第一位的1变为false

90 重复子集问题   加上去重复就行

###### leetcode 递增子序列
PS：似乎声明：= 如果当前函数没有用的话也不太行  还是var  然后 = 这样子函数中用了就可[[Pasted image 20250318110528.png|used函数]]
[[Pasted image 20250318112144.png|树层要去重]]  这里就要注意used定义的位置了，每层的:=其实用的used都是不一样的，不能用一开始的一个used去统计所有的，那样比如477就统计不到了！！！！！！！
进一步理解树枝和树层

###### leetcode 46 全排列
不用考虑去重，因为不重复
排列和组合不同比如123和321排列中是不一样的！  不过还是利用树的结构去理解，取出和还可以取的
关键就是在于灵活运用树形结构！！
利用st去判断元素是和否使用，第二次递归的时候会还判断所依132是可以的 （i每次都是0开始）

###### leetcode 47 全排序去重
注意if st\[i] 作为条件判断加入！  以及continue和return 的理解    树层去重


## 动态规划
### 概念
动规基础，背包问题，打家劫舍，股票问题，子序列问题
需要掌握本质步骤：五部曲
dp数组的定义以及下标的含义   dp\[i]\[j] 或者dp\[j]
递归公式
dp数组初始化
遍历顺序（比如先遍历背包还是物品） 前后，几层
打印数组进行检查    看看是哪里的问题

###### leetcode 斐波那契数 509
确定dp数组： dp\[i]是第i个斐波那契数
递归公式：就是公式
初始化： dp数组的0和1 为1
遍历顺序 从前向后
打印dp数组 （用于debug） 

动态规划，递归，普通的sum相加然后交换，方法都要理解
注意 n<2 的坑就行～


###### leetcode 70 爬楼梯
递归理解：爬到第n阶是由n-1和n-2 的组合来的！（都是一步就到了/ 爬2步，如果n-2爬1加1其实相当于n-1的类型里面的，很有意思）
dp\[i]达到i阶有多少种方法   所依就是i = i-1 + i-2i-2zi-2z

自己才发现语法，声明过res，赋值i的时候res\[i]直接=就行


###### leetcode 746 最小花费爬楼梯
自己：应该是计算每一步或者每两步哪个花费更少？但是不如1，2，3，走了1还是要2，3，不如直接2
dp\[i]数组含义：到达i所需要的花费 （不包括当前位置的）
 画出cost和dp数组，就很容易想清楚初始化长度，以及遍历的时候i从多少开始终止条件是多少
 这题对于cost 10，15，20 就知道dp应该是四个格子，然后求第四个格子的数值

###### leetcode 62 不同路径
自己：dp应该是表示？移动i，j格子有多少种方法  ij 对应i-1，j和i，j-1的方法相加
自己做出来了！注意初始化第一行和第一列都是1，而不止01和10
注意数组的实际大小理解好，其实就可以的

###### leetcode 63 不同路径2
用GO没赋值默认是0！   同时注意 设置0的话就不用多余的if判断了
不过这题也算自己a出来的，也考虑到了1到初始行的情况～      

###### leetcode 343 整数拆分
dp\[i]就是拆分i之后得到的最大数  比如10->2和8  又可以拆8的
j* dp\[i-j]遍历后其实就包括了所有的可能性 
注意新增加的j\*(i-j)本身的可能性以及之前的dp\[i]需要遍历求最大,dp\[1]= 0或者1其实似乎不影响，往往都是要取平均，中间的大
###### leetcode 96 不同搜索二叉树
自己：感觉应该也是，从之前的结果递推过来？
思路：注意以j为头节点的时候遍历，有j-1个比j小的，所以在左边，有i-j个比j打的就在右边力
dp\[i]的含义是i个数的可能的二叉搜索树的情况       所以左边乘右边（笛卡尔积）就是所有可能  （只用管个数就行，肯定是一个小到大的排列，所以n个节点也就等于dp\[n]，不用管数值绝对值（比如123和789组成的二叉搜索树种类是一样的））
初始化注意dp\[0]也算是一种

###### 背包问题
面试而言 01背包和完全背包足够
01背包：n种，每种一个      01背包的暴力解法就是取和不取  回溯可以解决，复杂度2的n次方
- dp\[i]\[j]   的含义是\[0,i]物品，容量为j的背包
	不放物品i就是\[i-1]\[j]
	放物品i: dp\[i-1]\[j-weight\[i]] +value\[i]      如果j-weight小于0就说明放不下
初始化，结合实际含义
对于01背包二维数组的实现先遍历背包或者 物品其实都可以（因为主要是由正上方和左上方数值推导出来）

完全背包问题： n种，每种无限个    其实都是由01背包衍生过来的
对于一维数组而言，需要前面的状态，座椅是++而不是--
遍历顺序：两层for循环可以调换（一位数组） 因为二维来看都是从“前面的状态（左侧）”推导过来的


滚动数组
原理：如果把dp\[i - 1]那一层拷贝到dp\[i]上，表达式完全可以是：dp\[i]\[j] = max(dp\[i]\[j], dp\[i]\[j - weight\[i]] + value\[i]);
初始化为0就行，后面会覆盖
遍历顺序：遍历物品然后倒序遍历背包  （保证每个物品只被添加一次）  倒序和先遍历物品比较要求严格
j从最大重量遍历到等于当前循环物品的重量（--） 避免超出下标
倒序是因为会如果正序会取之前物品的状态（重复取同一个物品的状态！之前二维数组不会，是从上一个物品的状态衍生过来的）
一定是先遍历物品再遍历背包，因为如果后遍历物品，这样存放的就只是一个物品数值了（因为覆盖，会在一个位置先把所有物品过一次，那么就是留下最后一个物品的）
###### leetcode  416 分割等和子串
我们需要找到一些元素，使得它们的和等于 target。这与 0-1 背包问题中选择物品使得总重量等于某个值非常相似
理解成物品的重量和价值是一样的，就是数字
问“能不能装满”

###### leetcode 1049 最后一块石头
关键在于想到将石头分几乎等质量的两堆       价值就是重量
假设其中一组的重量为 dp\[target]，那么另一组的重量就是 sum - dp\[target] 所以就get了，后者肯定大
问“怎么装最满  ”


###### leetcode 494 目标和
感觉都是分为两部分！分为加法的集合和减法的集合
利用数学公式推导  left 和right  相加为sum，相减为target  得到关于left的要求表达式。
这题的本质是问有多少种方法能装满容量为left的01背包 （确实和价值的理解就不可以一样了，需要精确等于才能算是“方法”
此时需要注意dp数组含义，装满j有j种方法    （是由前面的加和而来，对应不同的nums\[i]）


###### leetcode 474 一和零
装满这个容器（维度是m个0和n个1，两个维度，不只是“重量”）需要几个物品（子集合数量） 这类问题
	决定求的是需要几个物品就决定了递推公式的逻辑
	不是多重背包
物品的重量：x个0和y个1  （m和n是最大要求


完全背包问题
###### leetcode 零钱兑换
注意还是选择“多少种方法”，初始化为1     
dp数组含义是组成金额j有多少种方法！
- 先遍历物品再遍历背包是组合数（一个金额先用一次），先遍历背包，再遍历物品就会有排列数了
举例get就行
``` GO
//对于amount为3 【1，2】数值  先遍历物品
func change(amount int, coins []int) int {
    res := make([]int, amount + 1)
    res[0] = 1
    for _, coin := range coins {
        for j := coin; j <= amount; j++ {
            res[j] += res[j - coin]
        }
    }
    //[1,1,1,1] [1,1,2,2]  第一个1是dp[0]的1  两次循环对应两个coin
    return res[amount]
}
//先遍历背包
func change(amount int, coins []int) int {
    res := make([]int, amount + 1)
    res[0] = 1
    for j := 0; j <= amount; j++ {
        for _, coin := range coins {
            if j >= coin {
                res[j] += res[j - coin]
            }
        }
    }
    //[1100] [1120]  [1123] 三次输出对应三个位置
    return res[amount]
}
```


###### leetcode 377 组合总和
其实就是上一题，排序！ 
拓展：如果爬楼楼梯，一次可以爬m个台阶，本质上和这题是一样的（排列～）

###### leetcode 279 完全平方数
自己：需要返回的是“元素个数”而不是种类
取min？  ->需要“悲观初始化” 取都是1或者maxint都可以的


###### leetcode 139 单词拆分
自己：可以重复使用，不是说选最多或者最少，只要ok就可以返回true
也可以用回溯的方法
递推公式：如果dpj到i出现在字符串中，同时dpj是ture，那么dpi也是true
初始化：dp0应该是true
对于顺序有要求！！112和121不一样，所以应该先遍历背包再遍历物品（想一下也知道，先遍历物品会导致错过物品）


# 刷题记录
速刷吧，少年，思考5min不行就看题解理解，然后关键是自己写，最后再花时间总结。后面看自己写的总结写题复习
（初期方法）
#### hot100
###### 128 最长连续序列
利用hashset去重，然后判断条件是num-1是不是false，因为如果num-1是true其实从num-1开始遍历长度更好
- for num := range  nums  num是下标！   这样都没有用到map没有去重复会超时的
- for num := range map 取的是key

###### 11 最多水容器
一左一右压缩，往中间挤    谁小就动谁，记得底部长度的设置和变化

###### 3 无重复字符的最长子串
利用滑动窗口的思路： 第k个元素如果最长到ak，那么第k+1个就去掉k，然后肯定也会到ak，接着向下
注意一般都是利用map来去除重复！！！（刚写完18，如果用自己的方法也是）

###### 438 子串中的所有异位词
思路：由于长度是固定的，所以可以设置固定长度的滑动窗口，比较之中的字母是否在p中
字符串的基本都是利用\[26]int 这样存储比较   数组可以直接sCount == pCount ，但是切片就需要equal(charCount, windowCount) 这样进行比较的！
map就会有问题的，因为便利顺序不定（这里只能依靠对应的slice一样！）


###### 560 和为k的子数组
一开始自己采用的是双指针的思路，双指针依赖于单调性，这么处理是不对的！
这题基本方法可以暴力遍历： 0-i中间有个j  然后j到i遍历求和，这样就需要n3复杂度了，j-i其实可以从i返回去推，这样j-1到i就只用在j-i基础上加上j-1就可以了
这题利用前缀和的哈希表优化：（方法一的瓶颈在于对每个 i，我们需要枚举所有的 j 来判断是否符合条件）
pre\[i]−pre\[j−1]\==k   pre是对应的前缀和   移动后就是pre\[j−1]\==pre\[i]−k   （j-1就是-1到i-1）
 所以我们考虑以 i 结尾的和为 k 的连续子数组个数时只要统计有多少个前缀和为 pre\[i]−k 的 pre\[j] 即可 (j是从0到i的，所以可以到i-1～)
 注意初始化需要m\[0]= 1 代表有前缀和为0的一个
 关键就是记录：0-i之中满足前缀和和pre-k的个数，

增加m打印：fmt.Printf("  更新后 m=%v\n", m)


###### 76 最小覆盖子串
哈希，滑动窗口
tip：string遍历是rune类型    ，可以有map contain的比较方法，虽然遍历顺序不一样，但是可以取出key然后比较的
``` GO
func isContain(hashs, hasht map[rune]int) bool {
    for char, count := range hasht {  //只用取出hasht中的
        if hashs[char] < count {
            return false
        }
    }
    return true
}
```
和之前不一样的是这个是子串，不能数组\==来判断了
- hasht := make(map[rune]int)  需要初始化的！不能var只声明

思路：设置左位置用于获取切片长度的！
优化思路：针对每一个字母，判断是否需要再进行count计算
注意，增加和减少都需要小心判断是不是在hasht之中
补充考虑没有满足的情况，minlen会比len长
``` Go
if hashs[rune(s[left])] < hasht[rune(s[left])]{ //如果没有key，默认返回0！！
                    count --
         }
```

###### 53 最大子数组和
没有单调性，而且不能排序，所以用不了双指针
动态规划
```
f(i) 代表以第 i 个数结尾的「连续子数组的最大和」，那么很显然我们要求的答案就是：
max 0≤i≤n−1 f(i)
​
 {f(i)}

状态转移方程为：dp[i] = max(dp[i - 1] + nums[i], nums[i])
这是因为以 nums[i] 结尾的最大子序和要么是将 nums[i] 加入到以 nums[i - 1] 结尾的最大子序和中（如果 dp[i - 1] 为正，这样能增大和），要么就是 nums[i] 本身（如果 dp[i - 1] 为负，那么加上 dp[i - 1] 会使和变小，不如直接从 nums[i] 开始新的子数组）。
```

暴力解法理解，其实可以参考[[leetcode刷题笔记#560 和为k的子数组]] 类似倒着的加上去的遍历方法

###### 56 合并区间
如果利用双重循环罗列会非常复杂，最后卡在j是比i大的，这样就不能剔除j了
思路一：
左端点进行预排序sort.slice ，然后就可线性扫描合并了，res中放入interval\[0],因为如果重合就更新，不重合，就将interval加入res中，会继续和res的最后一个比较（这样才有重叠的可能）

经典sort.Slice:
``` GO
    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][0] < intervals[j][0]
    })  //进行排序了
```

###### 189 轮转数组
自己暴力比较复杂 k\*n的时间复杂度
方法二：环状替换：
位置 0 开始，最初令 temp=nums\[0]。根据规则，位置 0 的元素会放至 (0+k)modn 的位置,然后用temp接住对应(0+k)modn位置的元素，但是注意会有遗漏的，比如k=2，4长度，这样1，3遍历不到，所以也要移动初始位置
不过这个公约数方法理解有点麻烦，自己先不搞

方法三：reverse来看
取模之后，
尾部的k个会到前面，前面的n-k个会到尾部。所以整体翻转，然后分两部分翻转


######  238 除自身以外数组的乘积
左右数组的方法，存储数字左边和数字右边的乘积数值

###### 41 缺失的第一个正数
用hash表罗列（n，n） 或者遍历数组（n方，1）都不行，基于原有的数组进行修改！
思路：N长的数，结果肯定是1到N，不然如果1到N都有那么就是N+1
因此于遍历到的数 x，如果它在 [1,N] 的范围内，那么就将数组中的第 x−1 个位置（注意：数组下标从 0 开始）打上「标记」  最后返回最小打上标记的+1  。  如果都有标记就是N+1
考虑到有负数需要处理 ->变为N+1  这样就能采用标记方式：将对应x-1位置变为负数（表示下标为x的+1的数字对应正数已经出现过了）
最后就是找最先出现（相当于找最小）在数组中1到N中对应的数字，返回下标（此时下标是最后出现的正数）+1
（理解为x以及出现过了，就是）


###### 73 矩阵置零
利用行列锁定的思维
方法一：新建两个代表row和col切片，然后有0的话就对应位置行列置0    遍历两次实现（第一次标记，第二次看有标记就实现）
方法二：基础上直接利用矩阵的第一行和第一列作为标记（因为要是中间有0肯定也要置0）一开始判断是否要将第一行和第一列最后再初始化


###### 160 相交列表
这类题一般不会有dummyhead 说的头就是链表的第一个点
针对指针，pA，pB ：=headA，headB 是拷贝
&Listnode{}和nil也是不一样的！
var pre \*ListNode = &ListNode{}  这个就是0 而不是nil 了，var pre \*ListNode 就是nil！
双指针思路： a+c 和b+c （如果有共同c）  那么肯定会a+b+c （刚好相遇）再多走c，神奇
如果没有相遇就是都走了m+n （两个数组长度）最后同时会为nil

###### 234 回文列表
自己尝试结合翻转链表，复杂度（n，n）
可以基于原数组，slow和fast找到中点然后翻转 
注意不要写fast.Next.Next这样的，比如\[1] ,\[1,0,1] 就容易越界！

###### 环形链表
不用返回位置就好很多
```
fast==nil||fast.Next ==nil   和 fast.Next ==nil|| fast==nil   后者不对！
前者先检查 fast == nil，如果为 true，直接跳过 fast.Next 的判断（短路求值）。
```
除了快慢指针还有哈希法的方法

循环的条件是for slow！= fast～


######  21 合并两个有序列表
自己第一次用两个指针判断的方法a出来了   （更新时候先更新newhead还是head1都可以）不过一开始需要realhead存一下newhead，因为会
自己的方法其实就是迭代，复杂度n，1（（仅使用了常数级别的额外变量（如 newhead、realhead、head1、head2））

递归方法：视为这个函数每次都处理当前第一个节点
原因：每次递归调用 mergeTwoLists 时，系统会在内存的调用栈中压入一个新的栈帧（stack frame），用于保存当前函数的局部变量（如 l1、l2）、返回地址等信息。递归的深度取决于链表的长度，最坏情况下需要递归 m+n 次（两个链表的所有节点各被遍历一次），因此栈空间的大小为 O(m+n)


###### 94 中序遍历
自己第一次没意识到用指针！需要注意

迭代的方法是用栈来维护，一直寻找

发现var res \[]int{}, var res \[]int 和res = make(\[]int,0) 和res = \[]int{} 之间辩证


###### 随机链表的复制
本题中因为随机指针的存在，当我们拷贝节点时，「当前节点的随机指针指向的节点」可能还没创建。所以如果没有创建就立刻递归进入进行创建
同时为了防止重复拷贝（不能总是创建，可能多同时指向一个）
map\[\*Node\]\*Node 含义是建立原节点到新节点，需要的是存储的Node的数值


###### [148. 排序链表](https://leetcode.cn/problems/sort-list/)
这题主要是注意怎么找重点，利用2/4 个点来假设看看，还是next和nextnext作为条件好
采用递归的方法，归并排序


######  [146. LRU 缓存](https://leetcode.cn/problems/lru-cache/)
快速访问用hash，移动顺序用链表，两个要求
对于put和get操作，明确对应key在和不在时怎么变化
moveToHead 可以由remove和addToHead组成


######  [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/)
括号的是需要有匹配顺序的，如果直接有右括号不行  
\[ ( ] ) 这样有交叉括号也不行

######  [155. 最小栈](https://leetcode.cn/problems/min-stack/)
构建另外一个栈，原来的栈只要是最小的，另外一边就放入对应的最小值。 原来栈有更小的放入后，另一边的栈最小值也更新，后面也都放入这个    （一种“永久标记”）

######  [394. 字符串解码](https://leetcode.cn/problems/decode-string/)
嵌套括号，需要由内到外构建字符-> 栈的特点**先入后出**
数字放 multi    字母放 res 遇到\[  就同时放入栈， 遇到\]就出栈，res = 之前 res+现在 res*\multi

######  [739. 每日温度](https://leetcode.cn/problems/daily-temperatures/)
单调栈类题
单调栈中同时存放 i 以及对应的温度，如果有大的进来就做差然后取出对应的 i（温度）。这样维护
stack 中真实存放的只有 i，需要比较的时候直接利用数组索引去查就可以

######  [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/)
单调栈类题
暴力思路：对于任意一个矩形，向左边和右边分别扩散看能扩散多远  

改进：空间换时间，用单调栈
每一次计算最大宽度的时候，没有去遍历，而是使用了栈里存放的下标信息，以 O(1) 的时间复杂度计算最大宽度。
从左向右遍历，有严格小的时候就说明之前高的那个是固定了，弹出（符合后进先出的特点）。然后最右边增加一个 0 的哨兵，再从右到左出栈
注意正推和回推时候的公式 需要先 pop 再作差代表宽度

######  [数组中的第K个最大元素](https://leetcode.cn/problems/kth-largest-element-in-an-array/)
直接 sorted
或者快速排序

######  [347. 前 K 个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/)
方法很多，利用 counter 或者 heap 都可以
也可以快速排序或桶排序

[295. 数据流的中位数](https://leetcode.cn/problems/find-median-from-data-stream/)
使用两个堆实现是最好的
如果使用单一的数组：每次调用`findMedian()`时都对整个数组进行排序，导致时间复杂度达到O(nlogn) （排序所需要的算法）  虽然说这样插入上是 O（1）
而最优解应该实现O(logn)的插入和O(1)的查询
- 同时nums = sorted(self.nums) # 仅排序一次 可以后面统一用 nums 而不是总是 self 

######  [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)
枚举的话，第一天后面 n-1种可能，第二天后面 n-2种 ，加起来是 O（n 方） 不好搞
贪心思想：若在前 i 天选择买入，若想达到最高利润，**则一定选择价格最低的交易日买入**
分析：
- 前 i 天的最小 cost 我们需要获取  （这个遍历一次就行）
- 第 i 天的最大利润 = max（i 天之前的最大利润，第 i 天的钱-前 i 天的最小 cost）
这样其实只有遍历一次价格表就可以了

######  [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)
掌握贪心条件
设置一个“最远可到达距离” 进行维护，同时记得一开始 daoi 就需要判断一下

######  [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/)
不同之处在于返回最小跳跃次数
自己想法：应该是维护一个 minjump，应该要用到动态规划
可以建桥的方法->贪心
或者动规划   或者动规加贪心（每次从最能跳的地方开始追随到 i）

###### [763. 划分字母区间](https://leetcode.cn/problems/partition-labels/)
列出每个字母的所在区间（一共 26 个字母最多，不多）
然后就是合并区间->几个区间的长度就是答案，最长就是整个算一个


######  [1. 两数之和](https://leetcode.cn/problems/two-sum/)
Map {key:value}使用数组中元素的值作为键，索引可以作为值存储。
因为一定有答案，所以这样只用遍历一次链表


######   [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/)
排序后(如果可以放一块就是一样的)的字符串当作 key，原字符串组成的列表（即答案）当作 value 


[128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/)
如果采用排序然后算 max 的方法，排序复杂度有 ologn   那么看到要求 O（n），所以说用的就应该是哈希表



# 随意打击

###### [767. 重构字符串](https://leetcode.cn/problems/reorganize-string/)
思路：
每次取出统计中数目最多的两个字母拼在一起，后面再拼别的
GO：
需要用到 heap.Init() 需要自己定义结构体和方法 Less，Pop，Push 等
关键在于如何存入数据以及如何正确取出和返回
Python：
直接用 collection统计这也....太方便了吧，果然是思路导向的


# 排序算法

#### 快速排序
######  215的快速排序 （分治）： 
非常明显，较大分块
``` python
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        def quick_select(nums,k):
            # 随机选取基准
                pivot = random.choice(nums)
                big,equal,small = [],[],[]
                # 将大于，小于，等于的 pivot 元素划分到 big，small，equal 中
                for num in nums:
                    if num > pivot:
                        big.append(num)
                    elif num < pivot:
                        small.append(num)
                    else:
                        equal.append(num)
                if k <= len(big):
                    # 第 k 大元素在 big 中，递归划分
                    return quick_select(big, k)
                if len(nums) - len(small) < k:
                    # 第 k 大元素在 small 中，递归划分
                      return quick_select(small, k + len(small)-len(nums)) # 画图理解可以
                return pivot
        return quick_select(nums,k)
```






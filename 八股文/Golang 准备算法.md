记录 GO 刷算法时易忘记的语法
![[../attachments/Pasted image 20250605210856.png]] 这几个 goland 中的插件先关闭
## 常用的额外库

#### 排序
sort.Ints(nums) 对数组进行排序
sort.Float64s(a []float64)：对 float64 切片升序排序。
sort.Strings(a []string)：对 string 切片按字典序排序

#### 取绝对值
对于int： 自己定义函数
``` GO
func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```
因为 math.Abs 只支持浮点数的判断

#### 列表添加 nil
``` Go
func main() {
	var ans []*int
	ans = append(ans, nil)
	fmt.Println(ans)
}
```
这样是可以的，会添加 nil 占位，但是如果\[]int 就不可以添加 nil 了
#### 列表初始化
``` GO
var ans []int
ans[0]= 1   //由于没有初始化长度，不可以赋值！,不过去掉这一句可以正常打印空队列
fmt.Println(ans)
```


#### 构建最大值最小值
math.MinInt, math.MaxInt

#### 哈希表返回值
```  GO
hashmap := map[int]int{}
hashmap := map[int]bool{}
if i,ok := hashmap[target-x];ok{   
            return []int{i,j}
        }
```
i 会是数值，而不是索引！！ 如果是 bool 类型的定义就放不进 int 表了
如果希望获取下标可以自己定义函数或者双向映射：
``` GO
// 正向映射：key=下标, value=数值
hashmap := make(map[int]int)
// 反向映射：key=数值, value=下标
reverseMap := make(map[int]int)

// 初始化时同步维护两个map
for idx, num := range nums {
    hashmap[idx] = num
    reverseMap[num] = idx
}

// 使用时直接通过反向map获取下标
if i, ok := reverseMap[target-x]; ok {
    return []int{i, j}
}
```

#### 切片排序

``` Go
import (
    "fmt"
    "sort"
)

func main() {
    s := []int{1, 2, 3, 4, 5}
    sort.Slice(s, func(i, j int) bool {
        return i > j // 通过逆序比较实现反转
    })
    fmt.Println(s) // 输出: [5 4 3 2 1]
}
```

#### 切片翻转
自己写函数
``` Go
func reverseslice(tar []int) { //引用类型，直接放入
    left,right := 0,len(tar)-1
    for left<right{
        tar[left],tar[right] = tar[right],tar[left]
        left ++
        right--
    }
    return
}
```

通用：
``` GO
func reverseSlice[T any](s []T) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i] // 交换首尾元素
    }
}
// 使用示例
s := []string{"a", "b", "c"}
reverseSlice(s) // 输出: ["c", "b", "a"]
```

#### copy 深拷贝
``` GO
temp := make([]int, len(path))  // T为元素类型（如int、string等）
copy(temp, path)              // 复制path的所有元素到temp
ans = append(ans, temp)       // 将独立副本加入结果集
```
或者
``` GO
ans = append(ans, slices.Clone(path))
```
也是可以的

####  map 查找
对于
``` Go
onPath := make([]bool,len(nums))
if _, ok := onPath[j]; !ok 错误地使用了 map 的语法检查 onPath[j]
```
这个不对的，onPath 是切片，直接！onPath 判断就行了
直接遍历：
``` Go
 for j, on := range onPath {
            if !on {
```

#### 遍历的值拷贝问题
``` GO
f := []int{10, 20, 30}
for _, x := range f {
    x = 1
}
fmt.Println(f)  // 输出 [10, 20, 30]，未被修改
```


## ACM模式下

#### 基础输入输出
######  单行数字/字符串
``` GO
var a, b int
fmt.Scan(&a, &b)  // 读取空格或换行分隔的数字
fmt.Scanln(&a, &b) // 读取一行后停止（换行符结束）
```
###### 整行字符串
``` Go
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan() // 读取一行
line := scanner.Text()
```
###### 多组数据循环读取
``` Go
for {
    var a, b int
    _, err := fmt.Scan(&a, &b)
    if err != nil { // 遇到EOF或错误终止
        break
    }
    fmt.Println(a + b)
}
```

#### 字符串和数字转换
``` Go
num, _ := strconv.Atoi("123")  // string → int
str := strconv.Itoa(123)       // int → string
```
###### 字符串分割
``` GO
parts := strings.Split("a b c", " ") // 按空格分割为切片
```


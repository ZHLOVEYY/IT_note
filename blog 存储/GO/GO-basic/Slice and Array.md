切片：定义，切片长度（实际大小len），容量（底层数组大小cap）  三个组成
#### 基础知识
- \[]中有数字就是数组，没有就是切片！ （牢记）
	- \[]int{1,2,3}就是切片
- 通常利用make()进行创建  如果使用{}就是直接书写数组内容
- \[]中不可以是变量比如 k  ，后面这样可以：\[...] int {1,2,3,4}  / \[5] int {1,2,3,4,5}
- make 初始化的时候如果只指定一个参数，那么就是切片长度=容量 这样初始化   注意只要是初始化的，都会用 0 占位！！
- 理解res := \[]int{}   或者  var res \[]int 或者 res := make(\[]int,0)   不同的创建方式，经常还是用 make 更加直观
###### 数组和切片的区别
array：数组   （不过变量命名无所谓）   slice：切片
![[../../../attachments/Pasted image 20250525104949.png]]
基本用的都是切片
#### append函数
###### 通过append看切片和扩容
- 不能直接使用索引来突破切片的长度限制
``` go
	func main() {
	a := make([]int, 5, 10) //长度
	a = append(a, 1, 2, 3,4)
    fmt.Printf("追加后切片长度为：%d，容量为：%d\n", len(a), cap(a))
    fmt.Println(a[0:5])
    fmt.Printf("追加后赋值的切片长度为：%d，容量为：%d\n", len(s1), cap(s1))
    fmt.Printf("追加后原来的切片长度为：%d，容量为：%d\n", len(a), cap(a))
}
//追加后切片长度为：9，容量为：10
//[0 0 0 0 0]
//追加后赋值的切片长度为：12，容量为：20 
//追加后原来的切片长度为：9，容量为：10
```
######  通过 append 合并切片
``` GO
func main() {
    array := []int{1, 2, 3, 4, 5}
    array2 := []int{1, 2, 3, 4, 5}
    res := append(array,array2...) //需要加...
    fmt.Println(res)
    //[1 2 3 4 5 1 2 3 4 5]
}
```
算法中常用：
``` GO
a := int(len(numbers)/2) //取出中间元素的位置
nunbers = append(numbers[:a],numbers[a+1:]...) //去掉这个个数
```


 #### 利用\[:]的方式制造切片
 ###### 截取示范
 - :后面的是不包括在里面的
	``` GO
	func main() {
	    array := []int{1, 2, 3, 4, 5}
	    s1 := array[:3]
	    s2 := array[:]
	    fmt.Println(s1, s2) // [1 2 3] [1 2 3 4 5]
	}
	```

###### 截取的地址分析
- 二者地址上就差8个字节
``` GO
func main() {
    array := []int{1, 2, 3, 4, 5}
    s1 := array[1:]
    fmt.Printf("array 切片地址%p\n",array) //不是 array 这个变量的地址，做好区分
    fmt.Printf("s1 切片地址%p\n",s1)
    fmt.Printf("array 变量地址%p\n",&array) //和上面切片地址是不一样的
}
```

#### := 赋值的影响
赋值的操作本质上就是参数拷贝
利用:=来共享底层数组的切片，修改时会同时影响到
``` Go
func main() {
    array := []int{1, 2, 3, 4, 5}
    s1 := array
    s2 := array[1:3]
    array[0] = 100
    fmt.Println(s1)
    fmt.Println(s2)
    //[100 2 3 4 5]
    //[2 3]
}
```

#### range 的使用
``` GO
func main() {
  
      nums := []int{2, 3, 4}
      sum := 0
      for _, num := range nums {
          sum += num
      }
      fmt.Println("sum:", sum)
  
      for i, num := range nums {
          if num == 3 {
              fmt.Println("index:", i)
          }
          //index: 1 
      }
  
      kvs := map[string]string{"a": "apple", "b": "banana"}
      for k, v := range kvs {
          fmt.Printf("%s -> %s\n", k, v)
      }
      //两个变量的时候传递的是key - value
  
      for k := range kvs {
          fmt.Println("key:", k)
      }
      //只有一个的时候，传递的是 key
  
      for i, c := range "go" {
          fmt.Println(i, c)
      }
      //0 103
      // 1 111   本质因为字符串之中是 int32 存储的
  }


#### slice 相关库
######  slices.Equal(a,b)  return bool  判断是否相等
``` Go
func main() {
    array := []int{1, 2, 3, 4, 5}
    s1 := array
    res:=slices.Equal(s1, array)
    fmt.Println(res) // Output: true
}
```
比较的时候不会看切片的 cap 是不是一样的，就是看 len 中的

#### copy的使用
有多少就 copy 多少，(,)从后面的 copy 到前面的
``` Go
func main() {
	array := []int{1, 2, 3, 4, 5}
	s1 := make([]int, 6)
    s2:= make([]int, len(array))
    s3 := make([]int, 3)
	fmt.Println(len(s1), cap(s1)) // Output: 6,6
	copy(s1, array)
    copy(s2, array)
    copy(s3, array)
	fmt.Println(s1) // Output: [1 2 3 4 5 0]
    fmt.Println(s2) // Output: [1 2 3 4 5]
    fmt.Println(s3) // Output: [1 2 3]
	res := slices.Equal(s1, array)
    res2 := slices.Equal(s2, array)
	fmt.Println(res) // Output: false
    fmt.Println(res2) // Output: true
}
```



类似于 python 的字典
- key和value 键-值的对应
	- var mapname map\[keytype] valuetype  定义


#### 基本示例
``` Go
func main() {
	  
	  m := make(map[string]int)
	  //map will auto increase
  
	  m["k1"] = 7
	  m["k2"] = 13
  
	  fmt.Println("map:", m)
  
	  v1 := m["k1"]
	  fmt.Println("v1:", v1)
  
	  v3 := m["k3"]
	  fmt.Println("v3:", v3)
  
	  fmt.Println("len:", len(m))
  
	  delete(m, "k2")
	  fmt.Println("map:", m)
	  //delete certain key
	  clear(m)
	  fmt.Println("map:", m)
	  //clear all of the map
	  _, prs := m["k2"]
	  fmt.Println("prs:", prs)
	  //return the value and indication about the key's existence
	  //to disambiguate between missing keys and keys with zero values like 0 
	  n := map[string]int{"foo": 1, "bar": 2}
	  fmt.Println("map:", n)
  
	  n2 := map[string]int{"foo": 1, "bar": 2}
	  if maps.Equal(n, n2) {
		  fmt.Println("n == n2")
	  }
} 
```

#### map在并发读写中的问题
``` GO
aa := make(map[int]int)
	go func (	
		for {
			// aa[0] = 5  //不可以的，不支持并发写
			_ = aa[2]   //可以的，读操作不影响写操作
		}
	)()
	go func ()  {
		for {
			_ = aa[1]
		}
	}()
```

#### Map原理
###### map 结构
指定map长度为N，会初始化生成桶，桶数量为log2N
map在超过负载因子的时候会双倍重建，如果溢桶太大就会等量重建。当用到的时候旧桶才会放入新桶
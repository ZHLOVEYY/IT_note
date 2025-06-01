#### 基本例子
``` GO
package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d) // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha"))  // false
	fmt.Println(checkPassword2(&a, "haha")) // false
}

//不同传递方式
func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}
```

#### 传入数值和指针的区别
``` GO
ackage main

import "fmt"

func main() {
	//一个类型的所有方法最好使用相同的接收器类型（要么都是值，要么都是指针）
	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	u.ChangeName("Tom Changed!")
	u.ChangeAge(100)
	fmt.Printf("%v \n", u)

	//{Tom 100}

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age:  12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!")
	up.ChangeAge(120)

	//&{Jerry 120}

	fmt.Printf("%v \n", up)
}

type User struct {
	Name string
	Age  int
}

// 结构体接收器
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
} 
```
#### 初始化方法汇总
``` GO
func main(){
	 //定义struct的方法
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()

	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3 是 *ToyDuck
	duck3 := new(ToyDuck)
	duck3.Swim()

	// 当你声明这样的时候，Go 就帮你分配好内存
	// 不用担心空指针的问题，以为它压根就不是指针
	var duck4 ToyDuck
	duck4.Swim()

	// duck5 就是一个指针了
	//var duck5 *ToyDuck
	// 这边会直接panic 掉,因为指针需要初始化！
	//duck5.Swim()

	// 建议添加初始化示例
	var duck5 *ToyDuck = &ToyDuck{
		Color: "红色",
		Price: 200,
	}
	duck5.Swim()

	// 赋值，初始化按字段名字赋值
	duck6 := ToyDuck{
		Color: "黄色",
		Price: 100,
	}
	duck6.Swim()

	// 初始化按字段顺序赋值，不建议使用
	duck7 := ToyDuck{"蓝色", 1024}
	duck7.Swim()

	// 后面再单独赋值
	duck8 := ToyDuck{}
	duck8.Color = "橘色"

}

// ToyDuck 玩具鸭
type ToyDuck struct {
	Color string
	Price uint64
}

func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河，游过一群鸭，我是%s，%d一只\n", t.Color, t.Price)
} 
```



#### 结构体特点
###### 结构体方法定义基于另外一个结构体后，还是只能用自己的方法
``` GO
func main() {
	fake := FakeFish{}
	// fake 无法调用原来 Fish 的方法
	//如果有一样的swim方法也会调用自己的！！！
	// 这一句会编译错误
	//fake.Swim()
	fake.FakeSwim()

	// 转换为Fish
	td := Fish(fake)
	// 真的变成了鱼
	td.Swim()

}

// 定义了一个新类型，注意是新类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼，嘎嘎嘎\n")
}

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼，假装自己是一直鸭子\n")
} 
```

- 类型别名type fakeNews = News 
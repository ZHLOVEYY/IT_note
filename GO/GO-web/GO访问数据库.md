更多个人笔记：（仅供参考，非盈利）
gitee： https://gitee.com/harryhack/it_note
github： https://github.com/ZHLOVEYY/IT_note

本文是基于原生的库 database/sql进行初步学习
基于ORM等更多操作可以关注我的博客和笔记仓库
## 连接 MySQL 和基本 CRUD 操作
需要 `go get -u github.com/go-sql-driver/mysql`获取sql驱动
自己先连接到sql的root数据库，然后创建/fortest数据库

``` Go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 添加 MySQL 驱动。没有直接使用包的对象所以加_
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/fortest") //这里是user:password@tcp(location)/database的形式，我的密码是1234
	if err != nil {
		log.Fatal("连接失败", err)
	}
	defer db.Close()

	//测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal("连接失败", err)
	}
	fmt.Println("连接成功")

	//创建表  注意用if no exists避免重复创建
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gameusers (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50),
        age INT
    )`)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

	//插入数据
	result, err := db.Exec(`INSERT INTO gameusers (name, age) VALUES (?, ?)`, "Alice", 25)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId() //用于获取最新插入的ID
	fmt.Printf("插入成功,ID:%d\n", id)
	//再插入一条
	_, err = db.Exec(`INSERT INTO gameusers (name, age) VALUES (?, ?)`, "Alice2", 26)
	if err != nil {
		log.Fatal(err)
	}

	//查询单条数据
	var name string
	var age int //方便查询然后存储进来
	err = db.QueryRow("SELECT name, age FROM gameusers WHERE id = ?", id).Scan(&name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查询结果: name=%s, age=%d\n", name, age)

	//查询多条记录
	rows, err := db.Query("SELECT * FROM gameusers") //*就是所有列的了
	// rows, err := db.Query("SELECT id,name,age from gameusers") //显示指定列
	if err != nil {
		log.Fatal("查询失败", err)
	}
	defer rows.Close()
	//rows 是一个数据库结果集，会占用数据库连接和内存资源，如果不关闭，可能会导致资源泄露
	for rows.Next() { //类似于迭代器
		var userID int
		err := rows.Scan(&userID, &name, &age) //将当前行的数据读取到指定的变量中
		if err != nil {
			log.Fatal("扫描失败", err)
		}
		fmt.Printf("ID:%d ,姓名:%s,年龄:%d\n", userID, name, age)
	}

	//更新数据
	_, err = db.Exec("UPDATE gameusers SET age = ? WHERE id = ?", 26, id)
	if err != nil {
		log.Fatal("更新失败", err)
	}
	fmt.Println("更新成功")

	//删除数据
	_, err = db.Exec("DELETE FROM gameusers WHERE id =?", id)
	if err != nil {
		log.Fatal("删除失败", err)
	}
	fmt.Println("删除成功")

}

```

总结：
- open()：用于建立连接
	- 需要defer关闭数据库连接防止浪费资源
- ping()：用于测试是否建立连接
- exec()：用于执行创建表，CRUD等操作
- QueryRow()：用于查询单条数据
- Query()：用于多行查询
	- 注意结合.Next迭代和Scan写入
	- 需要defer关闭防止查询一直连接浪费资源


## 事务和结构体映射
事务的特点：
- 原子性：要么全部成功，要么全部失败
- 一致性：数据库从一个一致状态转换到另一个一致状态
- 隔离性：事务执行不受其他事务影响
- 持久性：一旦提交，修改就是永久的
例子理解：
- 假设你要给 A 转账 100 元给 B
- 需要两个操作：A 减 100，B 加 100
- 如果 A 减 100 成功，但 B 加 100 失败
- 这时就需要 Rollback，撤销 A 减 100 的操作
- 确保账户金额不会出错

``` GO
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 添加 MySQL 驱动。没有直接使用包的对象所以加_
	"log"
)

// User 结构体映射 users 表
type GameUser struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/fortest")
	if err != nil {
		log.Fatal("连接失败", err)
	}
	defer db.Close()

	//测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal("连接失败", err)
	}
	fmt.Println("连接成功")

	//创建表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gameusers (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50),
        age INT
    )`)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("事务开启失败", err)
	}

	//插入用户
	result, err := tx.Exec("INSERT INTO gameusers (name, age) VALUES (?, ?)", "Bob", 30)
	if err != nil {
		tx.Rollback()
		log.Fatal("插入失败", err)
	}
	id, _ := result.LastInsertId()
	fmt.Println("插入成功，用户ID:", id)

	//更新用户
	_, err = tx.Exec("UPDATE users SET age = ? WHERE id = ?", 31, id)
	if err != nil {
		tx.Rollback()
		log.Fatal("更新失败:", err)
	}

	// 提交事务
	err = tx.Commit() //通过commit进行提交
	if err != nil {
		log.Fatal("事务提交失败:", err)
	}
	fmt.Println("事务完成")

	//查询并映射到结构体
	var user GameUser
	err = db.QueryRow("SELECT id, name, age FROM gameusers WHERE id =?", id).Scan(&user.ID, &user.Name, &user.Age) //将查询结果映射到user结构体
	if err != nil {
		log.Fatal("查询失败", err)
	}
	fmt.Printf("用户信息: %+v\n", user)

}

```
执行完会发现报错：
``` bash
连接成功
插入成功，用户ID: 5
2025/04/25 10:12:03 更新失败:Error 1146 (42S02): Table 'fortest.users' doesn't exist
```
（这个用户id我插入很多次，会自己不断增加的）
这时候检查数据库会发现，bob并没有被插入，这是因为回滚了。那么接下来只要把上面代码中更新部分的users改为gameusers就可以了

如果不希望总是写tx.Rollback()
也可以增加：
``` GO
    // 定义一个函数，确保在出错时回滚事务
    defer func() {
        if err != nil {
            tx.Rollback()
            return
        }
    }()
```
这样就会在最后进行统一的判断

## 批量插入和预处理语句
预处理可以防sql注入，通过永远使用 ? 占位符，避免直接拼接 SQL 语句

``` Go
func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/fortest")
	if err != nil {
		log.Fatal("连接错误", err)
	}
	defer db.Close()

	//准备预处理
	stmt, err := db.Prepare("INSERT INTO gameusers (name,age) values (?,?)")
	if err != nil {
		log.Fatal("预处理失败", err)
	}
	defer stmt.Close()

	//批量插入
	users := []struct {   //定义结构体列表的同时初始化
		name string
		age  int
	}{
		{"Charlie", 22},
		{"David", 28},
		{"Eve", 19},
	}
	for _, u := range users {
		_, err := stmt.Exec(u.name, u.age) //通过Exec传递参数进去
		if err != nil {
			log.Fatal("批量插入失败", err)
		}
	}
	fmt.Println("批量插入完成")

	//查询所有用户
	rows, err := db.Query("SELECT id,name,age FROM gameusers")
	if err != nil {
		log.Fatal("查询失败", err)
	}
	defer rows.Close()
	
	for rows.Next() { //读取数据
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("读取是吧", err)
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", id, name, age)
	}
}
```
还有比如查询大数据时，限制返回行数（如 LIMIT），以及设置连接数和空闲连接数SetMaxOpenConns、SetMaxIdleConns等等，可以进一步拓展学习


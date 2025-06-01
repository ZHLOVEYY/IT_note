## 基本示例
``` GO
package main

import (
	"fmt"
	"html/template"
	"os"
)

// 定义数据结构
type Person struct {
	Name    string
	Age     int
	Email   string
	IsAdmin bool
}

type Company struct {
	Name      string
	Employees []Person
	Founded   int
}

func main() {
	// 示例1：基本语法演示
	basicTemplateDemo()

	// 示例2：条件判断和循环
	conditionalAndLoopDemo()

	// 示例3：HTML模板
	htmlTemplateDemo()

	// 示例4：Must的使用
	mustTemplateDemo()
}

// 基本语法演示
func basicTemplateDemo() {
	fmt.Println("=== 基本语法演示 ===")

	// {{.}} 表示当前数据对象
	tmpl1 := `Hello {{.}}!`
	t1 := template.Must(template.New("basic").Parse(tmpl1))
	t1.Execute(os.Stdout, "World")
	fmt.Println()

	// {{.Name}} 访问结构体字段
	person := Person{Name: "张三", Age: 30, Email: "zhangsan@example.com"}
	tmpl2 := `姓名: {{.Name}}, 年龄: {{.Age}}, 邮箱: {{.Email}}`
	t2 := template.Must(template.New("person").Parse(tmpl2)) //person 是模板名字
	t2.Execute(os.Stdout, person)   //对应的结构体
	fmt.Println("\n")
}

// 条件判断和循环演示
func conditionalAndLoopDemo() {
	fmt.Println("=== 条件判断和循环演示 ===")

	company := Company{
		Name:    "科技公司",
		Founded: 2020,
		Employees: []Person{
			{Name: "李四", Age: 25, Email: "lisi@example.com", IsAdmin: true},
			{Name: "王五", Age: 28, Email: "wangwu@example.com", IsAdmin: false},
			{Name: "赵六", Age: 32, Email: "zhaoliu@example.com", IsAdmin: false},
		},
	}

	// 使用 {{range .}} 循环和 {{if}} 条件判断
	// {{- 和 -}} 用于去除空白字符
	tmpl := `公司: {{.Name}} (成立于 {{.Founded}}年)
        员工列表:
        {{- range .Employees}}
        - 姓名: {{.Name}}
        {{- if .IsAdmin}} (管理员){{else}} (普通员工){{end}} 
        年龄: {{.Age}}, 邮箱: {{.Email}}
        {{- end}}
        `
        //通过IsAdmin 进行判断是否是管理员

	t := template.Must(template.New("company").Parse(tmpl))
	t.Execute(os.Stdout, company) //传入的 company 是一个大结构体，里面包括 Employees 列表
	fmt.Println()
}

// HTML模板演示
func htmlTemplateDemo() {
	fmt.Println("=== HTML模板演示 ===")

	// HTML模板会自动转义特殊字符
	htmlTmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Name}} - 员工列表</title>
</head>
<body>
    <h1>{{.Name}}</h1>
    <p>成立年份: {{.Founded}}</p>
    
    <h2>员工信息</h2>
    <table border="1">
        <tr>
            <th>姓名</th>
            <th>年龄</th>
            <th>邮箱</th>
            <th>角色</th>
        </tr>
        {{- range .Employees}}
        <tr>
            <td>{{.Name}}</td>
            <td>{{.Age}}</td>
            <td>{{.Email}}</td>
            <td>
                {{- if .IsAdmin -}}
                    <span style="color: red;">管理员</span>
                {{- else -}}
                    普通员工
                {{- end -}}
            </td>
        </tr>
        {{- end}}
    </table>
</body>
</html>
`

	company := Company{
		Name:    "示例科技公司",
		Founded: 2020,
		Employees: []Person{
			{Name: "张三", Age: 30, IsAdmin: true},
			{Name: "李四", Age: 25, IsAdmin: false},
		},
	}

	// 创建HTML文件
	file, err := os.Create("company.html")  //创建接收的文件
	if err != nil {
		panic(err)
	}
	defer file.Close()

	t := template.Must(template.New("html").Parse(htmlTmpl))
	t.Execute(file, company)  //这里不是 os.stdout输出了，而是输出到了文件中
	fmt.Println("HTML文件已生成: company.html")
}

// Must的使用演示
func mustTemplateDemo() {
	fmt.Println("=== Must使用演示 ===")

	// Must会在模板解析失败时直接panic
	// 这是一个正确的模板
	validTemplate := `Hello {{.Name}}!`
	t1 := template.Must(template.New("valid").Parse(validTemplate))

	person := Person{Name: "测试用户"}
	fmt.Print("正确模板输出: ")
	t1.Execute(os.Stdout, person)
	fmt.Println()

	// 演示错误处理（注释掉以避免程序崩溃）
	/*
		// 这是一个错误的模板语法，Must会直接panic
		invalidTemplate := `Hello {{.Name!` // 缺少右括号
		t2 := template.Must(template.New("invalid").Parse(invalidTemplate))
	*/

	// 不使用Must的错误处理方式
	invalidTemplate := `Hello {{.Name!`  //这里缺少了右边括号    
	t2, err := template.New("invalid").Parse(invalidTemplate) //没有用 Must，用 Must 会 panic 导致程序退出
	if err != nil {
		fmt.Printf("模板解析错误: %v\n", err)
	} else {
		t2.Execute(os.Stdout, person)
	}
}

```


## 语法
### 1. 基本输出语法
- `{{.}}`: 输出当前数据对象
- `{{.Name}}`: 输出结构体的Name字段
- `{{.Field.SubField}}`: 访问嵌套字段
### 2. 控制结构
- `{{if .IsAdmin}}...{{else}}...{{end}}`: 条件判断
- `{{range .Employees}}...{{end}}`: 循环遍历
- `{{with .Field}}...{{end}}`: 设置上下文
### 3. 空白字符控制
- `{{-`: 去除左侧空白字符
- `-}}`: 去除右侧空白字符
- `{{- if .IsAdmin -}}`: 去除两侧空白字符
### 4. Must函数
- `template.Must()`: 在模板解析失败时直接panic
- 适用于程序启动时的模板初始化
- 类似于panic的检验机制
- 如果需要接受错误就不要用 Must

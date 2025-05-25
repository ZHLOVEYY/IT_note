## 基本例子

``` GO
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// 1. 获取当前进程信息
	fmt.Println("--- 进程信息 ---")
	fmt.Println("进程ID:", os.Getpid())
	fmt.Println("父进程ID:", os.Getppid())
	fmt.Println("用户ID:", os.Getuid())
	fmt.Println("组ID:", os.Getgid())

	// 2. 获取系统信息
	fmt.Println("\n--- 系统信息 ---")
	fmt.Println("操作系统:", runtime.GOOS)
	fmt.Println("CPU核心数:", runtime.NumCPU())
	hostname, _ := os.Hostname()
	fmt.Println("主机名:", hostname)

	// 3. 环境变量操作
	fmt.Println("\n--- 环境变量 ---")
	fmt.Println("PATH:", os.Getenv("PATH"))
	os.Setenv("TEST_ENV", "test_value")
	fmt.Println("TEST_ENV:", os.Getenv("TEST_ENV"))

	// 4. 执行系统命令
	fmt.Println("\n--- 执行命令 ---")
	cmd := exec.Command("echo", "Hello, Go!")
	output, _ := cmd.Output()
	fmt.Printf("命令输出: %s", output)

	// 5. 文件系统操作
	fmt.Println("\n--- 文件操作 ---")
	_, err := os.Stat("test.txt")
	if os.IsNotExist(err) {
		fmt.Println("创建test.txt文件")
		os.WriteFile("test.txt", []byte("测试内容"), 0644)
	} else {
		data, _ := os.ReadFile("test.txt")
		fmt.Println("文件内容:", string(data))
	}

	// 6. 退出进程
	fmt.Println("\n--- 进程退出 ---")
	defer fmt.Println("清理工作...") // defer语句会在函数退出前执行
	// os.Exit(0) // 立即退出，不执行defer
	// syscall.Exit(0) // 系统调用方式退出

	// 7. 创建子进程
	fmt.Println("\n--- 创建子进程 ---")
	attr := &os.ProcAttr{ //创建ProcAttr结构体定义子进程属性
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //Files字段设置子进程的标准输入/输出/错误流，这里复用父进程的IO
	}
	process, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, attr) //启动/bin/ls程序执行ls -l命令
	if err != nil {
		fmt.Println("启动失败:", err)
		return
	}
	fmt.Println("子进程ID:", process.Pid) //输出子进程 pid
	state, _ := process.Wait()
	fmt.Println("子进程退出状态:", state.Success()) //检查退出状态

}

```


1. `os.Getpid()` - 获取当前进程ID
2. `exec.Command()` - 执行系统命令
3. `os.Stat()` - 检查文件状态
4. `os.StartProcess()` - 创建子进程
5. `os.Getenv()`/`Setenv()` - 环境变量操作
6. `runtime`包 - 获取运行时信息
7. `signal`包 - 处理系统信号(示例中已注释)
8. `os.Exit()` - 控制进程退出
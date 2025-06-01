## if else
注意格式就行
``` go
package main

import "fmt"

func main() {
	a := 1
	b := 2
	if a > b {
		fmt.Println("a > b")
	} else if a == b {
		fmt.Println("a = b")
	} else {
		fmt.Println("a < b")
	}
}
```

## switch代码
``` GO
package main

import "fmt"

func main() {
	tag := "h"
	switch tag {
	case "h":
		fmt.Println("高")
	case "m":
		fmt.Println("中")
	case "l":
		fmt.Println("低")
	default:
		fmt.Println("未知")
	}
}
```
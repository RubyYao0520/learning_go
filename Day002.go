//2025/8/15
//简单的控制结构

package main

import "fmt"

func Day002() {
	var ans int
	for i := 1; i <= 10; i++ {
		ans += i
	}
	//多个变量
	for i, j := 0, 10; i < j; {
		ans += i + j
		i++
		j--
	}
	fmt.Println(ans)
	var c string
	//在使用if-else结构时发现对代码的格式也有要求
	//go语言中for循环
	fmt.Print("Please input string: ")
	fmt.Scan(&c)
	if len(c) > 10 {
		fmt.Println("length > 10")
	} else {
		fmt.Println("length <= 10")
	}

	//不用break
	var inputVar int
	fmt.Print("Input variable: ")
	fmt.Scan(&inputVar)
	switch inputVar {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}

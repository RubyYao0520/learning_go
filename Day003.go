//2025/8/17
//数组

package main

import "fmt"

func Day003() {
	//定义一个数组
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i * i
	}
	//打印下标和元素
	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Println("----------")
	//只打印下标
	for v := range a {
		fmt.Println(v)
	}
	fmt.Println("----------")
	//只打印值
	for _, v := range a {
		fmt.Println(v)
	}
}

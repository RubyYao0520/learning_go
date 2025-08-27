//2025/8/27
//继续学习复合数据类型

package main

import (
	"fmt"
)

func Day004() {
	fmt.Println("This is func Day004")
	//尝试使用数组的不同初始化方式
	//数组的长度是固定的
	q := []int{1, 2, 3}
	fmt.Println(Sum(q))
}

func Sum(arr []int) int {
	var s int = 0
	for _, value := range arr {
		s += value
	}
	return s
}

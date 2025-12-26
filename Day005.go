package main

//2025/12/26
//指针数组与数组指针

import (
	"fmt"
)

func Day005() {

	k := []byte{1, 1, 1}
	fmt.Print(k)

	// 这是指针数组 - 存储3个byte指针的数组
	var p [3]*byte

	var i int
	for i = 0; i < 3; i++ {
		p[i] = &k[i]
	}
	zero(p)
	fmt.Println("使用指针数组清零后:", k)

	fmt.Println("\n使用数组指针")
	k2 := []byte{1, 1, 1}
	fmt.Println("清零前:", k2)

	// 将切片转换为数组指针
	var p2 *[3]byte
	p2 = (*[3]byte)(k2) // 这是关键：将切片转换为数组指针
	zero2(p2)

	fmt.Println("使用数组指针清零后:", k2)

}

func zero(ptr [3]*byte) {
	for i := range ptr {
		*ptr[i] = 0
	}
}

// 使用数组指针直接操作整个数组
func zero2(ptr *[3]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

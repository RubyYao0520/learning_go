package main

import "fmt"

//2026/03/15

func Day011() {
	//slice的声明方式

	//声明slice1是一个切片，并且初始化，默认值是{1,2,3},长度是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//声明slice是一个切片，但是没有分配空间
	var slice2 []int
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	//使用make分配三个空间
	slice2 = make([]int, 3)

	var slice3 = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	//判断一个slice是否为0
	if slice2 == nil {
		fmt.Println("slice2 is nil")
	} else {
		fmt.Println("slice2 is not nil")
	}

	//这里numbers一共分配了三个空间，但是其容量是5
	//只是尾指针指向3,3之后的空间是不可访问的
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d", len(numbers))

	//向numbers追加元素
	numbers = append(numbers, 1)
	//此时 len = 4, cap = 5
	fmt.Printf("len = %d", len(numbers))

	//如果向一个容量已满的slice追加元素，会开辟新的内存空间，大小为cap

	//copy 可以将底层数组的slice一起进行拷贝

}

package main

//2026/03/14

//匿名导入，无法使用导入包的方法，但是会使用init方法
//给导入的包起别名
import (
	"fmt"
	_ "learning/lib1"
	//以下方式导入之后不用通过包名调用，可以直接使用方法名调用
	//. "learning/lib1"
	mylib2 "learning/lib2"
)

func Day010() {
	fmt.Println("Day010")
	mylib2.Lib2Test()

	var a int = 10
	var b int = 20

	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)

	//固定长度数组
	var myArray1 [10]int
	//遍历数组
	for i := 0; i < 10; i++ {
		fmt.Println(myArray1[i])
	}

	myArray2 := [10]int{1, 2, 3, 4}
	for index, value := range myArray2 {
		fmt.Println("index ", index, "value ", value)
	}

	//查看数组的数据类型
	fmt.Printf("arr1 type =%T\n", myArray1)
}

// 引用传值
func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

package main

import "fmt"

//2026/03/15

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func printMap(cityMap map[string]string) {
	//参数是引用传递
	for key, value := range cityMap {
		fmt.Print(key, " ")
		fmt.Println(value)
	}
}

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

	//map的声明
	//key是string，value是string

	//第一种声明方式
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is nil")
	}
	//开辟空间
	myMap1 = make(map[string]string, 10)

	//扩容机制与slice相同
	myMap1["1"] = "Python"
	myMap1["2"] = "Java"
	myMap1["3"] = "Rust"
	fmt.Println(myMap1)

	fmt.Println("---------")

	//第二种声明方式
	myMap2 := make(map[int]string, 5)
	myMap2[1] = "Java"

	//第三种声明方式
	myMap3 := map[string]string{
		"1": "Python",
		"2": "Java",
	}
	fmt.Println(myMap3)

	cityMap := map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"Korea": "Seoul",
		"USA":   "New York",
	}

	for key, value := range cityMap {
		fmt.Println(key)
		fmt.Println(value)
	}

	//修改
	cityMap["USA"] = "DC"
	//删除
	delete(cityMap, "Japan")
	//遍历
	for _, value := range cityMap {
		fmt.Println(value)
	}
	printMap(cityMap)

	var book1 Book
	book1.title = "nil"
	book1.author = "zhangsan"
	fmt.Println(book1)

}

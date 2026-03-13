package main

//2026/03/13

import "fmt"

const (
	// 可以在const()添加一个关键字 iota，每一行的iota都会累加1,第一行的iota的默认值是0
	BEIJING = iota
	SHANGHAI
	SHENZHEN
	GUANGZHOU
)

// iota只能出现在const的括号里
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f = iota * 3, iota * 4
	g, h
)

func Day009() {
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("++++++++++++++++++++")

	f1("hello", 999)
	r1, r2 := f2("haha", 111)
	fmt.Println(r1, r2)

	r3, r4 := f3("hahs", 222)
	fmt.Println(r3, r4)
}

func f1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r := 0
	return r
}

// 返回多个返回值，匿名的
func f2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 111, 222
}

// 返回多个返回值，有形参名称的
func f3(a string, b int) (ret1 int, ret2 int) {
	fmt.Println(a, b)

	return ret1, ret2
}

//init函数在main函数之前执行

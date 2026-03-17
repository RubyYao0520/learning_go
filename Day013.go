package main

import "fmt"

//2026/03/17

// 变量的内部构造
func main() {
	var a string
	//pair<type:string, value:"acde">
	a = "acde"

	//pair<type:string, value:"acde">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}

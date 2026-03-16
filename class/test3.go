package main

import "fmt"

//通用万能类型

// arg就是一个万能类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is running")
	fmt.Println(arg)

	//给 interface{} 提供“类型断言”机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println(value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"GoLang"}
	myFunc(book)
	myFunc(3.14)
}

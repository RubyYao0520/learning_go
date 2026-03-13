package lib1

import "fmt"

// lib1包提供接口
func Lib1Test() {
	fmt.Println("lib1 test")
}

func init() {
	fmt.Println("lib1.init() ...")
}

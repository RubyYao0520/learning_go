package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("function main")
	//Day009()
	//lib1.Lib1Test()
	//lib2.Lib2Test()

	//写入defer关键字
	//表示函数调用结束前
	//多个defer使用栈的形式
	//defer在return之后调用
	defer fmt.Println("main end")
	defer fmt.Println("main end2")
	//Day010()
	Day011()
}

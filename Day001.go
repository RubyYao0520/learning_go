//2025/8/14
//基本的输入和输出
//命名变量

package main

//导入格式化输入输出库
import (
	"fmt"
)

func Day001() {
	//var关键字+变量名+变量类型
	var a int
	var b int

	//使用Scan读入
	fmt.Scan(&a, &b)
	fmt.Print("a+b=", a+b)

}

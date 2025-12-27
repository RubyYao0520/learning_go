package main

//2025/12/27

import (
	"fmt"
	"sort"
)

// Slice结构
func Day006() {
	//这里数组代表索引
	//跳过了0这个下标
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)

	//左闭右开区间
	//多个Slice之间可以共享底层数据
	summer := months[4:7]
	month := months[0:1]
	fmt.Println("现在输出0~1下标的内容")
	fmt.Println(month)
	fmt.Println(len(summer))
	fmt.Println(summer)

	months2 := [...]string{"Jan", "Feb", "Mar"}
	fmt.Println(months2)

	//append添加元素
	summer2 := months2[0:2]
	fmt.Println(summer2)
	summer2 = append(summer2, "May")
	fmt.Print("添加元素后 ")
	fmt.Println(summer2)

	//Map
	//使用内置函数make可以创建一个Map
	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2
	ages["c"] = 3
	ages["d"] = 4
	ages["f"] = 10
	ages["e"] = 9
	fmt.Println(ages)
	//使用delete删除Map中的元素
	delete(ages, "a")
	fmt.Print("第一遍遍历:")
	fmt.Println(ages)
	fmt.Print("第二遍遍历:")
	fmt.Println(ages)
	//Map的迭代顺序是不确定的
	var characters []string
	for char := range ages {
		characters = append(characters, char)
	}
	//对key排序之后可以按顺序遍历Map
	sort.Strings(characters)
	fmt.Print("排序后遍历:")
	fmt.Println(ages)
}

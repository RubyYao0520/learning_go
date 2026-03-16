package main

import "fmt"

//2026/03/16

// 面向对象
// 主要是通过结构体来完成
// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果类的属性首字母大写，表示该属性能够对外访问，否则智能对内访问
	Name  string
	Ad    int
	Level int
}

/*func (this Hero) Show() {
	fmt.Println(this)
}

// 表示绑定到当前的结构体上
func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(Name string) {
	//this 是调用该方法的对象的一个副本
	this.Name = Name
}*/

// 方法首字母大写表示可以被其他包访问
func (this *Hero) Show() {
	fmt.Println(this)
}

// 表示绑定到当前的结构体上
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(Name string) {
	//this 是调用该方法的对象的一个副本，所以要使用指针
	this.Name = Name
}

func Day012() {
	//创建对象
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("lisi")
	hero.Show()
}

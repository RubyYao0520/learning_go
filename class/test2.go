package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eats")
}

func (this *Human) Walk() {
	fmt.Println("Human walks")
}

// 继承Human类的方法
type SuperMan struct {
	Human
	level int
}

// 重新定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat")
}

// 定义子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly")
}

func main() {
	h := Human{name: "Jack", sex: "male"}
	h.Eat()
	h.Walk()

	//定义一个子类对象
	s := SuperMan{Human{"Alice", "female"}, 2}
	//父类的方法
	s.Walk()
	s.Eat()
	s.Fly()

}

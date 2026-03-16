package main

import "fmt"

// 本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string //获取动物的种类
}

// 具体的类
// 继承接口只需要实现接口里面的方法
// 必须实现接口里面每一个方法，不能部分实现！
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal Animal) {
	animal.Sleep()
}

func main() {

	//接口的数据类型、父类指针
	/*var animal Animal
	animal = &Cat{"Green"}
	//此时调用的就是Cat的Sleep()方法
	animal.Sleep()
	animal = &Dog{"Yellow"}
	animal.Sleep()*/
	cat := Cat{"red"}
	showAnimal(&cat)
}

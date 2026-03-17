package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type TextBook struct {
}

func (this *TextBook) ReadBook() {
	fmt.Println("Read a book")
}

func (this *TextBook) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	b := &TextBook{}

	var r Reader
	r = b

	r.ReadBook()
	var w Writer
	w = r.(Writer)
	w.WriteBook()
}

package main

//2026/03/12

import (
	"fmt"
	"time"
)

func goFunc(i int) {
	fmt.Println("go", i)
}

func Day008() {
	for i := 0; i < 20; i++ {
		//并发
		go goFunc(i)
	}

	time.Sleep(time.Second)
}

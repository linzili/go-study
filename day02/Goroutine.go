package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")

	for i := 0; i < 10; i++ {
		go fmt.Println(i)

		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("end")

	//go fmt.Println("hello world")
	//go hello()
	//go func() {
	//	fmt.Println("hello world")
	//}()
}

//func hello() {
//	fmt.Println("hello world")
//}

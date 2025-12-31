package main

import (
	"fmt"
	"time"
)

//func main() {
//	// 创建三个管道
//	chA := make(chan int)
//	chB := make(chan int)
//	chC := make(chan int)
//	defer func() {
//		close(chA)
//		close(chB)
//		close(chC)
//	}()
//
//	//go func() {
//	//	chA <- 1
//	//}()
//
//	l := make(chan struct{})
//
//	go Send(chA)
//	go Send(chB)
//	go Send(chC)
//	go func() {
//	Loop:
//		for {
//			select {
//			case n, ok := <-chA:
//				fmt.Println(n, ok)
//			case n, ok := <-chB:
//				fmt.Println(n, ok)
//			case n, ok := <-chC:
//				fmt.Println(n, ok)
//			case <-time.After(time.Second):
//				break Loop
//			}
//		}
//		l <- struct{}{}
//	}()
//	<-l
//}

func Send(ch chan<- int) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond)
		ch <- i
	}
}

func main() {
	chA := make(chan int)

	defer close(chA)

	go func() {
		time.Sleep(time.Second * 2)
		chA <- 1
	}()

	select {
	case n := <-chA:
		fmt.Println(n)
	case <-time.After(time.Second):
		fmt.Println("超时")
	}
}

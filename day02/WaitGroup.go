package main

import (
	"fmt"
	"sync"
)

//func main() {
//
//	var wait sync.WaitGroup
//
//	wait.Add(1)
//
//	go func() {
//		fmt.Println("hello world")
//		wait.Done()
//	}()
//
//	wait.Wait()
//	fmt.Println("end")
//}

//func main() {
//	var mainWait sync.WaitGroup
//	var wait sync.WaitGroup
//
//	mainWait.Add(10)
//	fmt.Println("start")
//	for i := 0; i < 10; i++ {
//		wait.Add(1)
//		go func() {
//			fmt.Println(i)
//			wait.Done()
//			mainWait.Done()
//		}()
//		wait.Wait()
//	}
//	mainWait.Wait()
//
//	fmt.Println("end")
//
//}

func main() {
	var mainWait sync.WaitGroup
	mainWait.Add(1)
	hello(mainWait)
	mainWait.Wait()
	fmt.Println("end")

}

func hello(wait sync.WaitGroup) {
	fmt.Println("start")
	wait.Done()
}

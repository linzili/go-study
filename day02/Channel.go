package main

import "fmt"

// no cache
//func main() {
//	intCh := make(chan int)
//	defer close(intCh)
//	go func() { intCh <- 123 }()
//	fmt.Println(<-intCh)
//}

// cache
/*func main() {
	ch := make(chan int, 5)
	chW := make(chan struct{})
	chR := make(chan struct{})

	defer func() {
		close(ch)
		close(chW)
		close(chR)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("写入", i)
		}
		chW <- struct{}{}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println("读取", <-ch)
		}
		chR <- struct{}{}
	}()

	fmt.Println("写入完毕", <-chW)
	fmt.Println("读取完毕", <-chR)

}
*/

//func main() {
//	ch := make(chan struct{})
//	defer close(ch)
//	go func() {
//		fmt.Println(2)
//		ch <- struct{}{}
//	}()
//	<-ch
//	fmt.Println(1)
//
//}

//func main() {
//	ch := make(chan int, 10)
//	go func() {
//		for i := 0; i < 10; i++ {
//			ch <- i
//		}
//		close(ch)
//	}()
//	for n := range ch {
//		fmt.Println(n)
//	}
//}

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	// 关闭管道
	close(ch)
	// 再读取数据
	for i := 0; i < 6; i++ {
		n, ok := <-ch
		fmt.Println(n, ok)
	}
}

// 互斥锁
var count = 0

var lock = make(chan struct{}, 1)

func Add() {
	lock <- struct{}{}
	fmt.Println("当前计数为", count, "执行加法")
	count++

	<-lock
}

func Sub() {
	lock <- struct{}{}
	fmt.Println("当前计数为", count, "执行减法")
	count--
	<-lock
}

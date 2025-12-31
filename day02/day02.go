package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//nums = nums[2:]
	fmt.Println(nums)

	nums = nums[:2]

	fmt.Println(nums)
	//nums = nums[2 : len(nums)-2]

	//fmt.Println(nums)
	//fmt.Println("hello world")
	//ch := make(chan int)
	//
	//go func() {
	//	fmt.Println("run")
	//	ch <- 1
	//	close(ch)
	//}()
	//
	//time.Sleep(time.Second)

	//ch := make(chan int, 10)
	//x := <-ch

	//fmt.Println(x)

	//fmt.Println(test())
}

func test() (x int) {

	x = 1

	defer func() {
		x *= 2
	}()

	x++
	return x
}

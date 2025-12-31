package main

import "fmt"

func main() {
	//sum := Sum(1, 2)
	//sum := Sum(1)
	//fmt.Println(sum(2, 3))
	//fmt.Println(sum(5, 4))
	//fmt.Println(sum(5, 6))

	fib := Fib()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func Fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func Sum(sum int) func(int, int) int {
	return func(a, b int) int {
		sum += a + b
		return sum
	}
}

//func Sum(a, b int) func(int, int) int {
//	return func(int, int) int {
//		return a + b
//	}
//}

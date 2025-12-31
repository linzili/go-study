package main

import "fmt"

//func Do() {
//
//	defer func() {
//		fmt.Println("1")
//	}()
//	fmt.Println("2")
//}

func Do() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("3")
	}()
	defer func() {
		fmt.Println("4")
	}()
	fmt.Println("2")
	defer func() {
		fmt.Println("5")
	}()

}
func main() {
	Do()
}

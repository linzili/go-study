package main

import "fmt"

type IntSlice []int

func (i IntSlice) Get(index int) int {
	return i[index]
}

func (i IntSlice) Set(index, val int) {
	i[index] = val
}

func (i IntSlice) Len() int {
	return len(i)
}

type MyInt int

func (i *MyInt) Set(val int) {
	*i = MyInt(val)
}

type Animal interface {
	Run()
}

type Dog struct {
}

func (d *Dog) Run() {
	fmt.Println("Dog Run")
}

func main() {
	var an Animal
	an = &Dog{}
	an.Run()
	//var intSlice IntSlice
	//intSlice = []int{1, 2, 3, 4, 5}
	//fmt.Println(intSlice.Get(0))
	//intSlice.Set(0, 2)
	//fmt.Println(intSlice)
	//fmt.Println(intSlice.Len())

	//myInt := MyInt(1)
	//myInt.Set(2)
	//fmt.Println(myInt)
}

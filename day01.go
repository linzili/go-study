package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sort"
)

type User struct {
	Id   int
	Name string
}

type Reader interface {
	Read() string
}

type File struct{}

func (f File) Read() string {
	return "file"
}

//func add(a int, b int) int {
//	return a + b
//}

func add(a, b int) int {
	return a + b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func sum(a, b int) (res int) {
	res = a + b
	return
}

func calc(f func(int, int) int) int {
	return f(5, 10)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func get() int {
	return rand.IntN(100)
}

func main() {

	if v := get(); v > 0 {
		fmt.Println(v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	var x = get()
	//for x < 10 {
	//	x++
	//	fmt.Println(x)
	//}

	//for {
	//	fmt.Println(x)
	//}

	switch x {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	switch {
	case x > 10:
		fmt.Println("big")

	case x > 0:
		fmt.Println("zero")

	}

	array := []int{3, 1, 5, 9, 4}

	fmt.Println(array[:])
	fmt.Println(array[0:2])
	fmt.Println(array[1:2])
	sort.Ints(array)
	fmt.Println(array)
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	fmt.Println(array)

	//var x = calc(add)
	//if x > 0 {
	//	fmt.Println(x)
	//}
	//a := 10
	//p := &a
	//*p = 20
	//fmt.Println(a)
	//fmt.Println(*p)

	//var f File
	//fmt.Println(f.Read())

	//m := make(map[string]int)
	//m["a"] = 1
	//m["b"] = 2
	//fmt.Println(m)
	//
	//v, ok := m["a"]
	//
	//fmt.Println(v, ok)
	//s := []int{1, 2, 3}

	//s := make([]int, 0, 10)
	//s = append(s, 1)

	//fmt.Println(s)
	//fmt.Println("hello world")
	//var a int = 10
	//var b = 20
	//c := 30
	//
	//var (
	//	x int    = 20
	//	y string = "hello"
	//)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(x)
	//fmt.Println(y)
	//const PI = 3.14
	//
	//fmt.Println(PI)
	//u1 := User{Id: 1, Name: "Tom"}
	//u1.Name = "Jack"
	//
	//u2 := &User{}
	//u2.Name = "Tom"
	//
	//fmt.Println(u1)
	//fmt.Println(u2)
	//fmt.Println(reflect.TypeOf(u1))
	//fmt.Println(reflect.TypeOf(u2))

}

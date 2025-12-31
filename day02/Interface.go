package main

import "fmt"

type Person interface {
	Say(string) string
	Walk(int)
}

type Crane interface {
	JackUp() string
	Hoist() string
}

type CraneA struct {
	work int
}

func (c CraneA) Work() {
	fmt.Println("使用技术A")
}

func (c CraneA) JackUp() string {
	c.Work()
	return "jackup"
}

func (c CraneA) Hoist() string {
	c.Work()
	return "hoist"
}

type CraneB struct {
	boot string
}

func (c CraneB) Boot() {
	fmt.Println("使用技术B")
}

func (c CraneB) JackUp() string {
	c.Boot()
	return "jackup"
}

func (c CraneB) Hoist() string {
	c.Boot()
	return "hoist"
}

type ConstructionCompany struct {
	Crane Crane
}

func (c *ConstructionCompany) Build() {
	fmt.Println(c.Crane.JackUp())
	fmt.Println(c.Crane.Hoist())
	fmt.Println("建筑完成")
}

type Func func()

func (f Func) Say(s string) string {
	f()
	return "bibibibi"
}

func (f Func) Walk(i int) {
	f()
	fmt.Println("can not walk")
}

func main() {

	var function Func
	function = func() {
		fmt.Println("do something")
	}
	function()
	function.Walk(1)
	fmt.Println(function.Say("hello"))
	//var person Person
	//fmt.Println(person)

	//company := ConstructionCompany{CraneA{}}
	//company.Build()
	//fmt.Println()

	//company.Crane = CraneB{}
	//company.Build()
}

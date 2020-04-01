package main

import "fmt"

// interface接口

// 引出接口的实例

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("嘤嘤嘤~")
}

func da(x speaker) {
	// 接收一个参数，传进来什么，我就打什么
	x.speak() // 挨打了就要叫
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker // 定义一个接口类型：speaker 的变量：ss
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}

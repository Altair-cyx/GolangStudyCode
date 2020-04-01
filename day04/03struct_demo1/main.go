package main

import "fmt"

// struct结构体

type person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

func main() {
	// 声明一个person类型的变量p
	var p person
	// 通关字段赋值
	p.name = "周林"
	p.age = 9000
	p.sex = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	// 访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	var p2 person
	p2.name = "周林1"
	p2.age = 9000
	p2.sex = "男1"
	p2.hobby = []string{"篮球1", "足球1", "双色球"}

	fmt.Println(p2)
	// 访问变量p的字段
	fmt.Printf("%T\n", p2)
	fmt.Println(p2.name)

	// 匿名结构体：多用于临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "嘿嘿嘿"
	s.age = 100
	fmt.Printf("type:%T value:%v\n", s, s)
}

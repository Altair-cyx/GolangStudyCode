package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
	workPlace
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "周林",
		age:  9000,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	// fmt.Println(p1.name, p1.addr.city)
	fmt.Println(p1.name, p1.address.city, p1.workPlace.city)
}

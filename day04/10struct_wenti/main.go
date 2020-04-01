package main

import "fmt"

// 结构体遇到的问题

// 1.myInt(100)是个啥？
type myInt int

type person struct {
	name string
	age  int
}

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	// 声明一个int32类型的变量x，他的值是10
	// 方法1：
	// var x int32
	// x = 10
	// 方法2：
	// var x int32 = 10
	// 方法3：
	// var x = int32(10)
	// 方法4：
	// x := int32(10)
	// fmt.Println(x)

	// 声明一个myInt类型的变量m，他的值是100
	// 方法1：
	// var m myInt
	// m = 100
	// 方法2：
	// var m myInt = 100
	// 方法3：
	// var m = myInt(100)
	// 方法4：
	// m := myInt(100) // 强制类型转换
	// fmt.Println(m)

	// 问题2：结构体初始化
	// type person struct {
	// 	name string
	// 	age  int
	// }
	// 方法1：
	var p person // 声明一个person类型的变量p
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)

	var p1 person
	p1.name = "周林"
	p1.age = 9000
	fmt.Println(p1)
	// 方法2：
	// 键值对初始化
	var p2 = person{
		name: "鸡冠花",
		age:  20,
	}
	fmt.Println(p2)
	// 值列表初始化
	var p3 = person{
		"理想",
		200,
	}
	fmt.Println(p3)
}

// 问题3：为什么使用构造函数
func newPerson(name string, age int) person {
	// 别人调用我，我能给他一个person类型的变量
	return person{
		name: name,
		age:  age,
	}
}

// func newPerson(name string, age int) *person {
// 	// 别人调用我，我能给他一个person类型的变量
// 	return person{
// 		name: name,
// 		age:  age,
// 	}
// }

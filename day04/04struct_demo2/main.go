package main

import "fmt"

// 结构体是值类型

type person struct {
	name, sex string
}

// Go语言中函数参数永远是副本
func f(x person) {
	x.sex = "女" // 修改的是副本的gender
}

func f2(x *person) {
	(*x).sex = "女" // 根据内存地址找到那个原变量，修改的就是原来的变量
	x.sex = "女"    // 语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "周林"
	p.sex = "男"
	fmt.Println(p.sex)
	f(p)
	fmt.Println(p.sex)
	f2(&p)
	fmt.Println(p.sex)

	// 结构体指针1
	var p2 = new(person)
	p2.name = "理想"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)
	fmt.Printf("%p\n", &p2)

	// 结构体指针2
	// key-value初始化
	var p3 = &person{
		name: "元帅",
		sex:  "男",
	}
	fmt.Printf("%#v\n", p3)

	// 使用值列表的形式初始化，值得顺序要和结构体定义时字段顺序一致
	p4 := &person{
		"小王子", "男",
	}
	fmt.Printf("%#v\n", p4)
}

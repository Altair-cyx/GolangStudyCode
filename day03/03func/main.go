package main

import "fmt"

// 函数：一段代码的封装

func f1() {
	fmt.Println("Hello 沙河！")
}

func f2(name string) {
	fmt.Println("Hello", name)
}

func f3(x int, y int) int {
	sum := x + y
	return sum
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(x string, y ...int) int {
	fmt.Println(y)
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y
	return
}

// Go语言中支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("理想")
	f2("姬无命")
	fmt.Println(f3(100, 200))

	f5("lixiang", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

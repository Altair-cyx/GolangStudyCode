package main

import (
	"fmt"
)

// goroutine

// GMP
// G:goroutine
// M:操作系统虚拟线程
// P:管理goroutine分配给M

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 1000; i++ {
		// go hello(i) // 开启一个单独的goroutine去执行hello函数（任务）
		go func(i int) {
			fmt.Println(i) // 用的是函数参数的i，不是外面的i
		}(i)
	}
	fmt.Println("main")
	// time.Sleep(time.Second)
	// main函数结束了 由main函数启动的goroutine也都结束了
	// goroutine对应的函数结束了，goroutine结束了
}

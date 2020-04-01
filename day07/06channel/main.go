package main

import (
	"fmt"
	"sync"
)

// channel
// CSP，提供通关通信共享内存而不是通过共享内存而实现通信
// goroutine是Go语言并发的执行体，channel就是它们之间的连接，channel是可以让一个goroutine

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)     // nil
	b = make(chan int) // 无缓冲区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 // hang住了
	fmt.Printf("10发送到通道b中了...\n")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)         // nil
	b = make(chan int, 10) // 带缓冲区通道的初始化
	b <- 10
	fmt.Printf("10发送到通道b中了...\n")
	b <- 20
	fmt.Printf("20发送到通道b中了...\n")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}

func bufChannel1() {
	b = make(chan int, 10)
	b <- 10
	b <- 20
}

func bufChannel2() {
	for {
		select {
		case j := <-b:
			fmt.Println(j)
		default:
		}
	}
}

func main() {

	bufChannel1()
	bufChannel2()

}

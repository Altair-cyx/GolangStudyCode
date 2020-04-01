package main

import (
	"fmt"

	calc "github.com/Altair-cyx/day05/08calc"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行...")
	fmt.Println(x, pi)
}

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}

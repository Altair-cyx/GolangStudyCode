package main

import "fmt"

func main() {
	// 回文判断
	// 字符串从左往右读和从右往左读是一样的，那么就是回文
	ss := "山西运煤车煤运西山"
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println(r)
	for i := 0; i < len(r); i++ {

		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}

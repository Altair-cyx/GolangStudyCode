package main

import "fmt"

// array数组练习题
// 1.求数组{1, 3, 5, 7, 8}所有元素的和
// 找出数组中和为指定值的两个元素的下标，比如从数组{1, 3, 5, 7, 8}中找出和为8的两个元素下标分别为{0, 3}和{1, 2}

func main() {

	// 1.求数组{1, 3, 5, 7, 8}所有元素的和
	a1 := [...]int{1, 3, 5, 7, 8}
	s1 := 0
	// for i := 0; i < len(a1); i++ {
	// 	s1 += a1[i]
	// }
	for _, v := range a1 {
		s1 += v
	}
	fmt.Println(s1)

	// 找出和为8的两个元素下标分别为{0, 3}和{1, 2}
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("{%d, %d}\n", i, j)
			}
		}
	}
}

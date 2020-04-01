package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析出对应的数据
	intStr := "10000"
	intValue, _ := strconv.ParseInt(intStr, 10, 64)
	fmt.Printf("%#v %T\n", intValue, intValue)

	// Atoi从字符串中解析出对应的数据
	retInt, _ := strconv.Atoi(intStr)
	fmt.Printf("%T\n", retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串中解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

	// 把数字转换成字符串类型
	i := 97
	// ret2 := string(i) // "a"
	ret2 := fmt.Sprintf("%d", i) // "97"
	fmt.Println(ret2)

	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v\n", ret3)
}

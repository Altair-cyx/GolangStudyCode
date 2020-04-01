package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
// 111

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	// write
	fileObj.Write([]byte("周林蒙蔽了！\n"))
	// writeString
	fileObj.WriteString("周林解释不了")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello沙河\n") // 写到缓存中
	wr.Flush()                  // 将缓存中的内容推送到文件
}

func writeDemo3() {
	str := "hello 沙河1"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}

func main() {
	writeDemo3()
}

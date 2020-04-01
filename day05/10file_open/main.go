package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件
func readFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// tmp := make([]byte,128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

func readFromFileBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}
	fmt.Print(string(ret))
}

func main() {
	// readFromFileBufio()
	readFromFileIoutil()
}

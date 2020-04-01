package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	strByte, err := ioutil.ReadFile("./index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(strByte)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在URL上（query param），请求体里面是没有数据的
	queryParam := r.URL.Query()
	fmt.Println(queryParam.Get("name"))
	fmt.Println(queryParam.Get("age"))
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) // 我在服务端打印客户端发来的请求
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/ports/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

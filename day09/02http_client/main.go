package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

// 共用一个client适用于请求比较频繁
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:8080/xxx/")
	// if err != nil {
	// 	fmt.Println("get url failed,err:", err)
	// 	return
	// }
	urlObj, _ := url.Parse("http://127.0.0.1:8080/xxx/")
	data := url.Values{} // url values
	data.Set("name", "周林")
	data.Set("age", "90000")
	urlStr := data.Encode() // URL encode之后的URL
	fmt.Println(urlStr)
	urlObj.RawQuery = urlStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("get url failed,err:", err)
	// 	return
	// }
	// 请求不是特别频繁，用完就关闭该链接
	// 禁用KeepAlive的client
	// tr := &http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: tr,
	// }
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get url failed,err:", err)
		return
	}
	defer resp.Body.Close() // 一定要记得关闭resp.Body
	// 发请求
	// 从resp中把服务端返回的数据读出来
	str, err := ioutil.ReadAll(resp.Body) // 我在客户端读出服务端返回的内容
	if err != nil {
		fmt.Println("read body failed,err:", err)
		return
	}
	fmt.Println(string(str))
}

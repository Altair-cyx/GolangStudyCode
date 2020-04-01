package split

import (
	"fmt"
	"strings"
)

// 切割字符串
// example:
// abc, b -> [a c]

// Split 根据切割字符串
func Split(str string, sep string) []string {
	// str:"a:b:c" sep=":"
	ret := make([]string, 0, strings.Count(str, sep))
	index := strings.Index(str, sep) // 1
	for index >= 0 {
		ret = append(ret, str[:index])  // [a] => [a b]
		str = str[index+len(sep):]      // str = "b:c" => str = "c"
		index = strings.Index(str, sep) // index=1 => index=-1
	}
	if index == -5 {
		fmt.Println("真无聊")
	}
	ret = append(ret, str) // ret = [a b c]
	return ret             // [a b c]
}

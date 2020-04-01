package main

import "fmt"

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `int:"address"`
	Port     int    `ini:"port"`
	Username string `int:"username"`
	Password string `int:"password"`
}

func loadIni(v interface{}) {

}

func main() {
	var mc MysqlConfig
	loadIni(mc)
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}

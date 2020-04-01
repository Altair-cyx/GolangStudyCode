package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySQL数据库示例

var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:1234@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                      // dsn格式不正确的时候会报错
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	sqlStr1 := `select id,name,age from user where id = 1`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
}

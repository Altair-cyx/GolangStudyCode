package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL
// 数据库
// 常见的数据库SQLite、MySQL、SQLServer、PostgreSQL、Oracle
// MySQL主流的关系型数据库，类似的还有postgreSQL
// 关系型数据库：
// 用表来存一类数据

// MySQL知识点
// SQL语句
// DDL：操作数据库的
// DML：表的增删改查
// DCL：用户及权限

// 存储引擎
// MySQL支持插件式的存储引擎
// 常见的存储引擎：MyISAM和innoDB
// MyISAM:
// 	1.查询速度快
// 	2.只支持表锁
// 	3.不支持事务
// InnoDB：
// 	1.整体速度快
// 	2.支持表锁和行锁
// 	3.支持事务

// 事务：
// 把多个SQL操作当成一个整体。

// 事务的特点：
// ACID
// 原子性：事务要么成功要么失败
// 一致性：数据库的完整性没有被破坏
// 隔离性：事务之间是互相隔离的。
// 		1.隔离的四个级别
// 持久性：事务操作的结果是不会丢失的

// 索引：
// 索引的原理是：B树和B+树
// 索引的类型
// 索引的命中
// 分库分表
// SQL注入
// SQL慢查询优化

// MySQL主从
// 		binlog
// MySQL读写分离

// Go操作MySQL
// database/sql
// 原生支持连接池，是并发安全的。
// 这个标准库没有具体的实现，只是列出了一些需要第三方库实现的具体内容。

// 下载驱动
// go get -u github.com/go-sql-driver/mysql
//
func main() {
	// 数据库信息
	dsn := "root:1234@tcp(127.0.0.1:3306)/studentsystem"
	// 连接数据库
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                   // dsn格式不正确的时候会报错
		fmt.Printf("dsn:%s invalid,err:%v", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v", dsn, err)
		return
	}
	fmt.Println("连接成功")
	db.Close()
}

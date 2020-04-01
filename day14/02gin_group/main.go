package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context){
	name := c.DefaultQuery("name","jack")
	c.String(200,fmt.Sprintf("hello %s\n",name))
}

func submit(c *gin.Context){
	name := c.DefaultQuery("name","lily")
	c.String(200,fmt.Sprintf("hello %s\n",name))
}

func main(){
	// 创建路由
	r := gin.Default()

	// 路由组1：处理GET请求
	v1 := r.Group("/v1")
	// {}书写规范
	{
		v1.GET("/login",login)
		v1.GET("/submit",submit)
	}
	// 路由组2：处理POST请求
	v2 := r.Group("/v2")
	{
		v2.POST("/login",login)
		v2.POST("/submit",submit)
	}

	// 启动8000端口，监听
	r.Run(":8000")
}

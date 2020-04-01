package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// 多种响应方式

func main(){
	r := gin.Default()
	// JSON
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"someJSON","status":200})
	})
	// 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct{
			Name string
			Message string
			Number int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200,msg)
	})

	// XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200,gin.H{"message":"abc"})
	})

	// YAML
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200,gin.H{"name":"chen"})
	})

	// protobuf格式 谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	r.GET("/someProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1),int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200,data)
	})

	r.Run(":8080")
}

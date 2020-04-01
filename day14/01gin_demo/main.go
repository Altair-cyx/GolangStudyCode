package _1gin_demo

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

// gin的helloworld

func jsonEncode(content interface{}) (str string) {
	b, err := json.Marshal(content)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}
	str = string(b)
	return
}

func htmlEncode(filePath string,c *gin.Context){
	b,err := ioutil.ReadFile(filePath)
	if err != nil{
		log.Println("read file failed,err:",err)
		return
	}
	c.Writer.Write(b)
	c.Writer.Flush()
}

func main() {

	// 1.创建路由
	// 默认使用了2个中间件Logger()，Recovery()
	r := gin.Default()
	// 也可以创建不带两个中间件的
	// r := gin.New()

	r.GET("/login", func(c *gin.Context) {
		htmlEncode("./page/login.html",c)
	})

	r.GET("/upload", func(c *gin.Context) {
		htmlEncode("./page/upload.html",c)
	})

	r.GET("/upload2", func(c *gin.Context) {
		htmlEncode("./page/upload2.html",c)
	})

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	//r.GET("/xxx",func(c *gin.Context){
	//	c.String(http.StatusOK,str)
	//})
	//r.POST("/xxx")
	//r.PUT("/xxx")

	// api参数 :获取单个参数，*获取多个参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		//
		name := c.Param("name")
		action := c.Param("action")

		c.String(http.StatusOK,fmt.Sprintf("name is %s,action is %s",name,action))
	})

	// URL参数
	r.GET("/sload", func(c *gin.Context) {
		c.Query("name")
		// 默认参数
		c.DefaultQuery("name","chen")
	})

	// 表单参数
	r.POST("/form", func(c *gin.Context) {
		uName := c.PostForm("username")
		pwd := c.PostForm("password")
		// 默认提交参数
		//c.DefaultPostForm()
		c.String(http.StatusOK, fmt.Sprintf("username=%s \t password=%s", uName, pwd))
	})

	// 单个文件上传
	r.POST("/upload", func(c *gin.Context) {
		// 表单传文件
		file,_ := c.FormFile("file")
		log.Println(file.Filename)
		// 传到项目根目录，名字就用本身的
		_ = c.SaveUploadedFile(file,file.Filename)

		// 打印信息
		c.String(200,fmt.Sprintf("%s upload!",file.Filename))
	})
	// 限制表单上传大小 8MB 默认32MB
	r.MaxMultipartMemory = 8 << 20
	// 多个文件上传
	r.POST("/upload2", func(c *gin.Context) {
		form,err := c.MultipartForm()
		if err != nil{
			c.String(http.StatusBadRequest,fmt.Sprintf("get error %s",err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _,file := range files{
			// 逐个存
			if err := c.SaveUploadedFile(file,file.Filename); err != nil{
				c.String(http.StatusBadRequest,fmt.Sprintf("upload error %s",err.Error()))
				return
			}
		}
		c.String(200,"upload ok %d files",len(files))
	})

	// 3.监听端口，默认在8080
	err := r.Run(":8000")
	if err != nil{
		return
	}
}

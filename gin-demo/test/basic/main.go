package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Login struct {
	User     string `form:"user" json:"user" uri:"user" xml:"user" bind:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" bind:"required"`
}

func main() {
	//	1. 创建默认路由
	r := gin.Default()

	//通过Context的Param方法获取URL中的PATH参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		//params := c.Params
		//for _, param := range params {
		//	fmt.Println(param)
		//}

		name := c.Param("name")
		action := c.Param("action")
		//截取action
		action = strings.Trim(action, "/")

		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
			"status": 200,
		})
	})

	//获取URL中的queryString
	r.GET("/query", func(c *gin.Context) {
		//第一种，直接使用Query
		//name := c.Query("name")
		//age := c.Query("age")
		//第二种，使用DefaultQuery
		//name := c.DefaultQuery("name", "somebody")
		//age := c.DefaultQuery("age", "18")
		//第三种，使用GetQuery
		name, ok := c.GetQuery("name")
		if !ok {
			name = "somebody"
		}
		age, ok := c.GetQuery("age")
		if !ok {
			age = "28"
		}

		c.String(http.StatusOK, "%s", name+age)
	})

	//获取表单参数，传输格式限定为 application/x-www-form-urlencoded 或 application/form-data
	r.POST("/form", func(c *gin.Context) {
		//第一种方式，直接PostForm
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		//第二种方式，不存在则输出默认值的方式, DefaultPostForm
		username := c.DefaultPostForm("username", "lxx")
		password := c.DefaultPostForm("password", "222")

		c.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s", username, password))
	})

	//参数绑定，自动识别content-type请求数据类型，并利用反射机制自动提取请求中query/form/json/xml等参数到结构体中
	//绑定json的示例(application/json)
	r.POST("/loginJson", func(c *gin.Context) {
		//声明接收数据的变量
		var login Login
		//将request的body中的数据，自动按照json格式解析到结构体中;
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})
	//绑定form表单的示例(application/x-www-urlencoded 或 application/form-data)
	r.POST("/loginForm", func(c *gin.Context) {
		//声明指定结构体类型的变量
		var login Login
		//ShouldBind会根据content-type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	//绑定URL中的QueryString示例
	r.POST("/loginQuery", func(c *gin.Context) {
		//声明变量接收数据
		var login Login
		if err := c.ShouldBindQuery(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	//绑定URI中的PATH
	r.GET("/loginUser/:user/:password", func(c *gin.Context) {
		//声明接收的变量
		var login Login

		if err := c.ShouldBindUri(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	//默认监听8080端口，可指定
	_ = r.Run(":9090")
}

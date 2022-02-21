package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制上传文件大小
	r.MaxMultipartMemory = 8 << 20 //20MB
	//配置获取请求的路由
	r.POST("/upload", func(c *gin.Context) {
		//获取多个文件
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusInternalServerError, "上传文件出错, err=%s", err.Error())
			return
		}
		//获取key为f1对应的所有文件
		files := form.File["f1"] // []*file.Header, f1表示key值
		//fmt.Println("files:", files)
		//遍历文件切片，依次保存每个文件
		for index, file := range files {
			fmt.Println(file.Filename)
			//逐个存文件
			dst := fmt.Sprintf("file/%s_%d", file.Filename, index)
			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Code":    http.StatusInternalServerError,
					"Message": fmt.Sprintf("文件保存失败，err: %s", err.Error()),
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": fmt.Sprintf("%d 个文件上传成功！", len(files)),
		})
	})

	_ = r.Run(":9090")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制上传文件的最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "上传图片出错, err=%s", err.Error())
			return
		}
		dst := fmt.Sprintf("file/%s", file.Filename)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.String(http.StatusInternalServerError, "图片保存失败, err=%s", err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded '%s'!", file.Filename, dst),
		})
	})

	_ = r.Run(":9090")
}

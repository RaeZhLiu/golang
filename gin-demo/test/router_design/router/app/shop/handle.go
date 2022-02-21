package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleShopping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "shopping test",
	})
}

package utils

import (
	"github.com/gin-gonic/gin"
)

func ResOkFormater(c *gin.Context, m string, v interface{}, s int) {

	c.JSON(s, gin.H{
		"success": true,
		"code":    s,
		"message": m,
		"data":    v,
	})
}

func ResFailFormater(c *gin.Context, m string, err interface{}, s int) {
	c.JSON(s, gin.H{
		"success": false,
		"code":    400,
		"message": m,
		"errors":  err,
	})
}

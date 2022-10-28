package authorization

import "github.com/gin-gonic/gin"

func Autho(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": "asdasd",
	})
}

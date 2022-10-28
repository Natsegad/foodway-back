package authenticaton

import "github.com/gin-gonic/gin"

func Authe(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": "ok",
	})
}

package register

import (
	"github.com/gin-gonic/gin"
	"io"
)

func Register(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"token": string(data),
	})
}

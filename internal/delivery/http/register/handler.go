package register

import (
	"encoding/json"
	"foodway/internal/domain"
	"foodway/internal/service"
	"foodway/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Register(c *gin.Context) {
	log := logger.GetLogger()

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("Error ReadAll json: %v ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json error 1",
		})
		return
	}

	var user domain.UserInfo

	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Errorf("Error Unmarshal json: %v ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json error 2",
		})
		return
	}

	service.Registration(user)

	c.JSON(200, gin.H{
		"token": 1,
	})
}

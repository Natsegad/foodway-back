package authorization

import (
	"encoding/json"
	"foodway/internal/domain"
	"foodway/internal/service"
	"foodway/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Autho(c *gin.Context) {
	log := logger.GetLogger()

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("Error ReadAll json: %v ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json error 1",
		})
		return
	}

	var user domain.JsonUserInfo

	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Errorf("Error Unmarshal json: %v ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json error 2",
		})
		return
	}

	resp, err := service.Authorization(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, resp)
}

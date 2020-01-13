package controller

import (
	"github.com/gin-gonic/gin"
)

type KafkaController struct{}

func GetKafkaController() *KafkaController {
	return &KafkaController{}
}

func (controller *KafkaController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

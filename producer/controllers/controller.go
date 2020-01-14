package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type KafkaController struct{}

func GetKafkaController() *KafkaController {
	return &KafkaController{}
}

func (controller *KafkaController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}

func (controller *KafkaController) Post(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(400, gin.H{
			"message": "Bad Response",
		})
		return
	}
	text := fmt.Sprintf("Success: %s was posted\n", id)
	c.JSON(200, gin.H{
		"message": text,
	})
}

package controller

import (
	"../../config"
	senders "../../producer/senders"
	"github.com/gin-gonic/gin"
)

type KafkaController struct {
	Config config.Config
}

func GetKafkaController(c config.Config) *KafkaController {
	return &KafkaController{
		Config: c,
	}
}

func (controller *KafkaController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}

func (controller *KafkaController) Post(c *gin.Context) {
	text := c.PostForm("text")
	topic := c.PostForm("topic")
	if text == "" {
		c.JSON(400, gin.H{
			"message": "Bad Response",
		})
		return
	}
	s := senders.GetKafkaSender(controller.Config)
	s.Send(text, topic)
	c.JSON(202, gin.H{
		"message": topic + ": " + text,
	})
}

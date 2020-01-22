package infrastructure

import (
	"../../config"
	controllers "../../producer/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	conf := config.GetConfig()
	router := gin.Default()
	ctrl := controllers.GetKafkaController(*conf)
	router.GET("/", func(c *gin.Context) { ctrl.Index(c) })
	router.POST("/send", func(c *gin.Context) { ctrl.Post(c) })

	Router = router
}

package infrastructure

import (
	controllers "../../producer/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	ctrl := controllers.GetKafkaController()
	router.GET("/", func(c *gin.Context) { ctrl.Index(c) })
	router.POST("/send", func(c *gin.Context) { ctrl.Post(c) })

	Router = router
}

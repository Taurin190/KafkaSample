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

	Router = router
}

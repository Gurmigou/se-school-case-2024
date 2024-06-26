package main

import (
	"github.com/gin-gonic/gin"
	"se-school-case/controller"
	"se-school-case/initializer"
	"se-school-case/service"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
	initializer.RunMigrations()
	service.StartScheduledEmail()
}

func main() {
	r := gin.Default()

	/*
		Util end-points:
		/api/ping - ping-pong server
		/api/notify - explicitly notifies all users using email without schedules interval
	*/
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/api/notify", controller.PostExplicitlyNotify)

	// Required end-points
	r.POST("/api/subscribe", controller.PostAddUserEmail)
	r.GET("/api/rate", controller.GetExchangeRate)

	initializer.StartServer(r)
}

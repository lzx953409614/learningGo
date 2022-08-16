package main

import (
	"user/api"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	user := router.Group("/api/v1/user/")
	{
		user.POST("/add", api.AddPerson)
		user.POST("/get", api.GetPerson)
		user.GET("/getAll", api.GetPersonAllPerson)
		user.POST("/update", api.UpdatePerson)
		user.POST("/delete", api.DeletePerson)
	}
	redis := router.Group("api/v1/redis")
	{
		redis.GET("/setRedis", api.SetRedis)
		redis.GET("/getRedis", api.GetRedis)
	}
	rabbitmq := router.Group("api/v1/rabbitmq")
	{
		rabbitmq.GET("/sendMsg", api.SendMsg)
		rabbitmq.GET("/subMsg", api.SubMsg)
	}
	mqtt := router.Group("/api/v1/mqtt")
	{
		mqtt.GET("/publishMqttMsg", api.PublishMqttMsg)
		mqtt.GET("/subcribeMqttMsg", api.SubcribeMqttMsg)
	}
	return router
}

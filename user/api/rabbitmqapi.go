package api

import (
	"net/http"
	"user/rabbitmq"

	"github.com/gin-gonic/gin"
)

func SendMsg(c *gin.Context) {
	msg := c.Query("msg")
	rabbitmq.PublishMsg(msg)
	c.JSON(http.StatusOK, gin.H{
		"msg": "send msg success",
	})
}

func SubMsg(c *gin.Context) {
	//rabbitmq.SubcribeRabbitMqMsg()
	c.JSON(http.StatusOK, gin.H{
		"msg": "sub msg success",
	})
}

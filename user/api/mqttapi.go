package api

import (
	"net/http"
	"user/mqtt"

	"github.com/gin-gonic/gin"
)

func PublishMqttMsg(c *gin.Context) {
	msg := c.Query("msg")
	topic := c.Query("topic")
	mqtt.SendMqttMsg(topic, msg)
	c.JSON(http.StatusOK, gin.H{
		"msg": "publish mqtt msg success",
	})
}

func SubcribeMqttMsg(c *gin.Context) {
	topic := c.Query("topic")
	mqtt.SubMqttMsg(topic)
	c.JSON(http.StatusOK, gin.H{
		"msg": "subcribe mqtt msg success",
	})
}

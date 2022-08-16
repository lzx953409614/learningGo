package api

import (
	"net/http"
	"user/db"

	"github.com/gin-gonic/gin"
)

func SetRedis(c *gin.Context) {
	db.RedisClient.Set("test1", "测试单机版redis设置key", -1)
	//db.RedisClusterClient.Set("test2", "测试集群版redis设置key", -1)
}

func GetRedis(c *gin.Context) {
	rs1, err1 := db.RedisClient.Get("test1").Result()
	if err1 != nil {
		c.JSON(http.StatusOK, "system error!")
	}

	//rs2, err2 := db.RedisClusterClient.Get("test2").Result()
	//if err2 != nil {
	//	c.JSON(http.StatusOK, "system error!")
	//}

	c.JSON(http.StatusOK, gin.H{
		"key1":   "test1",
		"value1": rs1,
		//"key2":   "test2",
		//"value2": rs2,
	})
}

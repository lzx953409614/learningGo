package db

import (
	"log"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

//var RedisClusterClient *redis.ClusterClient

func init() {
	var err error
	//单机版(如果能使用阿里云redis的话，业务量很大也可以用单机方式)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	//集群版
	// RedisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003",
	// 		"127.0.0.1:7004", "127.0.0.1:7005", "127.0.0.1:7006"},
	// 	Password:     "",
	// 	PoolSize:     500,
	// 	MinIdleConns: 50,
	// 	ReadTimeout:  3000 * time.Millisecond,
	// })
	// RedisClusterClient.Do("set", "name", "local")
	// _, err = RedisClusterClient.Ping().Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

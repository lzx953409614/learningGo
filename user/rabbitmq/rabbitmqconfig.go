package rabbitmq

import (
	"github.com/streadway/amqp"

	"log"
)

func InitRabbitMqConn() *amqp.Connection {
	//连接rabbitmq 默认可以使用Dial,指定虚拟机和指定最大通道数需要用DialCOnfig来定义
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	RabbitMqConn, err := amqp.DialConfig("amqp://guest:guest@localhost:5672/", amqp.Config{
		Vhost:      "/test", //指定虚拟机
		ChannelMax: 999999,  //最大通道数据
	})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("初始化rabbitmq成功! %v", RabbitMqConn)
	return RabbitMqConn
}

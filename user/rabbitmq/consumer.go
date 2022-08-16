package rabbitmq

import (
	"fmt"
	"log"
)

func init() {
	/**
	*新启动一个go协程去启动rabbitmq消费者监听
	*直接init()方法启动消费者会导致gin router初始化失败
	 */
	go SubcribeRabbitMqMsg("consumer1")
	go SubcribeRabbitMqMsg("consumer2")
}

func SubcribeRabbitMqMsg(consumrename string) {
	//创建连接
	conn := InitRabbitMqConn()
	//创建通道，其实是创建tcp连接
	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	defer ch.Close()

	//创建队列
	q, err := ch.QueueDeclare(
		"test_queue", // name 声明队列的名称
		false,        // durable，是否持久化
		false,        // delete when unused 声明一个自动删除队列（服务器将在不再使用它时将其删除）
		false,        // exclusive 声明一个独占队列
		false,        // no-wait
		nil,          // arguments 队列的其他属性（构造参数）
	)
	if err != nil {
		log.Panic(err)
	}

	msgch, err := ch.Consume(
		q.Name,       // queue
		consumrename, // consumer
		false,        //autoAck 是否自动确认，一旦消息队列将消息发送给消息消费者后，就会从内存中将这个消息删除
		false,        //exclusive 是否排外（排外：queue只被一个消费者使用并且在消费者断开连接时queue被删除），一般默认为false
		false,        //noLocal
		false,        // noWait 当nowait为true时，不要等待服务器确认请求就立即开始消费消息。如果不能消费，有可能引发通道异常并关闭通道。一般默认设置为false。
		map[string]interface{}{},
	)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("rabbitmq" + consumrename + "订阅发起成功！")
	for msg := range msgch {
		fmt.Printf(consumrename+"接收到消息：%s", msg.Body)
		//休眠时长 单位纳秒
		//休眠3s 3e9表示3x1000000000纳秒
		//time.Sleep(1e9)
		msg.Ack(true) //消息应答
	}
}

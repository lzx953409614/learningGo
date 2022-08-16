package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitmqCon *amqp.Connection

func init() {
	RabbitmqCon = InitRabbitMqConn()
}

func PublishMsg(msg string) {
	//创建连接
	//conn := InitRabbitMqConn()
	//创建通道，其实是创建tcp连接
	ch, err := RabbitmqCon.Channel()
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
	//推送数据
	err = ch.Publish(
		"",     // exchange ， 交换机
		q.Name, // routing key，声明队列的名称，指定发送哪个队列中
		false,  // mandatory，当mandatory设置为true时，如果exchange根据自身类型和消息routeKey无法找到一个符合条件的queue，那么会调用basic.return方法将消息返还给生产者；当mandatory设为false时，出现上述情形broker会直接将消息扔掉
		false,  // immediate， 当immediate设置为true时，如果exchange在将消息route到queue时发现对应的queue上没有消费者，那么这条消息不会放入队列中。当与消息routeKey关联的所有queue(一个或多个)都没有消费者时，该消息会通过basic.return方法返还给生产者
		amqp.Publishing{
			ContentType: "text/plain", //指定消息格式
			Body:        []byte(msg),  // 消息体
		})
	if err != nil {
		log.Panic(err)
	}
	log.Printf(" [x] Sent %s\n", msg)
}

package mqtt

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const broker = "tcp://ip:port"
const username = "admin"
const password = "public"

func InitMqttClient(ClientID string) mqtt.Client {
	//配置
	clinetOptions := mqtt.NewClientOptions().AddBroker(broker).SetUsername(username).SetPassword(password)
	clinetOptions.SetClientID(ClientID)
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	clinetOptions.SetKeepAlive(60)
	//连接
	MqttClient := mqtt.NewClient(clinetOptions)
	//客户端连接判断
	if token := MqttClient.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Printf("mqtt client [%s] init success", ClientID)
	return MqttClient
}

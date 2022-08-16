package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var SubClient mqtt.Client

//message的回调
var onMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[%s] -> %s\n", msg.Topic(), msg.Payload())
}

func init() {
	SubClient = InitMqttClient("go_mqtt_sub_test")
}

func SubMqttMsg(topic string) {
	qos := 0
	token := SubClient.Subscribe(topic, byte(qos), onMessage)
	token.Wait()
}

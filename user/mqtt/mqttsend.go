package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var PubClient mqtt.Client

func init() {
	PubClient = InitMqttClient("go_mqtt_pub_test")
}

func SendMqttMsg(topic string, msg string) {
	qos := 0
	payload := fmt.Sprintf("message: %s", msg)
	if token := PubClient.Publish(topic, byte(qos), false, payload); token.Wait() && token.Error() != nil {
		fmt.Printf("publish failed, topic: %s, payload: %s\n", topic, payload)
	} else {
		fmt.Printf("publish success, topic: %s, payload: %s\n", topic, payload)
	}
}

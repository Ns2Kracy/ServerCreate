package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//const topic = "$thing/down/property/PZA7QTDRXX/test001"
const topic = "t/a"

type payload struct {
	Gy              float64 `json:"gy"`
	Dht             float64 `json:"dht"`
	HeartRate       float64 `json:"heart_rate"`
	Temperature     float64 `json:"temperature"`
	AlcoholStrength float64 `json:"alcohol_strength"`
}

func main() {
	client := MqttClient()
	go Subscribe(client)        // 在主函数里, 我们用另起一个 go 协程来订阅消息
	time.Sleep(time.Second * 1) // 暂停一秒等待 subscribe 完成
	Publish(client)
}

func MqttClient() mqtt.Client {
	var username = ""
	var password = ""

	connectAddress := "ws://139.224.19.236:8083/mqtt"
	client_id := fmt.Sprintf("go-client")

	fmt.Println("connect address: ", connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)

	//设定username
	opts.SetUsername(username)
	//设定password
	opts.SetPassword(password)
	opts.SetClientID(client_id)
	opts.SetKeepAlive(60)

	client := mqtt.NewClient(opts)

	token := client.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}

func Publish(client mqtt.Client) {
	qos := 0
	for {
		payloadTest := payload{
			Gy:              1,
			Dht:             1,
			HeartRate:       1,
			Temperature:     1,
			AlcoholStrength: 1,
		}
		test, _ := json.Marshal(payloadTest)
		if token := client.Publish(topic, byte(qos), false, test); token.Wait() && token.Error() != nil {
			fmt.Printf("publish failed, topic: %s, payload: %s\n", topic, test)
		} else {
			fmt.Printf("publish success, topic: %s, payload: %s\n", topic, test)
		}
		time.Sleep(time.Second * 5)
	}
}

func Subscribe(client mqtt.Client) {
	var qos int
	fmt.Sprintf("Set qos: %d", qos)
	client.Subscribe(topic, byte(qos), func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received `%s` from `%s` topic\n", msg.Payload(), msg.Topic())
	})
}

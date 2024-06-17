package mq

import (
	"github.com/streadway/amqp"
	"log"
	"tiktok/config"
)

var RabbitmqConn *amqp.Connection

func LinkRabbitmq() {
	url := "amqp://" + config.RabbitmqUserName + ":" + config.RabbitmqPassword + "@" + config.RabbitmqIP + "/"
	conn, err := amqp.Dial(url) //("amqp://guest:guest@192.168.239.126:5672/")
	if err != nil {
		panic(err)
	}
	RabbitmqConn = conn
	log.Println("connecting to rabbitmq successes")
}

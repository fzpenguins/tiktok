package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func PublishMsg(data []byte, to string) (err error) {
	//rabbitmq采用direct模式
	//声明通道
	ch, err := RabbitmqConn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	//声明交换机
	err = ch.ExchangeDeclare(
		"direct_exchange", // Exchange 名称
		"direct",          // Exchange 类型
		true,              // 持久化
		false,             // 不自动删除
		false,             // 不等待服务器响应
		false,
		nil, // 不设置额外参数
	)
	if err != nil {
		log.Println(err)
		return
	}
	// 声明队列
	queue, err := ch.QueueDeclare(
		to,    // 队列名称
		false, // 不持久化
		false, // 不自动删除
		false, // 不独占
		false, // 不等待服务器响应
		nil,   // 不设置额外参数
	)
	if err != nil {
		log.Println(err)
		return
	}

	// 交换机和队列绑定
	err = ch.QueueBind(
		queue.Name,        // 队列名称
		to,                // 路由键，用于绑定 Exchange 和队列
		"direct_exchange", // Exchange 名称
		false,             // 不等待服务器响应
		nil,               // 不设置额外参数
	)
	if err != nil {
		log.Println(err)
		return
	}
	// 发布消息
	err = ch.Publish(
		"direct_exchange", // Exchange 名称
		to,                // 路由键
		false,             // 不等待服务器响应
		false,             // 不设置额外参数
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		log.Println(err)
		return
	}
	return nil
}

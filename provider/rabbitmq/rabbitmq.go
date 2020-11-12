package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channel *amqp.Channel
var queuse map[string]amqp.Queue

func init() {
	fmt.Println(".:::AMQP:::.")
	connection = connect()
	channel = CreateChannel()
	queuse = make(map[string]amqp.Queue)
}

func connect() *amqp.Connection {
	url := "amqp://hjguplwm:I-wumDDxm3KgPbyATuFGplnJ5POkO_Hp@lion.rmq.cloudamqp.com/hjguplwm"
	connection, err := amqp.Dial(url)
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	return connection
}

func CreateChannel() *amqp.Channel {
	channel, err := connection.Channel()
	if err != nil {
		panic("could not create channel:" + err.Error())
	}
	return channel
}

func QueueDeclare(queueName string, durable bool) amqp.Queue {
	_, ok := queuse[queueName]
	if ok {
		return queuse[queueName]
	}
	q, err := channel.QueueDeclare(
		queueName,
		durable,
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic("Failed to declare a queue:" + err.Error())
	}
	queuse[queueName] = q
	return q
}

func Publish(queueName string, body string) {
	queue := QueueDeclare(queueName, false)
	err := channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		panic("Failed to publish a messagel:" + err.Error())
	}
}

func Consume(queueName string) <-chan amqp.Delivery {
	msgs, err := channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}
	return msgs
}

package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channel *amqp.Channel
var exchanges map[string]bool
var queuse map[string]amqp.Queue

func init() {
	fmt.Println(".:::AMQP:::.")
	connection = connect()
	channel = CreateChannel()
	exchanges = make(map[string]bool)
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

func ExchangeDeclare(exchangeName string, exchangeType string, durable bool) {
	_, ok := exchanges[exchangeName]
	if !ok {
		err := channel.ExchangeDeclare(
			exchangeName,
			exchangeType,
			durable, // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			panic("Failed to declare a queue:" + err.Error())
		}
		exchanges[exchangeName] = true
	}
}

func QueueDeclare(queueName string, durable bool) {
	_, ok := queuse[queueName]
	if !ok {
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
	}
}

func CreatePublisher(exchangeName string, exchangeType string, durable bool, queueName string) {
	ExchangeDeclare(exchangeName, exchangeType, durable)
	QueueDeclare(queueName, false)

}

func Publish(exchangeName string, queueName string, body string) {
	err := channel.Publish(
		exchangeName, // exchange
		queueName,    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		panic("Failed to publish a messagel:" + err.Error())
	}
}

func QueueBind(exchangeName string, queueName string) {
	err := channel.QueueBind(
		queueName,
		queueName, // routing key
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		panic("Failed to publish a messagel:" + err.Error())
	}

}

func Consume(exchangeName string, queueName string) (<-chan amqp.Delivery, error) {
	QueueBind(exchangeName, queueName)
	return channel.Consume(queueName, "", false, false, false, false, nil)
}

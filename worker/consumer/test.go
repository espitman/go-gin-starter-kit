package consumer

import (
	"fmt"
	"jettster/provider/rabbitmq"
)

func TestConsumer() {
	msgs, _ := rabbitmq.Consume("ginTestExchange", "ginTestQueue")
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			body := string(msg.Body)
			fmt.Println("message received: " + string(body))
			_ = msg.Ack(true)
		}
	}()
	<-forever
}

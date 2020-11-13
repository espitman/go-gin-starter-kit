package consumer

import (
	"fmt"
	"jettster/provider/rabbitmq"
)


func (t *T) BozConsumer() {
	msgs, _ := rabbitmq.Consume("ginTestExchange2", "ginTestQueue2")
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			body := string(msg.Body)
			fmt.Println("message received2: " + string(body))
			_ = msg.Ack(true)
		}
	}()
	<-forever
}

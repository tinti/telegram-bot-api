package main

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/streadway/amqp"
)

type RequestMessage struct {
	Op string
	Value int
}

type ResponseMessage struct {
	Op string
	Value int
	ResponseValue int
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func double(n int) int {
	return n * 2
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/guest")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n := RequestMessage{}
			err := json.Unmarshal(d.Body, &n)
			failOnError(err, "Failed to convert body to RequestMessage")

			var r ResponseMessage
			if n.Op == "fib" {
				log.Printf(" [.] fib(%d)", n.Value)
				r = ResponseMessage{"fib", n.Value, fib(n.Value)}
			} else {
				log.Printf(" [.] dou(%d)", n.Value)
				r = ResponseMessage{"dou", n.Value, double(n.Value)}
			}

			response, _ := json.Marshal(r)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(response),
				})
			failOnError(err, "Failed to publish a ResponseMessage")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

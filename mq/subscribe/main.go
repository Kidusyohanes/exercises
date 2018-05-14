package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

const maxConnRetries = 5

func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

func main() {
	mqAddr := reqEnv("MQADDR")
	mqName := reqEnv("MQNAME")

	conn, err := amqp.Dial("amqp://" + mqAddr)
	if err != nil {
		log.Fatalf("error dialing MQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error getting channel: %v", err)
	}
	q, err := channel.QueueDeclare(mqName,
		false, //non-durable
		false, //won't auto-delete q after idle
		false, //non-exclusive
		false, //wait for the server to respond
		nil)   //server-specific args
	if err != nil {
		log.Fatalf("error declaring queue: %v", err)
	}

	msgs, err := channel.Consume(q.Name, //name of the queue
		"",    //name of this consumer
		false, //don't want to auto-ack
		false, //non-exclusive
		false, //noLocal (not supported by RabbitMQ)
		false, //wait for the server to respond before delivering
		nil)   //extra server-specific args
	if err != nil {
		log.Fatalf("error consuming messages: %v", err)
	}

	for msg := range msgs {
		log.Printf("received message: %s", string(msg.Body))
		msg.Ack(false)
	}
}

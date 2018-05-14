package main

import (
	"log"
	"net/http"
	"os"
	"time"

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

func connectToMQ(addr string) (*amqp.Connection, error) {
	mqURL := "amqp://" + addr
	var conn *amqp.Connection
	var err error
	for i := 1; i <= maxConnRetries; i++ {
		conn, err = amqp.Dial(mqURL)
		if err == nil {
			log.Printf("successfully connected to %s", mqURL)
			return conn, nil
		}
		log.Printf("error connecting to MQ at %s: %v", mqURL, err)
		log.Printf("will retry in %d seconds", i*2)
		time.Sleep(time.Second * time.Duration(i*2))
	}
	return nil, err
}

func processMessages(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		log.Printf("received message: %s", string(msg.Body))
		msg.Ack(false)
	}
}

func main() {
	mqAddr := reqEnv("MQADDR")
	mqName := reqEnv("MQNAME")

	conn, err := connectToMQ(mqAddr)
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

	go processMessages(msgs)

	http.ListenAndServe("127.0.0.1:4000", nil)
}

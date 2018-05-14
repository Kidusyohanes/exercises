package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

const usage = `
usage:
	publish <your-message>
`

func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	mqAddr := reqEnv("MQADDR")
	mqName := reqEnv("MQNAME")
	message := os.Args[1]

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
	channel.Confirm(false)
	channel.Publish("", //default exchange
		q.Name,
		true,  //must be added to a queue
		false, //subscriber doesn't need to read right away
		amqp.Publishing{
			Body: []byte(message),
		})

	log.Printf("sent message: %s", message)
	channel.Close()
	conn.Close()
}

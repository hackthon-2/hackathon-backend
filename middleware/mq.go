package middleware

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var conn *amqp.Connection

func ConnectMQ() {
	user := os.Getenv("MQ_USER")
	passwd := os.Getenv("MQ_PASSWD")
	var err error
	conn, err = amqp.Dial("amqp://" + user + ":" + passwd + "@api.onesnowwarrior.cn:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
}

func MQ() *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func QueueInit(userId uint, ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"USER_"+strconv.Itoa(int(userId)),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	return q
}

func Publisher(watcherId uint, message string, ch *amqp.Channel) error {
	err := ch.Publish(
		"",
		"USER_"+strconv.Itoa(int(watcherId)),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		})
	return err
}

func Consumer(userId uint, ch *amqp.Channel) ([]string, error) {
	msg, err := ch.Consume(
		"USER_"+strconv.Itoa(int(userId)),
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	over := make(chan bool)
	data := make([]string, len(msg))
	go func() {
		var sum = 0
		for d := range msg {
			data[sum] = string(d.Body)
			sum++
		}
		over <- true
	}()
	<-over
	return data, nil
}

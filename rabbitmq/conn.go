package rabbitmq

import (
	"fmt"

	"github.com/coderc/onlinedisk-util/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	// once sync.Once
	// conn *amqp.Connection
	cfg *config.RabbitMQConfigStruct
)

func Init(config *config.RabbitMQConfigStruct) {
	cfg = config
	// var err error
	// tcpAddr := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port)
	// conn, err = amqp.Dial(tcpAddr)
	// if err != nil {
	// 	panic(err)
	// }
}

func Send(qName string, msg []byte) error {
	tcpAddr := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(tcpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)

	return err
}

func GetMsg(qName string, chanMsg chan []byte) {
	tcpAddr := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		panic(err)
	}

	for {
		select {
		case d := <-msgs:
			if len(d.Body) == 0 {
				continue
			}
			fmt.Println(fmt.Sprintf("get msg from %s: [%s]", qName, string(d.Body)))
			chanMsg <- d.Body
		}
	}
}

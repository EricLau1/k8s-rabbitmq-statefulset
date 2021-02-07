package rmq

import (
	"app/types"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Consumer interface {
	WaitMessages()
	Close() error
}

type consumer struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func NewConsumer(conn *amqp.Connection, queue string) (Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		queue,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	return &consumer{ch, q}, nil
}

func (c *consumer) getMessages() <-chan amqp.Delivery {
	msgs, err := c.ch.Consume(
		c.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		log.Fatalln("Erro ao consumir mensagem:", err.Error())
	}
	return msgs
}

func (c *consumer) WaitMessages() {
	msgs := c.getMessages()
	for m := range msgs {
		message := new(types.Message)
		err := json.Unmarshal(m.Body, message)
		if err != nil {
			log.Println("Mensagem inválida:", err.Error())
		} else {
			log.Printf("Mensagem recebida: %+v\n", message)
		}
	}
}

func (c *consumer) Close() error {
	return c.ch.Close()
}

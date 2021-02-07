package rmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type Producer interface {
	Publish(data interface{}) error
	Close() error
}

type producer struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func NewProducer(conn *amqp.Connection, queue string) (Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	return &producer{ch: ch, q: q}, err
}

func (p *producer) Publish(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.ch.Publish("", // exchange
		p.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
}

func (p *producer) Close() error {
	return p.ch.Close()
}

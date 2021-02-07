package rmq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func NewConnection() (*amqp.Connection, error) {
	rmqUser := os.Getenv("RABBITMQ_USER")
	rmqPass := os.Getenv("RABBITMQ_PASS")
	rmqHost := os.Getenv("RABBITMQ_HOST")
	rmqPort := os.Getenv("RABBITMQ_PORT")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s", rmqUser, rmqPass, rmqHost, rmqPort)

	log.Println("rabbitmq url:", url)

	return amqp.Dial(url)
}

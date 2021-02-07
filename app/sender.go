package main

import (
	"app/rmq"
	"app/types"
	"log"
	"time"
)

func main() {
	conn, err := rmq.NewConnection()
	if err != nil {
		log.Fatalln("Erro na conexão com o RabbitMQ:", err.Error())
	}
	defer conn.Close()

	producer, err := rmq.NewProducer(conn, "hello")
	if err != nil {
		log.Fatalln("Erro ao criar um Producer:", err.Error())
	}
	defer producer.Close()

	msg := &types.Message{
		Title:     "RabbitMQ",
		Content:   "Hello World",
		Author:    "John",
		CreatedAt: time.Now(),
	}

	err = producer.Publish(msg)
	if err != nil {
		log.Fatalln("Erro ao publicar mensagem:", err.Error())
	}

	log.Printf("Mensagem publicada: %+v\n", msg)
}

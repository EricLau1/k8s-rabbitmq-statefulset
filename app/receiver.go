package main

import (
	"app/rmq"
	"log"
)

func main() {
	conn, err := rmq.NewConnection()
	if err != nil {
		log.Fatalln("Erro na conexão com o RabbitMQ:", err.Error())
	}
	defer conn.Close()

	consumer, err := rmq.NewConsumer(conn, "hello")
	if err != nil {
		log.Fatalln("Erro ao criar um Consumer:", err.Error())
	}
	defer consumer.Close()

	forever := make(chan bool)

	go consumer.WaitMessages()

	log.Printf(" [*] Esperando mensagens. Pressione CTRL+C para sair")

	<-forever
}

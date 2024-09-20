package main

import (
	"log"

	"github.com/streadway/amqp"
)

type Conusmer struct {
	channel *amqp.Channel
    errMsg  chan string
}

func (c *Conusmer) Consume(ch *amqp.Channel, q amqp.Queue) {
	defer ch.Close()
	msgs, _ := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if msgs == nil {
		log.Print("Nil channel")
	}
	go func() {
		//ProcessOrders(msgs)
	}()
}

func (c *Conusmer) Listen() {
    
}

func NewConsumer(conn *amqp.Connection) *Conusmer{
    ch, err := conn.Channel()
    if err != nil{
        return err
    }
    return &Conusmer{
        channel: ch,
    }
}

package queue

import (
	"log"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/streadway/amqp"
)

var _ contracts.IQueue = (*RabbitMQQueue)(nil)

type RabbitMQQueue struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQQueue(url string) contracts.IQueue {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Erro ao conectar no RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao abrir um canal: %v", err)
	}

	return &RabbitMQQueue{conn: conn, channel: ch}
}

func (r *RabbitMQQueue) Publish(queue string, message []byte) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = r.channel.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        message,
	})

	return err
}

func (r *RabbitMQQueue) Consume(queue string, worker func([]byte)) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := r.channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			worker(msg.Body)
		}
	}()

	return nil
}

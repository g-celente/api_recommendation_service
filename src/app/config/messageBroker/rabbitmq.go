package messagebroker

import (
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQ(dsn string, queue string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn:      conn,
		channel:   ch,
		queueName: queue,
	}, nil
}

func (r *RabbitMQ) ConsumeMessages() (<-chan amqp091.Delivery, error) {
	msgs, err := r.channel.Consume(
		r.queueName, "", true, false, false, false, nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

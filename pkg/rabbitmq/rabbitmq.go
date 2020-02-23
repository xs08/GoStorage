package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// RabbitMQ is an implementation of mq handler
type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	exchange string
}

// New RabbitMQ
func New(s string) *RabbitMQ {
	conn, e := amqp.Dial(s)
	if e != nil {
		panic(e)
	}

	ch, e := conn.Channel()
	if e != nil {
		panic(e)
	}

	q, e := ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if e != nil {
		panic(e)
	}

	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name

	return mq
}

// Bind to exchange
func (mq *RabbitMQ) Bind(exchange string) {
	e := mq.channel.QueueBind(
		mq.Name,  // queue name
		"",       // routing key
		exchange, // exchange
		false,    //
		nil,      // arguments
	)
	if e != nil {
		panic(e)
	}
	mq.exchange = exchange
}

// Send msg
func (mq *RabbitMQ) Send(queue string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}

	e = mq.channel.Publish(
		"",    // queue name
		queue, //
		false, //
		false, //
		amqp.Publishing{
			ReplyTo: mq.Name,
			Body:    []byte(str),
		},
	)
	if e != nil {
		panic(e)
	}
}

// Publish msg
func (mq *RabbitMQ) Publish(exchange string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = mq.channel.Publish(
		exchange, // exchange
		"",       //
		false,    //
		false,    //
		amqp.Publishing{
			ReplyTo: mq.Name,
			Body:    []byte(str),
		},
	)
	if e != nil {
		panic(e)
	}
}

// Consume msg
func (mq *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := mq.channel.Consume(
		mq.Name, //
		"",      //
		true,    //
		false,   //
		false,   //
		false,   //
		nil,     //
	)
	if e != nil {
		panic(e)
	}
	return c
}

// Close connection
func (mq *RabbitMQ) Close() {
	mq.channel.Close()
}

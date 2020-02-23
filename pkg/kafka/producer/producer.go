package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"
	partition := 0

	fmt.Println("start connection")

	conn, err := kafka.DialLeader(context.Background(), "tcp", "172.17.0.4:9092", topic, partition)
	if err != nil {
		fmt.Printf("dial error: %v\n", err)
		return
	}

	fmt.Println("connection ok")

	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	conn.Close()
	fmt.Println("Send ok")
}

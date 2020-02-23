package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// to consume messages
	topic := "my-topic"
	partition := 0

	fmt.Println("connect start")

	conn, err := kafka.DialLeader(context.Background(), "tcp", "172.17.0.4:9092", topic, partition)
	if err != nil {
		fmt.Errorf("dial error: %v", err)
		return
	}
	fmt.Printf("Listen ok\n")

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	// fetch 10KB min, 1MB max
	batch := conn.ReadBatch(10e3, 1e6)

	// 10KB max per message
	b := make([]byte, 10e3)
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	batch.Close()
	conn.Close()
}

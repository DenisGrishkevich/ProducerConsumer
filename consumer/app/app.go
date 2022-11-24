package app

import (
	"consumer/repository"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type consumer struct{}

func NewConsumer() *consumer {
	return &consumer{}
}

func normalizeId(unnormalizedId string) string {
	unnormalizedId = strings.Replace(unnormalizedId, "%3A", ":", -1)
	unnormalizedId = strings.Replace(unnormalizedId, "%3a", ":", -1)
	if strings.HasPrefix(unnormalizedId, "ac:") {
		unnormalizedParts := strings.SplitN(unnormalizedId, ":", 3)[:2]
		return strings.Join(unnormalizedParts, ":")
	}

	return strings.SplitN(unnormalizedId, ":", 2)[0]
}

func (c *consumer) consumeMessage(ch *amqp.Channel, q amqp.Queue) {
	msgs, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	go func() {
		for d := range msgs {
			message := normalizeId(string(d.Body))
			fmt.Printf("Consuming message: %v\n", message)
			repository.InsertMessageIntoMongo(message)
		}
		time.Sleep(500 * time.Millisecond)
	}()
}

func (c *consumer) Run(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	ch, q := repository.RabbitInit()
	repository.MongoDBInit()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Consumer exit.")
			return
		default:
			c.consumeMessage(ch, q)
		}
	}
}

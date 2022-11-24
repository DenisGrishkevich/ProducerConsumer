package app

import (
	"bufio"
	"context"
	"fmt"
	"producer/repository"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type producer struct {
	ids      []string
}

func NewProducer(filePath string) *producer {
	p := &producer{}
	p.ids = openAndReadFile(filePath)
	return p
}

func openAndReadFile(filePath string) ([]string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic("File not found")
	}
	defer file.Close()

	var ids []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	return ids
}

func (p *producer) createMessage(ch *amqp.Channel, q amqp.Queue) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(len(p.ids))
	message := p.ids[randomNum]
	fmt.Printf("Producing message: %v\n", message)
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "application/json",
		  Body: []byte(message),
	})
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(100 * time.Millisecond)
}

func (p *producer) Run(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	ch, q := repository.RabbitInit()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("Received interrupt, shutting down...")
			fmt.Println("Context Done - closing channel.")
			ch.Close()
			fmt.Println("Producer exit.")
			return
		default:
			p.createMessage(ch, q)
		}
	}
}
package repository

import (
	"context"
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func MongoDBInit() {
    clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = (*mongo.Collection)(client.Database("MongoMessages").Collection("Messages"))
}

func InsertMessageIntoMongo(message string) error {
	newMessage := bson.D{{"message", message}}
	_, err := collection.InsertOne(ctx, newMessage)
	return err
}

func RabbitInit() (ch *amqp.Channel, q amqp.Queue) {
	conn, _ := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	q, err = ch.QueueDeclare(
		"messages", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	return ch, q
}
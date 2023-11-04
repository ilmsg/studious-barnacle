package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("auth-api")

	mongoURL := os.Getenv("mongoURL")
	dbName := os.Getenv("dbName")
	collectionName := os.Getenv("collectionName")

	_ = getMongoCollection(mongoURL, dbName, collectionName)

	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	groupID := os.Getenv("groupID")

	kafkaWriter := getKafkaWriter(kafkaURL, topic)
	defer kafkaWriter.Close()

	kafkaRreader := getKafkaReader(kafkaURL, topic, groupID)
	defer kafkaRreader.Close()

	fmt.Println("string consuming ... !!!")

	// for {
	// 	msg, err := kafkaRreader.ReadMessage(context.Background())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("messate at topic:%v partition:%v offset:%v %s = %s\n",
	// 		msg.Topic,
	// 		msg.Partition,
	// 		msg.Offset,
	// 		string(msg.Key),
	// 		string(msg.Value),
	// 	)
	// }

	// /auth/register 	{ email, password }
	// /auth/login		{ email, password }
	// /auth/profile	{ first_name, last_name, birthday}, jwt token
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth-api"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e3, // 10KB
	})
}

func getMongoCollection(mongoURL, dbName, collectionName string) *mongo.Collection {
	opts := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB ... !!!")
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return collection
}

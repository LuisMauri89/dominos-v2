package service

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "dominos-orders-topic"
	brokerAddress = "localhost:9092"
	groupID       = "deliveries-group-id"
)

type Payload struct {
	ID         string
	Name       string
	Address    string
	Items      map[string]string
	TotalPrice float64
	Action     string
}

type KafkaService interface {
	StartKafkaListener(ctx context.Context, listener chan Payload)
}

type kafkaService struct {
}

func NewKafkaService() KafkaService {
	return &kafkaService{}
}

func (s *kafkaService) StartKafkaListener(ctx context.Context, listener chan Payload) {
	go consume(ctx, listener)
}

func consume(ctx context.Context, listener chan Payload) {
	logger := log.New(os.Stdout, "kafka reader: ", 0)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupID,
		Logger:  logger,
	})
	for {
		payloadMsg, err := reader.ReadMessage(ctx)
		if err != nil {
			logger.Fatal("could not read message " + err.Error())
		}

		payload := Payload{}
		err = json.Unmarshal([]byte(string(payloadMsg.Value)), &payload)
		if err != nil {
			logger.Fatal("could not parse message " + err.Error())
		}
		listener <- payload
		logger.Println("received: ", string(payloadMsg.Value))
	}
}

package repository

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
)

// UserRepository represent the user's repository contract
type UserRepository interface {
	Publish(ctx context.Context, data string, topic string) error
	GetUUID(ctx context.Context, uuid string) (string, error)
}

type UserRepositoryImpl struct {
	Redis *redis.Client
	Kafka *kafka.Producer
	// RabbitMQ *amqp091.Connection
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
// func NewUserRepository(rabbitConn *amqp091.Connection) UserRepository {
func NewUserRepository(redisConnect *redis.Client, kafkaProduce *kafka.Producer) UserRepository {
	return &UserRepositoryImpl{
		Redis: redisConnect,
		Kafka: kafkaProduce,
	}
}

func (m *UserRepositoryImpl) Publish(ctx context.Context, data string, topic string) error {
	err := m.Kafka.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sending message success: %s", data)

	return err
}

func (m *UserRepositoryImpl) GetUUID(ctx context.Context, uuid string) (res string, err error) {
	res, err = m.Redis.Get(ctx, uuid).Result()
	return res, err
}

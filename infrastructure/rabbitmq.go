package infrastructure

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenRabbitMQ() (conn *amqp.Connection, err error) {
	dsn := os.Getenv("RABBITMQ_USERNAME") + ":" + os.Getenv("RABBITMQ_PASSWORD") + "@" + os.Getenv("RABBITMQ_HOST") + ":" + os.Getenv("RABBITMQ_PORT") + "/"
	conn, err = amqp.Dial("amqp://" + dsn)

	return conn, err
}

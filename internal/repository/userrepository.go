package repository

// UserRepository represent the user's repository contract
type UserRepository interface {
	// Publish(ctx context.Context, data string)
}

type UserRepositoryImpl struct {
	// RabbitMQ *amqp091.Connection
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
// func NewUserRepository(rabbitConn *amqp091.Connection) UserRepository {
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		// RabbitMQ: rabbitConn,
	}
}

// func (m *UserRepositoryImpl) Publish(ctx context.Context, data string) {
// 	ch, err := m.RabbitMQ.Channel()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"logger", //name
// 		false,    // durable
// 		false,    //delete when unused
// 		false,    // exclusive
// 		false,    // no-wait
// 		nil,      // arguments
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// body := "Hi halovina, keep in touch"
// 	err = ch.PublishWithContext(
// 		ctx,
// 		"",     // exchange
// 		q.Name, // routing key
// 		false,  // mandatory
// 		false,  // immadiate
// 		amqp091.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(data),
// 		})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("Sending message success: %s", data)
// }

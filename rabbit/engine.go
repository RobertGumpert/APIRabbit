package rabbit

import (
	"errors"
	"github.com/streadway/amqp"
)

type Engine struct {
	ConnectionString string
	Connection       *amqp.Connection
	Chanel           *amqp.Channel
	Queues           map[string]*amqp.Queue
}

func NewEngine() *Engine {
	return &Engine{
		ConnectionString: "",
		Connection:       nil,
		Chanel:           nil,
		Queues:           make(map[string]*amqp.Queue),
	}
}

func (engine *Engine) createConnection(connectionString string) error {
	conn, err := amqp.Dial(connectionString)
	if err == nil {
		engine.ConnectionString = connectionString
		engine.Connection = conn
		return nil
	} else {
		engine.CloseRabbitConnection()
		return errors.New("Error : rabbit.createConnection : 'err := amqp.Dial(connectionString)' not nil ")
	}
}

func (engine *Engine) createRabbitChanel() error {
	if engine.Connection != nil {
		ch, err := engine.Connection.Channel()
		if err == nil {
			engine.Chanel = ch
			return nil
		} else {
			engine.CloseRabbitConnection()
			return errors.New("Error : rabbit.createRabbitChanel : 'err := engine.Connection.Channel' not nil ")
		}
	}
	return errors.New("Error : rabbit.createRabbitChanel : 'engine.Connection' is nil ")
}

func (engine *Engine) createRabbitQueue(queueName string, durable, delete, exclusive, nowait bool, args map[string]interface{}) error {
	if engine.Chanel != nil {
		queue, err := engine.Chanel.QueueDeclare(
			queueName,
			durable,
			delete,
			exclusive,
			nowait,
			args,
		)
		if err == nil {
			engine.Queues[queueName] = &queue
			return nil
		} else {
			engine.CloseRabbitConnection()
			return errors.New("Error : rabbit.createRabbitQueue : 'err := engine.Chanel.QueueDeclare' not nil ")
		}
	}
	return errors.New("Error : rabbit.createRabbitQueue : 'engine.Chanel' is nil ")
}

func (engine *Engine) createRabbitConsume(queueName, consumer string, ack, exclusive, nolocal, nowait bool, args map[string]interface{}) (<-chan amqp.Delivery, error) {
	messages, err := engine.Chanel.Consume(
		engine.Queues[queueName].Name,
		consumer,
		ack,
		exclusive,
		nolocal,
		nowait,
		args,
	)
	if err != nil {
		defer engine.CloseRabbitConnection()
		return nil, errors.New("Error : rabbit.ReceiveMessages : 'err := rabbitEngine.Chanel.Consume' not nil ")
	}
	return messages, nil
}

func (engine *Engine) CloseRabbitConnection() {
	if engine.Connection != nil {
		_ = engine.Connection.Close()
	}
}

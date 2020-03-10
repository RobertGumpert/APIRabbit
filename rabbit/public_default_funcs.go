package rabbit

import (
	"errors"
	"github.com/streadway/amqp"
)

type CreatorQueue func(
	durable,
	delete,
	exclusive,
	nowait bool,
	args map[string]interface{},
) error
type CreatorConsumer func(
	consumer string,
	ack,
	exclusive,
	nolocal,
	nowait bool,
	args map[string]interface{},
) (<-chan amqp.Delivery, error)

func StartRabbitEngine(connectionString string) *Engine {
	// "amqp://guest:guest@localhost:5672/"
	engine := NewEngine()
	if err := engine.createConnection(connectionString); err != nil {
	}
	if err := engine.createRabbitChanel(); err != nil {
	}
	return engine
}

func CreateRabbitConsumer(queueName string, engine *Engine) (CreatorQueue, CreatorConsumer, error) {
	if engine != nil {
		if q := engine.Queues[queueName]; q == nil {
			var (
				creatorQueue = func(durable, delete, exclusive, nowait bool, args map[string]interface{}) error {
					return engine.createRabbitQueue(queueName, durable, delete, exclusive, nowait, args)
				}
				creatorConsumer = func(consumer string, ack, exclusive, nolocal, nowait bool, args map[string]interface{}) (<-chan amqp.Delivery, error) {
					consume, err := engine.createRabbitConsume(queueName, consumer, ack, exclusive, nolocal, nowait, args)
					return consume, err
				}
			)
			return creatorQueue, creatorConsumer, nil
		}
	}
	return nil, nil, errors.New("Error")
}

func CreateRabbitSender(queueName string, engine *Engine){

}
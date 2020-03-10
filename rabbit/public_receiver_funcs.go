package rabbit

import (
	"github.com/streadway/amqp"
)

type Receiver struct {
	Engine          *Engine
	Delivery        *amqp.Delivery
	CreatorQueue    CreatorQueue
	CreatorConsumer CreatorConsumer
	Starter         func(*Receiver)
}

func CreateReceiver(queueName string, engine *Engine, starter func(*Receiver)) (*Receiver, error) {
	if creatorQueue, creatorConsumer, err := CreateRabbitConsumer(queueName, engine); err == nil && creatorQueue != nil && creatorConsumer != nil {
		receiver := &Receiver{
			Engine:          engine,
			CreatorQueue:    creatorQueue,
			CreatorConsumer: creatorConsumer,
			Starter:         starter,
		}
		starter(receiver)
		return receiver, nil
	} else {
		return nil, err
	}
}



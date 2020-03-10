package car

import (
	"../../rabbit"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"strings"
)

type RequestHandler func(car *Car) string

func CarReceiver(receiver *rabbit.Receiver) {
	localMessagesChanel := make(chan string)
	if errQueue := receiver.CreatorQueue(false, false, false, false, nil); errQueue != nil {
		fmt.Println("errQueue : ", errQueue)
	}
	if rabbitMessagesChanel, errReceiver := receiver.CreatorConsumer("", false, false, false, false, nil); errReceiver != nil {
		fmt.Println("errReceiver : ", errReceiver)
	} else {
		go readerMessages(rabbitMessagesChanel, localMessagesChanel)
		handlerMessages(localMessagesChanel)
	}
}

func Route(car *Car, path, endpoint string, r RequestHandler) {
	if strings.Contains(path, endpoint) {
		res := r(car)
		fmt.Println(res)
	}
}

func Index(car *Car, path string) {
	Route(car, path, "/add", AddElement)
	Route(car, path, "/delete", DeleteElement)
	/*
		/car/add:{ "ID": 12, "Name": "BMW" }
		/car/delete:{ "ID": 12, "Name": "BMW" }
	*/
}

func readerMessages(rabbitMessagesChanel <-chan amqp.Delivery, localMessagesChanel chan string) {
	for message := range rabbitMessagesChanel {
		result := string(message.Body)
		localMessagesChanel <- result
	}
}

func handlerMessages(localMessagesChanel <-chan string) {
	for message := range localMessagesChanel {
		slice := strings.Split(message, ":")
		var (
			car      = &Car{}
			endpoint = slice[0]
		)
		if err := json.Unmarshal([]byte(slice[1]), car); err != nil {
			Index(car, endpoint)
		}
	}
}



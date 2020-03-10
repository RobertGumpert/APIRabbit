package main

import (
	"./models/car"
	"./rabbit"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	engine := rabbit.StartRabbitEngine("amqp://guest:guest@localhost:5672/")
	_, _ = rabbit.CreateReceiver("queue", engine, car.CarReceiver)
	for scanner.Scan() {
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("Events.new", func(msg *nats.Msg) {
		fmt.Printf("Recived event: %s\n", string(msg.Data))
	})

	err = nc.Publish("Events.new", []byte("Hello from server event system"))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running ")
}

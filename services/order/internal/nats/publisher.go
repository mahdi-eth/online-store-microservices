package nats

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func InitNATS(url string) {
    var err error

	natsURL := os.Getenv("NATS_URL")
    if natsURL == "" {
        natsURL = nats.DefaultURL
    }

    nc, err = nats.Connect(natsURL)
    if err != nil {
        log.Fatalf("Error connecting to NATS: %v", err)
    }
    log.Println("Connected to NATS")
}

func PublishOrderCreated(orderID string) {
    err := nc.Publish("order.created", []byte(orderID))
    if err != nil {
        log.Printf("Error publishing order created event: %v", err)
    }
}
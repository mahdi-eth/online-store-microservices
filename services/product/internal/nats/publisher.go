package nats

import (
    "github.com/nats-io/nats.go"
    "log"
)

func publishProductUpdated(productID string) {
    nc, _ := nats.Connect(nats.DefaultURL)
    defer nc.Close()

    err := nc.Publish("product.updated", []byte(productID))
    if err != nil {
        log.Fatal(err)
    }
}

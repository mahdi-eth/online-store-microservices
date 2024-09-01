package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mahdi-eth/online-store/inventory-service/internal/controllers"
	"github.com/nats-io/nats.go"
)

var natsConn *nats.Conn

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/inventory", controllers.GetInventory).Methods("GET")
    r.HandleFunc("/inventory/{product_id}", controllers.GetInventoryItem).Methods("GET")
    r.HandleFunc("/inventory", controllers.CreateInventoryItem).Methods("POST")
    r.HandleFunc("/inventory/{product_id}", controllers.UpdateInventoryItem).Methods("PUT")
    r.HandleFunc("/inventory/{product_id}", controllers.DeleteInventoryItem).Methods("DELETE")

    natsURL := os.Getenv("NATS_URL")
    if natsURL == "" {
        natsURL = nats.DefaultURL
    }

    var err error
    natsConn, err = nats.Connect(natsURL)
    if err != nil {
        log.Fatalf("Error connecting to NATS: %v", err)
    }
    defer natsConn.Close()

    log.Println("Inventory Service is running on port 50053")
    log.Fatal(http.ListenAndServe(":50053", r))
}

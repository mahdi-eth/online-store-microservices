package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mahdi-eth/online-store/product-service/internal/controllers"
	"github.com/nats-io/nats.go"
)

var natsConn *nats.Conn

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
    r.HandleFunc("/products/{id}", controllers.GetProduct).Methods("GET")
    r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
    r.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

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

    log.Println("Products service is running on port 50051")
    log.Fatal(http.ListenAndServe(":50051", r))
}
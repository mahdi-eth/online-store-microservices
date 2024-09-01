package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mahdi-eth/online-store/order-service/internal/controllers"
	"github.com/nats-io/nats.go"
)

var natsConn *nats.Conn

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
    r.HandleFunc("/orders/{id}", controllers.GetOrder).Methods("GET")
    r.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
    r.HandleFunc("/orders/{id}", controllers.UpdateOrder).Methods("PUT")
    r.HandleFunc("/orders/{id}", controllers.DeleteOrder).Methods("DELETE")

    log.Println("Order Service is running on port 50052")
    log.Fatal(http.ListenAndServe(":50052", r))
}
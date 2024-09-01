package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/mahdi-eth/online-store/inventory-service/internal/db"
)

// GetInventory retrieves the inventory item based on product ID from Redis.
func GetInventory(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	productID := mux.Vars(r)["product_id"]
	item, err := db.RedisClient.Get(ctx, productID).Result()
	if err != nil {
		if err == redis.Nil {
			http.Error(w, "Inventory item not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var itemData map[string]interface{}
	if err := json.Unmarshal([]byte(item), &itemData); err != nil {
		http.Error(w, "Error parsing inventory item data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemData)
}

// GetInventoryItem retrieves a specific inventory item by its ID.
func GetInventoryItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	itemID := mux.Vars(r)["id"]
	item, err := db.RedisClient.Get(ctx, itemID).Result()
	if err != nil {
		if err == redis.Nil {
			http.Error(w, "Inventory item not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var itemData map[string]interface{}
	if err := json.Unmarshal([]byte(item), &itemData); err != nil {
		http.Error(w, "Error parsing inventory item data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemData)
}

// CreateInventoryItem creates a new inventory item and stores it in Redis.
func CreateInventoryItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	var itemData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&itemData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productID, ok := itemData["product_id"].(string)
	if !ok {
		http.Error(w, "Invalid or missing product_id", http.StatusBadRequest)
		return
	}

	itemJSON, err := json.Marshal(itemData)
	if err != nil {
		http.Error(w, "Error serializing inventory item data", http.StatusInternalServerError)
		return
	}

	err = db.RedisClient.Set(ctx, productID, itemJSON, 0).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(itemData)
}

// UpdateInventoryItem updates an existing inventory item in Redis.
func UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	itemID := mux.Vars(r)["id"]

	var itemData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&itemData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemJSON, err := json.Marshal(itemData)
	if err != nil {
		http.Error(w, "Error serializing inventory item data", http.StatusInternalServerError)
		return
	}

	err = db.RedisClient.Set(ctx, itemID, itemJSON, 0).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemData)
}

// DeleteInventoryItem deletes an inventory item from Redis by its ID.
func DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	itemID := mux.Vars(r)["id"]

	err := db.RedisClient.Del(ctx, itemID).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

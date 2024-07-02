package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Item represents an item in the shop
type Item struct {
	ItemID string  `json:"itemId"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Stock  int     `json:"stock"`
	Dates  *Dates  `json:"dates"`
}

// Dates represents manufacture and expiry dates
type Dates struct {
	ManufactureDate string `json:"manufactureDate"`
	ExpiryDate      string `json:"expiryDate"`
}

// shopDB is a fake database storing items
var shopDB []Item

func main() {
	// Seed the random number generator
	// rand.Seed(time.Now().UnixNano())

	// Initialize the router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/items", getAllItems).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", getOneItem).Methods("GET")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

// homePage handles the root route
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<h2> Welcome to the shop API </h2>`))
}

// getAllItems returns all items in the shopDB
func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shopDB)
}

// getOneItem returns a single item by its ID
func getOneItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the item ID from the request
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Loop through shopDB to find the item with the given ID
	for _, item := range shopDB {
		if item.ItemID == itemID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// If item not found, return 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Item not found")
}

// createItem creates a new item in the shopDB
func createItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create item")
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON request body into item struct
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validate item fields
	if item.isEmpty() {
		http.Error(w, "Item cannot be empty", http.StatusBadRequest)
		return
	}

	// Generate a new UUID for itemID
	item.ItemID = uuid.New().String()

	// Add item to shopDB
	shopDB = append(shopDB, item)

	// Encode the created item as JSON and send it in the response
	json.NewEncoder(w).Encode(item)
}

// updateItem updates an existing item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update item")
	w.Header().Set("Content-Type", "application/json")

	// Get the item ID from the request
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Loop through shopDB to find the item with the given ID
	for index, item := range shopDB {
		if item.ItemID == itemID {
			// Remove the old item
			shopDB = append(shopDB[:index], shopDB[index+1:]...)

			// Decode JSON request body into updatedItem struct
			var updatedItem Item
			if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
				http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
				return
			}

			// Assign the same itemID
			updatedItem.ItemID = itemID

			// Add updated item to shopDB
			shopDB = append(shopDB, updatedItem)

			// Encode the updated item as JSON and send it in the response
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	// If item not found, return 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Item not found")
}

// deleteItem deletes an item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete item")
	w.Header().Set("Content-Type", "application/json")

	// Get the item ID from the request
	vars := mux.Vars(r)
	itemID := vars["id"]

	// Loop through shopDB to find the item with the given ID and delete it
	for index, item := range shopDB {
		if item.ItemID == itemID {
			shopDB = append(shopDB[:index], shopDB[index+1:]...)
			json.NewEncoder(w).Encode("Item deleted")
			return
		}
	}

	// If item not found, return 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Item not found")
}

// isEmpty checks if an item is empty
func (i *Item) isEmpty() bool {
	return i.ItemID == "" && i.Name == "" && i.Price == 0 && i.Stock == 0 && i.Dates == nil
}

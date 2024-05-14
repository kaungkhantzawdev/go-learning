package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Item struct {
	ID string
	Name string
	Price float64
}

var items = map[string]Item{}

// CREATE
func create(w http.ResponseWriter, _ *http.Request) {

	var newItem Item

	newItem.ID = "1"
	newItem.Name = "Test One"
	newItem.Price = 2.1

	items[newItem.ID] = newItem

	fmt.Fprintf(w, "successfully created.")

}

// READ
func read(w http.ResponseWriter, r*http.Request) {
	
	itemID := r.URL.Query().Get("id")

	item, found := items[itemID]

	if !found {

		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "Item ID: %s, Name: %s, Price: %2f", item.ID, item.Name, item.Price)
}

// READ
func readAll(w http.ResponseWriter, r*http.Request) {

	jsonData, err := json.Marshal(items)

	if err != nil {
		http.Error(w, "unable to marshal items to JSON", http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-type", "application/json")

	w.Write(jsonData)
}


func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/read-all", readAll)

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil);
	
}

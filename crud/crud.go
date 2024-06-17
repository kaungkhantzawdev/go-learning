package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID string
	Name string
	Price float64
}

/** UPDATE */
func update(w http.ResponseWriter, r*http.Request) {


	if r.Method != http.MethodPatch {
		http.Error(w, "Method's not allowed.", http.StatusMethodNotAllowed)
		return
	}

	var updateItem Item

	err := json.NewDecoder(r.Body).Decode(&updateItem)

	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	itemID := updateItem.ID
	if itemID == "" {
		http.Error(w, "Missing item id", http.StatusBadRequest)
	}

	item, found := items[updateItem.ID]

	if !found {
		http.NotFound(w, r)
		return
	}

	item.Name = updateItem.Name
	item.Price = updateItem.Price
	items[itemID] = item

	fmt.Fprintln(w, item.Name)
}


var items = map[string]Item{}

// CREATE
func create(w http.ResponseWriter, r*http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method's not allowed.", http.StatusMethodNotAllowed)
		return
	}

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

// READ ALL DATA
func readAll(w http.ResponseWriter, r*http.Request) {

	jsonData, err := json.Marshal(items)

	if err != nil {
		http.Error(w, "unable to marshal items to JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")

	w.Write(jsonData)
}


// DELETE
func remove(w http.ResponseWriter, r*http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method's not allowed.", http.StatusMethodNotAllowed)
		return
	}

	itemID := r.URL.Query().Get("id")

	_, found := items[itemID]
	if !found {
		http.NotFound(w, r)
		return
	}

	delete(items, itemID)
	fmt.Fprintf(w, "successfully deleted item with ID %s", itemID)

}

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/read-all", readAll)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", remove)

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil);

}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Details for a specific item in a list
type Item struct {
	Due         time.Time `json:"due"`
	Completed   bool      `json:"completed"`
	CompletedOn time.Time `json:"completedOn"`
	Description string    `json:"description"`
}
type Items []Item

func GetListItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ALL THE LIST ITEMS!")
}

func GetListItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["ListID"]
	itemID := vars["ItemID"]
	fmt.Fprintln(w, "Acquiring list: ", listID, " Acquiring item: ", itemID)
}

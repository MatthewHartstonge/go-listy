package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Details for a list that contains multiple items
type List struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Should change items to a URL string as it makes items redundant
	Items *Items `json:"items"`
}
type Lists []List

func GetLists(w http.ResponseWriter, r *http.Request) {
	lists := Lists{
		List{
			ID:   1,
			Name: "Groceries",
			Items: &Items{
				Item{Description: "Chicken"},
				Item{Description: "Chocolate"},
				Item{Description: "Jim Beam"}}},
		List{
			ID:   2,
			Name: "CPUs",
			Items: &Items{
				Item{Description: "Intel"},
				Item{Description: "AMD"},
				Item{Description: "ARM"},
				Item{Description: "MIPS"}}}}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(lists)
	if err != nil {
		log.Fatal(err)
	}
}

func GetList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["ListID"]
	fmt.Fprintln(w, "Aqquiring list: ", listID)
}

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	fmt.Fprintf(w, "Hello, there%s", "visitor")

}

// GET -> recibe
func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

//POST -> tipo de entrega

func addCountry(w http.ResponseWriter, r *http.Request) {
	country := &Country{}

	err := json.NewDecoder(r.Body).Decode(country)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	countries = append(countries, country)
	fmt.Fprintf(w, "Country was added")
}

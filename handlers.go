package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome visitor\n")
}

func ItinerariesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//TODO: Display recent posted itineraries
	itineraries := Itineraries{
		Itinerary{Name: "Rome"},
		Itinerary{Name: "Naples"},
	}

	if err := json.NewEncoder(w).Encode(itineraries); err != nil {
		panic(err)
	}
}

func ItinerariesGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	itineraryId := vars["itineraryId"]
	fmt.Fprintf(w, "TODO: Get the itinerary with the given index number: ", itineraryId)
}

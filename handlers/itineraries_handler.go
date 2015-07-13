package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mishuan/itinerary_service/models"
	"net/http"
)

func ItinerariesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//TODO: Display recent posted itineraries
	itineraries := models.Itineraries{
		models.Itinerary{Name: "Rome"},
		models.Itinerary{Name: "Naples"},
	}

	if err := json.NewEncoder(w).Encode(itineraries); err != nil {
		panic(err)
	}
}

func ItinerariesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	itineraryId := vars["itineraryId"]
	fmt.Fprintf(w, "TODO: Get the itinerary with the given index number: ", itineraryId)
}

func ItinerariesCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

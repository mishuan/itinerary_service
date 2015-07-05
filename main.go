package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Server started and listening to port 8080.\n")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/itineraries", ItinerariesIndex)
	router.HandleFunc("/itineraries/{itineraryId}", ItinerariesGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ItinerariesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: Display recent posted itineraries")
	itineraries := Itineraries{
		Itinerary{Name: "Rome"},
		Itinerary{Name: "Naples"},
	}

	json.NewEncoder(w).Encode(itineraries)
}

func ItinerariesGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itineraryId := vars["itineraryId"]
	fmt.Fprintf(w, "TODO: Get the itinerary with the given index number: ", itineraryId)
}

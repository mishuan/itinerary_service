package main

import (
	"github.com/mishuan/itinerary_service/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", handlers.Index},
	Route{"Itineraries Index", "GET", "/itineraries", handlers.ItinerariesIndex},
	Route{"Get Specific Itinerary", "GET", "/itineraries/{itineraryId}", handlers.ItinerariesGet},
	Route{"Create New Itinerary", "POST", "/itineraries/create", handlers.ItinerariesCreate},
}

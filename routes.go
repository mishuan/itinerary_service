package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"Itineraries Index", "GET", "/itineraries", ItinerariesIndex},
	Route{"Get Specific Itinerary", "GET", "/itineraries/{itineraryId}", ItinerariesGet},
}

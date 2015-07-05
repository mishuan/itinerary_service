package main

import "time"

type Itinerary struct {
	Name      string    `json:"name"`
	VisitDate time.Time `json: "visit_date"`
	Duration  int       `json: "duration"`
	Country   string    `json: "country"`
	City      string    `json: "city"`
	Rating    int       `json: "rating"`
	Tansit    string    `json: "transit"`
	Plan      []string  `json: "plan"`
	Tips      []string  `json: "tips"`
	Tags      string    `json: "tags"`
}

type Itineraries []Itinerary

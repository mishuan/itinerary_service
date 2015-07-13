package models

import "time"

type Itinerary struct {
	Name      string      `json: "name"`
	VisitDate time.Time   `json: "visit_date"`
	Duration  int         `json: "duration"`
	Country   string      `json: "country"`
	City      string      `json: "city"`
	Rating    int         `json: "rating"`
	Plan      Attractions `json: "plan"`
	Tips      []string    `json: "tips"`
	Tags      string      `json: "tags"`
}

type Itineraries []Itinerary

package models

type Attraction struct {
	Name        string `json: "name"`
	Description string `json: "description"`
}

type Attractions []Attraction

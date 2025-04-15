package models

import "time"

type Author struct {
	Name           string   `json:"name" bson:"name"`
	Biography      string   `json:"biography" bson:"biography"`
	BirthDate      string   `json:"birth_date" bson:"birth_date"`
	DeathDate      string   `json:"death_date" bson:"death_date"`
	DeathPlace     string   `json:"death_place" bson:"death_place"`
	Congregation   string   `json:"congregation" bson:"congregation"`
	Denomination   []string `json:"denomination" bson:"denomination"`
	VenerationDate string   `json:"veneration_date" bson:"veneration_date"`
	PilgrimageSite string   `json:"pilgrimage_site" bson:"pilgrimage_site"`
	FestivalDate   string   `json:"festival_date" bson:"festival_date"`
	Attributes     []string `json:"attributes" bson:"attributes"`
	Patronage      []string `json:"patronage" bson:"patronage"`
}

type Quote struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Quote     string    `json:"quote" bson:"quote"`
	Author    Author    `json:"author" bson:"author"`
	Category  string    `json:"category" bson:"category"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

package models

import "time"

type Metadata struct {
	Source []string `json:"source" bson:"source"`
}

type Quote struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Quote     string    `json:"quote" bson:"quote"`
	Author    string    `json:"author" bson:"author"`
	Category  string    `json:"category" bson:"category"`
	Metadata  Metadata  `json:"metadata" bson:"metadata"` // Added metadata
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

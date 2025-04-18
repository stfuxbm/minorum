package models

import "time"

// Apostle mewakili struktur data umum untuk para rasul
type Saints struct {
	Name               string    `bson:"name" json:"name"`
	LatinName          string    `bson:"latin_name" json:"latin_name"`
	NickName           string    `bson:"nick_name" json:"nick_name"`
	Biography          string    `bson:"biography" json:"biography"`
	RoleInApostles     string    `bson:"role_in_apostles" json:"role_in_apostles"`
	FunFacts           []string  `bson:"fun_facts" json:"fun_facts"`
	Denomination       []string  `bson:"denomination" json:"denomination"`
	MissionaryActivity string    `bson:"missionary_activity" json:"missionary_activity"`
	FeastDay           []string  `bson:"feast_day" json:"feast_day"`
	Birth              BirthInfo `bson:"birth" json:"birth"`
	Death              DeathInfo `bson:"death" json:"death"`
	Patronage          []string  `bson:"patronage" json:"patronage"`
	Symbols            []string  `bson:"symbols" json:"symbols"`
	CreatedAt          time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt          time.Time `bson:"updated_at" json:"updated_at"`
	Metadata           Metadata  `bson:"metadata" json:"metadata"`
}

// BirthInfo menyimpan detail kelahiran
type BirthInfo struct {
	NameAtBirth string `bson:"name_at_birth" json:"name_at_birth"`
	Year        string `bson:"year" json:"year"`
	Place       string `bson:"place" json:"place"`
}

// DeathInfo menyimpan detail kematian
type DeathInfo struct {
	Year  string `bson:"year" json:"year"`
	Place string `bson:"place" json:"place"`
	Cause string `bson:"cause" json:"cause"`
}

// Metadata menyimpan informasi sumber data
type Metadata struct {
	Source            []string `bson:"source" json:"source"`
	RetrievedDate     string   `bson:"retrieved_date" json:"retrieved_date"`
	SourceDescription string   `bson:"source_description" json:"source_description"`
}

package models

import "time"

// Metadata menyimpan informasi tambahan terkait sumber kutipan
type Metadata struct {
	Source            []string `json:"source" bson:"source"`
	RetrievedDate     string   `json:"retrieved_date" bson:"retrieved_date"`
	SourceDescription string   `json:"source_description" bson:"source_description"`
}

// Quote menyimpan kutipan beserta informasi terkaitnya
type Quote struct {
	ID        string    `json:"id" bson:"_id,omitempty"`      // ID unik untuk kutipan
	Quote     string    `json:"quote" bson:"quote"`           // Isi kutipan
	Author    string    `json:"author" bson:"author"`         // Nama penulis kutipan
	Category  string    `json:"category" bson:"category"`     // Kategori kutipan (misal: Doa, Kasih)
	Metadata  Metadata  `json:"metadata" bson:"metadata"`     // Metadata terkait sumber kutipan
	CreatedAt time.Time `json:"created_at" bson:"created_at"` // Tanggal saat kutipan dibuat
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` // Tanggal saat kutipan terakhir diperbarui
}

package models

import "time"

// Ordo mewakili detail sebuah ordo
type Ordo struct {
	OrderName             string               `json:"order_name" bson:"order_name"`
	LatinName             string               `json:"latin_name" bson:"latin_name"`
	Nickname              string               `json:"nickname" bson:"nickname"`
	Title                 string               `json:"title" bson:"title"`       // e.g., OFM, SJ
	Category              string               `json:"category" bson:"category"` // e.g., Klerikal (Imam)
	Denomination          string               `json:"denomination" bson:"denomination"`
	FounderName           string               `json:"founder_name" bson:"founder_name"`
	FoundedYear           int                  `json:"founded_year" bson:"founded_year"`
	CurrentLeader         string               `json:"current_leader" bson:"current_leader"`
	Address               Address              `json:"address" bson:"address"`
	Description           string               `json:"description" bson:"description"`
	Slogan                string               `json:"slogan" bson:"slogan"`
	RelatedOrders         []string             `json:"related_orders" bson:"related_orders"`
	CoatOfArmsURL         string               `json:"coat_of_arms_url" bson:"coat_of_arms_url"`
	ConstitutionApproval  ConstitutionApproval `json:"constitution_approval" bson:"constitution_approval"`
	MissionaryActivity    MissionaryActivity   `json:"missionary_activity" bson:"missionary_activity"`
	Principles            Principles           `json:"principles" bson:"principles"`
	NotableSaintsBlesseds []NotableSaint       `json:"notable_saints_blesseds" bson:"notable_saints_blesseds"`
	Rules                 []string             `json:"rules" bson:"rules"`
	Metadata              Metadata             `json:"metadata" bson:"metadata"`
	CreatedAt             time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt             time.Time            `json:"updated_at" bson:"updated_at"`
}

// ConstitutionApproval berisi informasi tentang persetujuan konstitusi
type ConstitutionApproval struct {
	ApprovedByPope string `json:"approved_by_pope" bson:"approved_by_pope"`
	ApprovalDate   string `json:"approval_date" bson:"approval_date"`
}

// MissionaryActivity berisi informasi tentang tujuan dan wilayah misi
type MissionaryActivity struct {
	InitialGoals     string   `json:"initial_goals" bson:"initial_goals"`
	GlobalRegions    []string `json:"global_regions" bson:"global_regions"`
	IndonesiaRegions []string `json:"indonesia_regions" bson:"indonesia_regions"`
}

// Principles berisi informasi tentang prinsip dan praktik rohani ordo
type Principles struct {
	Motto             string   `json:"motto" bson:"motto"`
	SpiritualPractice []string `json:"spiritual_practice" bson:"spiritual_practice"`
}

// Address mewakili alamat untuk wilayah ordo
type Address struct {
	HeadQuarters string   `json:"head_quarters" bson:"head_quarters"`
	Street       string   `json:"street" bson:"street"`
	City         string   `json:"city" bson:"city"`
	PostalCode   string   `json:"postal_code" bson:"postal_code"`
	Country      string   `json:"country" bson:"country"`
	Phone        []string `json:"phone" bson:"phone"`
	Fax          string   `json:"fax" bson:"fax"`
}

// NotableSaint mewakili seorang santo atau beato yang terkenal dari ordo
type NotableSaint struct {
	Name         string `json:"name" bson:"name"`
	Contribution string `json:"contribution" bson:"contribution"`
}

// Metadata berisi informasi sumber data dan tanggal pengambilan data
type Metadata struct {
	Source            []string `json:"source" bson:"source"`
	RetrievedDate     string   `json:"retrieved_date" bson:"retrieved_date"`
	SourceDescription string   `json:"source_description" bson:"source_description"`
}

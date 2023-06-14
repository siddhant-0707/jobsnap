package model

import "gorm.io/gorm"

// Experience struct
type Experience struct {
	gorm.Model
	Company  int    `json:"company"`
	JobTitle string `json:"jobTitle"`
	WorkerID uint
	// TODO: add dates worked field
}

/*type payRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}*/

type Preference struct {
	gorm.Model
	JobTitle string `json:"JobTitle"`
	MinPay   int    `json:"minPay"`
	MaxPay   int    `json:"maxPay"`
	WorkerID uint
	// TODO: add location field
	// TODO: add locations field
}

// Worker struct
type Worker struct {
	gorm.Model
	Username     string       `gorm:"uniqueIndex;not null" json:"username"`
	Email        string       `gorm:"uniqueIndex;not null" json:"email"`
	Password     string       `gorm:"not null" json:"password"`
	Name         string       `json:"name"`
	Phone        string       `json:"phone"`
	Skills       []string     `gorm:"type:text[]" json:"skills"`
	Experience   []Experience `gorm:"foreignKey:WorkerID;constraint:OnDelete:CASCADE;default:[]" json:"Experience"`
	Availability bool         `gorm:"default:true" json:"availability"` // TODO: change to object which is based on date and time
	Preference   []Preference `gorm:"foreignKey:WorkerID;constraint:OnDelete:CASCADE;default:[]" json:"Preference"`
	AppliedJobs  []*Job       `gorm:"many2many:worker_jobs;" json:"appliedJobs"`

	// TODO: add education field
	// TODO: add Resume field (file upload)
}

package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title       string    `gorm:"uniqueIndex;not null" json:"title"`
	Description string    `json:"description"`
	MinPay      int       `json:"minPay"`
	MaxPay      int       `json:"maxPay"`
	Skills      []string  `gorm:"type:text[]" json:"skills"`
	Applicants  []*Worker `gorm:"many2many:worker_jobs;" json:"applicants"`
	//PostedBy    Poster   `json:"postedBy"`
	PosterID uint `json:"posterID"`
	// add start date
	// add requirements field
	// add location field
}

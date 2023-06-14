package model

import "gorm.io/gorm"

type Poster struct {
	gorm.Model
	Username   string `gorm:"uniqueIndex;not null" json:"username"`
	Email      string `gorm:"uniqueIndex;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	JobsPosted []Job  `gorm:"foreignkey:PosterID;constraint:OnDelete:CASCADE" json:"jobsPosted"`
}

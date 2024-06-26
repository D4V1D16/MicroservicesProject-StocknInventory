package models

import "gorm.io/gorm"

type Marca struct {
	gorm.Model

	Nombre      string `gorm:"not null" json:"brand_name"`
}
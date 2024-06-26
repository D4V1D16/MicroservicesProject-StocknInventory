package models

import "gorm.io/gorm"

type Proveedor struct {
	gorm.Model
	Nombre      string `gorm:"not null" json:"provider_name`
}
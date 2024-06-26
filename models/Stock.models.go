package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ProductoID uint  `gorm:"index"` 
	Cantidad   int32 `gorm:"not null"`
}
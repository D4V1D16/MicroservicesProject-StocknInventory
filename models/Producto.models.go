package models

import "gorm.io/gorm"

type Producto struct {
	gorm.Model
	IDProducto   uint     `gorm:"primaryKey; autoIncrement:true"`
	Nombre       string   `gorm:"not null"`
	Precio       float32  `gorm:"not null"`
	ProveedorID  int32    `gorm:"not null"`
	MarcaID      int32    `gorm:"not null"`
	Proveedor    Proveedor `gorm:"foreignKey:ProveedorID"`
	Marca        Marca     `gorm:"foreignKey:MarcaID"`
}
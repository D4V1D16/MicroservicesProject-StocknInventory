package utilities

import (
	"stockInventory/DB"
	"stockInventory/models"
)

func FindBrand(idBrand int) bool {
	var brand models.Marca

	DB.DB.First(&brand,idBrand)

	if brand.ID == 0{
		return false
	}
	return true
	
}
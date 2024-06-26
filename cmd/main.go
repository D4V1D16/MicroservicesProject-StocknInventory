package main

import (
	"fmt"
	"net/http"
	"stockInventory/DB"
	"stockInventory/models"
	"stockInventory/routes"
)

func main() {

	DB.DBConnection()
	
    DB.DB.AutoMigrate(models.Marca{})
    DB.DB.AutoMigrate(models.Proveedor{})
    DB.DB.AutoMigrate(models.Stock{})
	DB.DB.AutoMigrate(models.Producto{})

	//marca := models.Marca{}

	//productRoutes := routes.SetupProductRoutes()
	marcaRoutes := routes.SetupMarcasRoutes()

	fmt.Print("Sirviendo en el puerto 3000")


	http.ListenAndServe(":3000",marcaRoutes)
}
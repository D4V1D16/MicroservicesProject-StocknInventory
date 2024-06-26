package routes

import (
	"stockInventory/controllers"

	"github.com/gorilla/mux"
)



func SetupProductRoutes() *mux.Router{

	routesProduct := mux.NewRouter()

	//POST -> Creacion de Productos
	routesProduct.HandleFunc("/products",controllers.GetProductsHandler).Methods("GET")
	routesProduct.HandleFunc("/products",controllers.PostProductHandler).Methods("POST")
	routesProduct.HandleFunc("/products/{id}",controllers.DeleteProductHandler).Methods("DELETE")
	routesProduct.HandleFunc("/products/{id}",controllers.UpdateProductHandler).Methods("PUT")
	//GET -> Mostrar Productos

	//DELETE -> Borrar Productos

	


	return routesProduct

}





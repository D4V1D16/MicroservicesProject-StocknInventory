package routes

import (
	"stockInventory/controllers"

	"github.com/gorilla/mux"
)

func SetupMarcasRoutes() *mux.Router{
	marcasRoutes := mux.NewRouter()

	marcasRoutes.HandleFunc("/marcas",controllers.GetMarcas).Methods("GET")
	marcasRoutes.HandleFunc("/marcas",controllers.PostMarca).Methods("POST")
	marcasRoutes.HandleFunc("/marcas/{id}",controllers.DeleteMarca).Methods("DELETE")
	marcasRoutes.HandleFunc("/marcas/{id}",controllers.GetSingleMarca).Methods("GET")


	return marcasRoutes
}
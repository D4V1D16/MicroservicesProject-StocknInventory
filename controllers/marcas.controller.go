package controllers

import (
	"encoding/json"
	"net/http"
	"stockInventory/models"
	"stockInventory/DB"
)

func PostMarca(w http.ResponseWriter, r *http.Request) {
	var marca models.Marca

	json.NewDecoder(r.Body).Decode(&marca)
	createdMarca := DB.DB.Create(&marca)

	err := createdMarca.Error
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&marca)
	

}

func GetMarcas(w http.ResponseWriter, r *http.Request){
	var marcas[]models.Marca

	DB.DB.Find(&marcas)
	json.NewEncoder(w).Encode(&marcas)
}
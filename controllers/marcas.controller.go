package controllers

import (
	"encoding/json"
	"net/http"
	"stockInventory/DB"
	"stockInventory/models"

	"github.com/gorilla/mux"
)

func PostMarca(w http.ResponseWriter, r *http.Request) {
	var marca models.Marca

	json.NewDecoder(r.Body).Decode(&marca)

	if marca.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("brand_name cannot be empty"))
		return
	}
	createdMarca := DB.DB.Create(&marca)
	if createdMarca.Error != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating brand"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&marca)
}

func GetMarcas(w http.ResponseWriter, r *http.Request){
	var marcas []models.Marca

	DB.DB.Find(&marcas)

	json.NewEncoder(w).Encode(&marcas)
}


func DeleteMarca(w http.ResponseWriter, r *http.Request){
	var marca models.Marca

	params := mux.Vars(r)
	deleteBrand := DB.DB.Delete(&marca,params["id"])

	if deleteBrand.Error != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting brand"))
		return
	}

	if deleteBrand.RowsAffected == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Brand not found"))
		return
	}


	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Brand Deleted"))
	return
}



func GetSingleMarca(w http.ResponseWriter, r *http.Request){
	var marca models.Marca
	params := mux.Vars(r)
	DB.DB.First(&marca,params["id"])

	if marca.ID  == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Brand not found"))
		return
	}
	json.NewEncoder(w).Encode(&marca)
}
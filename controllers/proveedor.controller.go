package controllers

import (
	"encoding/json"
	"net/http"
	"stockInventory/DB"
	"stockInventory/models"
)

func PostProveeodr(w http.ResponseWriter, r *http.Request) {
	var proveedor models.Proveedor

	json.NewDecoder(r.Body).Decode(&proveedor)

	if proveedor.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provider_name cannot be empty"))
		return
	}
	createdProvider := DB.DB.Create(&proveedor)
	if createdProvider.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating provider"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&proveedor)
}

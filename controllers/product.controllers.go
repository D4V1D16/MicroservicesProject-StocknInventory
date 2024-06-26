package controllers

import (
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request){
	//Crear un Producto
	w.Write([]byte("Getting Products"))
}

func PostProductHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Posting Products"))
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deleting Products"))
}

func UpdateProductHandler (w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Updating Products"))
}
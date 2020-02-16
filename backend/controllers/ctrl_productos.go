package controllers

import (

	"encoding/json"
	"net/http"

)

func GetProductos(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{productos}")

}

func GetProducto(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{producto}")

}

func SaveProducto(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelProducto(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
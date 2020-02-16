package controllers

import (

	"encoding/json"
	"net/http"

)

func GetCategorias(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{categorias}")

}

func GetCategoria(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{categoria}")

}

func SaveCategoria(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelCategoria(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
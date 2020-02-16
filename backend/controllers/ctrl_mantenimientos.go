package controllers

import (

	"encoding/json"
	"net/http"

)

func GetMantenimientos(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{Mantenimientos}")

}

func GetMantenimiento(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{mantenimiento}")

}

func SaveMantenimiento(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelMantenimiento(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
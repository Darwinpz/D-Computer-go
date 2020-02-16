package controllers

import (

	"encoding/json"
	"net/http"

)

func GetHistoriales(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{historiales}")

}

func GetHistorial(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{historial}")

}

func SaveHistorial(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelHistorial(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
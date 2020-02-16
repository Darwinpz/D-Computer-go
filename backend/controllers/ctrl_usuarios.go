package controllers

import (

	"encoding/json"
	"net/http"

)

func GetUsuarios(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{Usuario}")

}

func GetUsuario(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{usuario}")

}

func SaveUsuario(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelUsuario(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
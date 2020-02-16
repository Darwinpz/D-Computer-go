package controllers

import (

	"encoding/json"
	"net/http"

)

func GetTipo_Usuarios(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{TipoUsuario}")

}

func GetTipo_Usuario(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{tipousuario}")

}

func SaveTipo_Usuario(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelTipo_Usuario(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
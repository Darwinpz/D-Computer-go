package controllers

import (

	"encoding/json"
	"net/http"

)

func GetAsistencias(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{asistencias}")

}

func GetAsistencia(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{asistencia}")

}

func SaveAsistencia(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelAsistencia(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
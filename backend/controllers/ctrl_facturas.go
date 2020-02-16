package controllers

import (

	"encoding/json"
	"net/http"

)

func GetFacturas(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{facturas}")

}

func GetFactura(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{factura}")

}

func SaveFactura(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelFactura(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
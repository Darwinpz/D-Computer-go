package controllers

import (

	"encoding/json"
	"net/http"

)

func GetProd_Facturas(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{prodfacturas}")

}

func GetProd_Factura(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{prodfactura}")

}

func SaveProd_Factura(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{guardar}")

}

func DelProd_Factura(res http.ResponseWriter, req *http.Request){

	json.NewEncoder(res).Encode("{eliminar}")

}
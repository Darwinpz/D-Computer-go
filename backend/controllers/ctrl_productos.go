package controllers

import (

	"encoding/json"
	"net/http"
	models "d-computer/backend/models"
	
)


func GetProductos(res http.ResponseWriter, req *http.Request){

	var lista_productos [] models.Productos

	var obj_producto models.Productos

	rows, err := GetDB().Query("SELECT * FROM productos")
	
	if err != nil {
		
		json.NewEncoder(res).Encode(lista_productos)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_producto.Prod_cod,
				&obj_producto.Prod_nombre,
				&obj_producto.Cat_cod,	
			)
	
			lista_productos = append(lista_productos, obj_producto)
		}

		defer rows.Close()
		
		json.NewEncoder(res).Encode(lista_productos)

	}	

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
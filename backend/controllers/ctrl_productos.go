package controllers

import (

	"encoding/json"
	"net/http"
	models "d-computer/backend/models"
	
)

func GetProductos(res http.ResponseWriter, req *http.Request){

	var productos [] models.productos

	obj_producto := new(models.productos)
	
	rows, err := GetDB().Query("SELECT * FROM productos")

	if err != nil {
		return json.NewEncoder(res).Encode(&models.productos{})
	}
	
	defer rows.Close()

	for rows.Next() {

		rows.Scan(
			&obj_producto.prod_cod,
			&obj_producto.prod_nombre,
			&obj_producto.cat_cod
		)

		productos = append(productos, obj_producto)
	}

	json.NewEncoder(res).Encode(productos)

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
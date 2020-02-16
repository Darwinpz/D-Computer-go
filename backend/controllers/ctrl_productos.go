package controllers

import (

	"encoding/json"
	"net/http"
	models "d-computer/backend/models"
	"github.com/gorilla/mux"
	//"fmt"
)


func GetProductos(res http.ResponseWriter, req *http.Request){

	var lista_productos [] models.Productos

	var obj_producto models.Productos

	rows, err := GetDB().Query("SELECT * FROM PRODUCTOS")
	
	if err != nil {
		
		json.NewEncoder(res).Encode(err)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_producto.Prod_cod,
				&obj_producto.Prod_nombre,
				&obj_producto.Prod_precio_costo,	
				&obj_producto.Prod_precio_venta,	
				&obj_producto.Cat_cod,	
				
			)
	
			lista_productos = append(lista_productos, obj_producto)
		}

		defer rows.Close()
		
		json.NewEncoder(res).Encode(lista_productos)

	}	

}

func GetProducto(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_producto models.Productos

	rows, err := GetDB().Query("SELECT * FROM PRODUCTOS where PROD_COD = "+params["id"])

	if err != nil {
		
		json.NewEncoder(res).Encode(err)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_producto.Prod_cod,
				&obj_producto.Prod_nombre,
				&obj_producto.Prod_precio_costo,	
				&obj_producto.Prod_precio_venta,	
				&obj_producto.Cat_cod,	
			)

		}
		
		json.NewEncoder(res).Encode(obj_producto)

	}


}

func SaveProducto(res http.ResponseWriter, req *http.Request){
	
	var obj_producto models.Productos

	smt, err := GetDB().Prepare("INSERT INTO PRODUCTOS (PROD_NOMBRE,PROD_PRECIO_COSTO,PROD_PRECIO_VENTA,CAT_COD) VALUES(:1,:2,:3,:4)")

	if err != nil {
		
		json.NewEncoder(res).Encode(err)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_producto)

		_, err := smt.Exec(obj_producto.Prod_nombre,obj_producto.Prod_precio_costo,obj_producto.Prod_precio_venta,obj_producto.Cat_cod)

		if err != nil {

			json.NewEncoder(res).Encode(err)

		}else{

			json.NewEncoder(res).Encode("Producto Guardado")

		}

		
	}


}

func DelProducto(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	smt, err := GetDB().Prepare("DELETE FROM PRODUCTOS WHERE PROD_COD =: 1")

	if err != nil {
		
		json.NewEncoder(res).Encode(err)
		
	}else{

		_, err := smt.Exec(params["id"])
		
		if err != nil {

			json.NewEncoder(res).Encode(err)

		}else{

			json.NewEncoder(res).Encode("Producto Eliminado")

		}

	
	}

	
}


func UpdateProducto(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_producto models.Productos

	smt, err := GetDB().Prepare("UPDATE PRODUCTOS SET PROD_NOMBRE =: 1,PROD_PRECIO_COSTO =: 2,PROD_PRECIO_VENTA =: 3,CAT_COD =: 4 WHERE PROD_COD =: 5")

	if err != nil {
		
		json.NewEncoder(res).Encode(err)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_producto)

		_, err := smt.Exec(obj_producto.Prod_nombre,obj_producto.Prod_precio_costo,obj_producto.Prod_precio_venta,obj_producto.Cat_cod,params["id"])

		if err != nil {

			json.NewEncoder(res).Encode(err)

		}else{

			json.NewEncoder(res).Encode("Producto Actualizado")

		}

	}
}
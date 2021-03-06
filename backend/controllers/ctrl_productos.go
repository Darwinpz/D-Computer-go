package controllers

import (

	"encoding/json"
	"net/http"
	models "d-computer/backend/models"
	"github.com/gorilla/mux"
	
)


func GetProductos(res http.ResponseWriter, req *http.Request){

	var lista_productos [] models.Productos

	var obj_producto models.Productos

	rows, err := GetDB().Query("SELECT * FROM PRODUCTOS")
	
	if err != nil {

		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_producto.Prod_cod,
				&obj_producto.Prod_nombre,
				&obj_producto.Prod_cantidad,
				&obj_producto.Prod_descripcion,
				&obj_producto.Prod_precio_costo,	
				&obj_producto.Prod_precio_venta,	
				&obj_producto.Cat_cod,	
				
			)
	
			lista_productos = append(lista_productos, obj_producto)
		}

		defer rows.Close()
		
		if len(lista_productos) != 0 {

			json.NewEncoder(res).Encode(lista_productos)

		}else{

			json.NewEncoder(res).Encode(make([]string, 0))

		}


	}	

}

func GetProducto(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_producto models.Productos

	rows, err := GetDB().Query("SELECT * FROM PRODUCTOS where PROD_COD = "+params["id"])

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_producto.Prod_cod,
				&obj_producto.Prod_nombre,
				&obj_producto.Prod_cantidad,
				&obj_producto.Prod_descripcion,
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

	smt, err := GetDB().Prepare("INSERT INTO PRODUCTOS (PROD_NOMBRE,PROD_CANTIDAD,PROD_DESCRIPCION,PROD_PRECIO_COSTO,PROD_PRECIO_VENTA,CAT_COD) VALUES(:1,:2,:3,:4,:5,:6)")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_producto)

		_, err := smt.Exec(obj_producto.Prod_nombre,obj_producto.Prod_cantidad,obj_producto.Prod_descripcion,obj_producto.Prod_precio_costo,obj_producto.Prod_precio_venta,obj_producto.Cat_cod)

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje:"Producto Guardado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

		
	}


}

func DelProducto(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	smt, err := GetDB().Prepare("DELETE FROM PRODUCTOS WHERE PROD_COD =: 1")

	if err != nil {

		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_, err := smt.Exec(params["id"])
		
		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje:"Producto Eliminado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	
	}

	
}


func UpdateProducto(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_producto models.Productos

	smt, err := GetDB().Prepare("UPDATE PRODUCTOS SET PROD_NOMBRE =: 1,PROD_CANTIDAD =:2,PROD_DESCRIPCION =:3,PROD_PRECIO_COSTO =: 4,PROD_PRECIO_VENTA =: 5,CAT_COD =: 6 WHERE PROD_COD =: 7")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_producto)

		_, err := smt.Exec(obj_producto.Prod_nombre,obj_producto.Prod_cantidad,obj_producto.Prod_descripcion,obj_producto.Prod_precio_costo,obj_producto.Prod_precio_venta,obj_producto.Cat_cod,params["id"])

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje:"Producto Actualizado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	}
}
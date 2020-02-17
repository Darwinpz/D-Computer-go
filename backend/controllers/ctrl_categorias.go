package controllers

import (

	"encoding/json"
	"net/http"
	models "d-computer/backend/models"
	"github.com/gorilla/mux"
	

)

func GetCategorias(res http.ResponseWriter, req *http.Request){

	var lista_categorias [] models.Categorias

	var obj_categoria models.Categorias

	rows, err := GetDB().Query("SELECT * FROM CATEGORIAS")
	
	if err != nil {

		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_categoria.Cat_cod,
				&obj_categoria.Cat_nombre,
			)
	
			lista_categorias = append(lista_categorias, obj_categoria)
		}

		defer rows.Close()
		
		if len(lista_categorias) != 0 {

			json.NewEncoder(res).Encode(lista_categorias)

		}else{

			json.NewEncoder(res).Encode(make([]string, 0))

		}
		
	}	

}

func GetCategoria(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_categoria models.Categorias

	rows, err := GetDB().Query("SELECT * FROM CATEGORIAS where CAT_COD = "+params["id"])

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_categoria.Cat_cod,
				&obj_categoria.Cat_nombre,
			)

		}
		
		json.NewEncoder(res).Encode(obj_categoria)

	}

}

func SaveCategoria(res http.ResponseWriter, req *http.Request){

	var obj_categoria models.Categorias

	smt, err := GetDB().Prepare("INSERT INTO CATEGORIAS (CAT_NOMBRE) VALUES(:1)")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_categoria)

		_, err := smt.Exec(obj_categoria.Cat_nombre)

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Categoria Guardada"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}
		
	}

}

func DelCategoria(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	smt, err := GetDB().Prepare("DELETE FROM CATEGORIAS WHERE CAT_COD =: 1")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_, err := smt.Exec(params["id"])
		
		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Categoria Eliminada"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	
	}

}

func UpdateCategoria(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_categoria models.Categorias

	smt, err := GetDB().Prepare("UPDATE CATEGORIAS SET CAT_NOMBRE =: 1 WHERE CAT_COD =: 2")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_categoria)

		_, err := smt.Exec(obj_categoria.Cat_nombre,params["id"])

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Categoria Actualizada"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	}

}
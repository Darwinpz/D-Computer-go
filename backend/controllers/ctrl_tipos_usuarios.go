package controllers

import (

	"encoding/json"
	"net/http"
	models "d-computer/backend/models"
	"github.com/gorilla/mux"

)

func GetTipo_Usuarios(res http.ResponseWriter, req *http.Request){

	var lista_tipo_usuarios [] models.Tipo_usuarios

	var obj_tipo_usuarios models.Tipo_usuarios

	rows, err := GetDB().Query("SELECT * FROM TIPO_USUARIOS")
	
	if err != nil {

		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_tipo_usuarios.Tipo_user_cod,
				&obj_tipo_usuarios.Tipo_user_nombre,
			)
	
			lista_tipo_usuarios = append(lista_tipo_usuarios, obj_tipo_usuarios)
		}

		defer rows.Close()
		
		if len(lista_tipo_usuarios) != 0 {

			json.NewEncoder(res).Encode(lista_tipo_usuarios)

		}else{

			json.NewEncoder(res).Encode(make([]string, 0))

		}
		
	}	

}

func GetTipo_Usuario(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_tipo_usuarios models.Tipo_usuarios

	rows, err := GetDB().Query("SELECT * FROM TIPO_USUARIOS where TIPO_USER_COD = "+params["id"])

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_tipo_usuarios.Tipo_user_cod,
				&obj_tipo_usuarios.Tipo_user_nombre,
			)

		}
		
		json.NewEncoder(res).Encode(obj_tipo_usuarios)

	}


}

func SaveTipo_Usuario(res http.ResponseWriter, req *http.Request){

	var obj_tipo_usuarios models.Tipo_usuarios

	smt, err := GetDB().Prepare("INSERT INTO TIPO_USUARIOS (TIPO_USER_NOMBRE) VALUES(:1)")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_tipo_usuarios)

		_, err := smt.Exec(obj_tipo_usuarios.Tipo_user_nombre)

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Tipo de usuario Guardado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}
		
	}

}

func DelTipo_Usuario(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	smt, err := GetDB().Prepare("DELETE FROM TIPO_USUARIOS WHERE TIPO_USER_COD =: 1")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_, err := smt.Exec(params["id"])
		
		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Tipo de usuario Eliminado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	
	}

}


func UpdateTipo_Usuario(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_tipo_usuarios models.Tipo_usuarios

	smt, err := GetDB().Prepare("UPDATE TIPO_USUARIOS SET TIPO_USER_NOMBRE =: 1 WHERE TIPO_USER_COD =: 2")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_tipo_usuarios)

		_, err := smt.Exec(obj_tipo_usuarios.Tipo_user_nombre,params["id"])

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Tipo de usuario Actualizado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	}

}
package controllers

import (

	"encoding/json"
	"net/http"
	"time"
	models "d-computer/backend/models"
	"github.com/gorilla/mux"

)

func GetUsuarios(res http.ResponseWriter, req *http.Request){

	var lista_usuarios [] models.Usuarios

	var obj_usuario models.Usuarios

	rows, err := GetDB().Query("SELECT * FROM USUARIOS")
	
	if err != nil {

		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_usuario.User_ced,
				&obj_usuario.User_nombre,
				&obj_usuario.User_correo,	
				&obj_usuario.User_telefono,	
				&obj_usuario.User_clave,
				&obj_usuario.User_tipo,	
				&obj_usuario.User_registro,
			)
	
			lista_usuarios = append(lista_usuarios, obj_usuario)
		}

		defer rows.Close()
		
		if len(lista_usuarios) != 0 {

			json.NewEncoder(res).Encode(lista_usuarios)

		}else{

			json.NewEncoder(res).Encode(make([]string, 0))

		}

	}	

}

func GetUsuario(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_usuario models.Usuarios

	rows, err := GetDB().Query("SELECT * FROM USUARIOS where USER_CED = "+params["id"])

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_usuario.User_ced,
				&obj_usuario.User_nombre,
				&obj_usuario.User_correo,	
				&obj_usuario.User_telefono,	
				&obj_usuario.User_clave,
				&obj_usuario.User_tipo,	
				&obj_usuario.User_registro,
			)

		}
		
		json.NewEncoder(res).Encode(obj_usuario)

	}

}

func SaveUsuario(res http.ResponseWriter, req *http.Request){

	var obj_usuario models.Usuarios

	smt, err := GetDB().Prepare("INSERT INTO USUARIOS (USER_CED,USER_NOMBRE,USER_CORREO,USER_TELEFONO,USER_CLAVE,USER_TIPO,USER_REGISTRO) VALUES(:1,:2,:3,:4,:5,:6,:7)")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_usuario)

		obj_usuario.User_registro = time.Now().UTC()

		_, err := smt.Exec(obj_usuario.User_ced,obj_usuario.User_nombre,obj_usuario.User_correo,obj_usuario.User_telefono,obj_usuario.User_clave,obj_usuario.User_tipo,obj_usuario.User_registro)

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje:"Usuario Guardado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

		
	}

}

func DelUsuario(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	smt, err := GetDB().Prepare("DELETE FROM USUARIOS WHERE USER_CED =: 1")

	if err != nil {

		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_, err := smt.Exec(params["id"])
		
		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje:"Usuario Eliminado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	
	}

}


func UpdateUsuario(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_usuario models.Usuarios

	smt, err := GetDB().Prepare("UPDATE USUARIOS SET USER_CED =: 1,USER_NOMBRE =: 2,USER_CORREO =: 3,USER_TELEFONO =: 4,USER_CLAVE =: 5,USER_TIPO =: 6 WHERE USER_CED =: 7")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_usuario)

		_, err := smt.Exec(obj_usuario.User_ced,obj_usuario.User_nombre,obj_usuario.User_correo,obj_usuario.User_telefono,obj_usuario.User_clave,obj_usuario.User_tipo,params["id"])

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje:"Usuario Actualizado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	}

}
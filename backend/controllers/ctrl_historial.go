package controllers

import (

	"encoding/json"
	"net/http"
	models "d-computer/backend/models"
	"github.com/gorilla/mux"
	"time"

)

func GetHistoriales(res http.ResponseWriter, req *http.Request){

	var lista_historial [] models.Historial

	var obj_historial models.Historial

	rows, err := GetDB().Query("SELECT * FROM HISTORIALES")
	
	if err != nil {

		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_historial.History_cod,
				&obj_historial.User_ced,
				&obj_historial.History_actividad,
				&obj_historial.History_fecha,
			)
	
			lista_historial = append(lista_historial, obj_historial)
		}

		defer rows.Close()
		
		if len(lista_historial) != 0 {

			json.NewEncoder(res).Encode(lista_historial)

		}else{

			json.NewEncoder(res).Encode(make([]string, 0))

		}
		
	}	

}

func GetHistorial(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_historial models.Historial

	rows, err := GetDB().Query("SELECT * FROM HISTORIALES where HISTORY_COD = "+params["id"])

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		for rows.Next() {

			rows.Scan(
				&obj_historial.History_cod,
				&obj_historial.User_ced,
				&obj_historial.History_actividad,
				&obj_historial.History_fecha,
			)

		}
		
		json.NewEncoder(res).Encode(obj_historial)

	}

}

func SaveHistorial(res http.ResponseWriter, req *http.Request){

	var obj_historial models.Historial

	smt, err := GetDB().Prepare("INSERT INTO HISTORIALES (USER_CED,HISTORY_ACTIVIDAD,HISTORY_FECHA) VALUES(:1,:2,:3)")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_historial)

		obj_historial.History_fecha = time.Now().UTC()

		_, err := smt.Exec(obj_historial.User_ced,obj_historial.History_actividad,obj_historial.History_fecha)

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Historial Guardado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}
		
	}


}

func DelHistorial(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	smt, err := GetDB().Prepare("DELETE FROM HISTORIALES WHERE HISTORY_COD =: 1")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_, err := smt.Exec(params["id"])
		
		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Historial Eliminado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	
	}

}

func UpdateHistorial(res http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	var obj_historial models.Historial

	smt, err := GetDB().Prepare("UPDATE HISTORIALES SET USER_CED =: 1,HISTORY_ACTIVIDAD =:2,HISTORY_FECHA =:3  WHERE HISTORY_COD =: 4")

	if err != nil {
		
		obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

		json.NewEncoder(res).Encode(obj_mensaje)
		
	}else{

		_ = json.NewDecoder(req.Body).Decode(&obj_historial)

		obj_historial.History_fecha = time.Now().UTC()

		_, err := smt.Exec(obj_historial.User_ced,obj_historial.History_actividad,obj_historial.History_fecha,params["id"])

		if err != nil {

			obj_mensaje := models.Mensajes {Tipo:"error", Mensaje: err.Error()}

			json.NewEncoder(res).Encode(obj_mensaje)

		}else{

			obj_mensaje := models.Mensajes {Tipo:"success", Mensaje: "Historial Actualizado"}

			json.NewEncoder(res).Encode(obj_mensaje)

		}

	}

}
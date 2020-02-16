package models


import (
	"time"
)


type Usuarios struct 
{
	User_ced int `json:"user_ced,omitempty"`
	User_nombre string `json:"user_nombre,omitempty"`
	User_correo string `json:"user_correo"`
	User_telefono string `json:"user_telefono"`
	User_clave string `json:"user_clave,omitempty"`
	User_tipo string `json:"user_tipo,omitempty"`
	User_registro time.Time `json:"user_registro,omitempty"`

}

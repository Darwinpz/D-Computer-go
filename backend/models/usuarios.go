package models


import (
	"time"
)


type Usuarios struct 
{
	User_ced string `json:"user_ced,omitempty"`
	User_nombre string `json:"user_nombre,omitempty"`
	User_correo string `json:"user_correo"`
	User_telefono string `json:"user_telefono"`
	User_clave string `json:"user_clave,omitempty"`
	Tipo_user_cod int `json:"Tipo_user_cod,omitempty"`
	User_registro time.Time `json:"user_registro,omitempty"`

}

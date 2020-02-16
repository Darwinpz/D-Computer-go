package models


import (
	"time"
)


type usuarios struct 
{
	user_ced int `json:"user_ced,omitempty"`
	user_nombre string `json:"user_nombre,omitempty"`
	user_correo string `json:"user_correo"`
	user_telefono string `json:"user_telefono"`
	user_clave string `json:"user_clave,omitempty"`
	user_tipo string `json:"user_tipo,omitempty"`
	user_registro time.Time `json:"user_registro,omitempty"`

}

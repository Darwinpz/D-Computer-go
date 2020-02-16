package models


import (
	"time"
)


type asistencia struct 
{
	asist_cod int `json:"asist_cod,omitempty"`
	user_ced string `json:"user_ced,omitempty"`
	asist_fecha_in time.Time `json:"asist_fecha_in"`
	asist_fecha_out time.Time `json:"asist_fecha_out"`
	

}
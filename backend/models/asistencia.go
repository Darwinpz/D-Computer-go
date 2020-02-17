package models


import (
	"time"
)


type Asistencia struct 
{
	Asist_cod int `json:"asist_cod,omitempty"`
	User_ced string `json:"user_ced,omitempty"`
	Asist_fecha_in time.Time `json:"asist_fecha_in"`
	Asist_fecha_out time.Time `json:"asist_fecha_out"`
	
}
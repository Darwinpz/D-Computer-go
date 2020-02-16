package models


import (
	"time"
)


type Historial struct 
{
	History_cod int `json:"history_cod,omitempty"`
	User_ced string `json:"user_ced,omitempty"`
	History_actividad string `json:"history_actividad,omitempty"`
	History_fecha time.Time `json:"history_fecha,omitempty"`

}

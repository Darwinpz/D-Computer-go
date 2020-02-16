package models


import (
	"time"
)


type historial struct 
{
	history_cod int `json:"history_cod,omitempty"`
	user_ced string `json:"user_ced,omitempty"`
	history_actividad string `json:"history_actividad,omitempty"`
	history_fecha time.Time `json:"history_fecha,omitempty"`

}

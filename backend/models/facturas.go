package models

import (
	"time"
)

type facturas struct 
{
	fact_cod int `json:"fact_cod,omitempty"`
	user_ced string `json:"user_ced,omitempty"`
	fact_fecha time.Time `json:"fact_fecha,omitempty"`
	fact_precio float32 `json:"fact_precio,omitempty"`
}

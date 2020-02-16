package models

import (
	"time"
)

type Facturas struct 
{
	Fact_cod int `json:"fact_cod,omitempty"`
	User_ced string `json:"user_ced,omitempty"`
	Fact_fecha time.Time `json:"fact_fecha,omitempty"`
	Fact_precio float32 `json:"fact_precio,omitempty"`
}

package models

import (
	"time"
)

type facturas struct 
{
	cod_factura int `json:"cod_factura,omitempty"`
	cod_usuario int `json:"cod_usuario,omitempty"`
	fecha_factura time.Time `json:"fecha_factura,omitempty"`
	precio_factura float32 `json:"precio_factura,omitempty"`
}

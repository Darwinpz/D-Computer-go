package models


type Productos struct {

	Prod_cod int `json:"prod_cod,omitempty"`
	Prod_nombre string `json:"prod_nombre,omitempty"`
	Prod_precio_costo float32 `json:"prod_precio_costo,omitempty"`
	Prod_precio_venta float32 `json:"prod_precio_venta,omitempty"`
	Cat_cod int `json:"cat_cod,omitempty"`

}

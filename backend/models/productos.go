package models


type Productos struct {

	Prod_cod int `json:"prod_cod,omitempty"`
	Prod_nombre string `json:"prod_nombre,omitempty"`
	Cat_cod int `json:"cat_cod,omitempty"`
}

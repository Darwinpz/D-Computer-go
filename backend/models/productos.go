package models


type productos struct 
{
	prod_cod int `json:"prod_cod,omitempty"`
	prod_nombre string `json:"prod_nombre,omitempty"`
	cat_cod int `json:"cat_cod,omitempty"`
}

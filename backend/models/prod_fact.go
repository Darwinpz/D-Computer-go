package models


type prod_fact struct 
{

	prod_fact_cod int `json:"prod_fact_cod,omitempty"`
	fact_cod int `json:"fact_cod,omitempty"`
	prod_cod int `json:"prod_cod,omitempty"`
	prod_fact_cantidad int `json:"prod_fact_cantidad,omitempty"`
}

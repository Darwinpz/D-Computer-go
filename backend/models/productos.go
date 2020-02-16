package models


type productos struct 
{
	cod_producto int `json:"cod_prod,omitempty"`
	nombre_producto string `json:"nombre_prod,omitempty"`
	cod_categoria int `json:"cod_categoria,omitempty"`
}

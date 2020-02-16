package main

import (
	ctrl "d-computer/controllers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)


func main(){

	router := mux.NewRouter()

	//productos
	router.HandleFunc("/dcomputer/api/productos",ctrl.GetProductos).Methods("GET")
	router.HandleFunc("/dcomputer/api/productos/{id}",ctrl.GetProducto).Methods("GET")
	router.HandleFunc("/dcomputer/api/productos",ctrl.SaveProducto).Methods("POST")
	router.HandleFunc("/dcomputer/api/productos/{id}",ctrl.DelProducto).Methods("DELETE")

	//facturas
	router.HandleFunc("/dcomputer/api/facturas",ctrl.GetFacturas).Methods("GET")
	router.HandleFunc("/dcomputer/api/facturas/{id}",ctrl.GetFactura).Methods("GET")
	router.HandleFunc("/dcomputer/api/facturas",ctrl.SaveFactura).Methods("POST")
	router.HandleFunc("/dcomputer/api/facturas/{id}",ctrl.DelFactura).Methods("DELETE")

	//categorias
	router.HandleFunc("/dcomputer/api/categorias",ctrl.GetCategorias).Methods("GET")
	router.HandleFunc("/dcomputer/api/categorias/{id}",ctrl.GetCategoria).Methods("GET")
	router.HandleFunc("/dcomputer/api/categorias",ctrl.SaveCategoria).Methods("POST")
	router.HandleFunc("/dcomputer/api/categorias/{id}",ctrl.DelCategoria).Methods("DELETE")

	//productos_facturas
	router.HandleFunc("/dcomputer/api/productos_facturas",ctrl.GetProd_Facturas).Methods("GET")
	router.HandleFunc("/dcomputer/api/productos_facturas/{id}",ctrl.GetProd_Factura).Methods("GET")
	router.HandleFunc("/dcomputer/api/productos_facturas",ctrl.SaveProd_Factura).Methods("POST")
	router.HandleFunc("/dcomputer/api/productos_facturas/{id}",ctrl.DelProd_Factura).Methods("DELETE")

	//usuarios
	router.HandleFunc("/dcomputer/api/usuarios",ctrl.GetUsuarios).Methods("GET")
	router.HandleFunc("/dcomputer/api/usuarios/{id}",ctrl.GetUsuario).Methods("GET")
	router.HandleFunc("/dcomputer/api/usuarios",ctrl.SaveUsuario).Methods("POST")
	router.HandleFunc("/dcomputer/api/usuarios/{id}",ctrl.DelUsuario).Methods("DELETE")

	//historial
	router.HandleFunc("/dcomputer/api/historial",ctrl.GetHistoriales).Methods("GET")
	router.HandleFunc("/dcomputer/api/historial/{id}",ctrl.GetHistorial).Methods("GET")
	router.HandleFunc("/dcomputer/api/historial",ctrl.SaveHistorial).Methods("POST")
	router.HandleFunc("/dcomputer/api/historial/{id}",ctrl.DelHistorial).Methods("DELETE")

	//tipos_usuarios
	router.HandleFunc("/dcomputer/api/tipos_usuarios",ctrl.GetTipo_Usuarios).Methods("GET")
	router.HandleFunc("/dcomputer/api/tipos_usuarios/{id}",ctrl.GetTipo_Usuario).Methods("GET")
	router.HandleFunc("/dcomputer/api/tipos_usuarios",ctrl.SaveTipo_Usuario).Methods("POST")
	router.HandleFunc("/dcomputer/api/tipos_usuarios/{id}",ctrl.DelTipo_Usuario).Methods("DELETE")

	//asistencia
	router.HandleFunc("/dcomputer/api/asistencia",ctrl.GetAsistencias).Methods("GET")
	router.HandleFunc("/dcomputer/api/asistencia/{id}",ctrl.GetAsistencia).Methods("GET")
	router.HandleFunc("/dcomputer/api/asistencia",ctrl.SaveAsistencia).Methods("POST")
	router.HandleFunc("/dcomputer/api/asistencia/{id}",ctrl.DelAsistencia).Methods("DELETE")

	//mantenimientos
	router.HandleFunc("/dcomputer/api/mantenimientos",ctrl.GetMantenimientos).Methods("GET")
	router.HandleFunc("/dcomputer/api/mantenimientos/{id}",ctrl.GetMantenimiento).Methods("GET")
	router.HandleFunc("/dcomputer/api/mantenimientos",ctrl.SaveMantenimiento).Methods("POST")
	router.HandleFunc("/dcomputer/api/mantenimientos/{id}",ctrl.DelMantenimiento).Methods("DELETE")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"Access-Control-Allow-Origin", "*"})
	
	log.Fatal(http.ListenAndServe(":4000",handlers.CORS(headers, methods, origin)(router)))
	
}
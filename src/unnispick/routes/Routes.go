package routes

import (
	"github.com/gorilla/mux"
	"unnispick/endpoint"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/brand", endpoint.BrandEndpointWithoutId).Methods("POST", "OPTIONS")
	r.HandleFunc("/brand/{Id}", endpoint.BrandEntpointWithId).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/produk", endpoint.ProductEndpointWithoutId).Methods("POST", "GET", "OPTIONS")
	r.HandleFunc("/produk/{Id}", endpoint.ProductEndpointWithId).Methods("DELETE", "PUT", "OPTIONS")
	return r
}
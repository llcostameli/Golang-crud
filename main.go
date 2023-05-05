package main

import (
	"fmt"
	"go-api-meli/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", controller.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products", controller.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", controller.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", controller.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/products/{productID}", controller.DeleteProduct).Methods(http.MethodDelete)

	router.HandleFunc("/cart", controller.AddProductToCart).Methods(http.MethodPost)
	router.HandleFunc("/cart/{id}", controller.GetCartById).Methods(http.MethodGet)

	fmt.Println("servidor rodando")
	log.Fatal(http.ListenAndServe(":5000", router))
}

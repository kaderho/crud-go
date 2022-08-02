package main

import (
	"crud/controllers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
	"html/template"
	"log"
	"net/http"
)

var index_template = template.Must(template.ParseFiles("web/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	err := index_template.ExecuteTemplate(w, "index.html", "Pratice crud go")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Println("Server started on: http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/", Index).Methods("GET")

	router.HandleFunc("/product-category", controllers.IndexProductCategory).Methods("GET")
	router.HandleFunc("/product-category/get", controllers.GetAllProductCategory).Methods("GET")
	router.HandleFunc("/product-category/create", controllers.CreateProductCategory).Methods("POST")
	router.HandleFunc("/product-category/update/{id}", controllers.UpdateProductCategory).Methods("PUT")
	router.HandleFunc("/product-category/get/{id}", controllers.GetOneProductCategory).Methods("GET")
	router.HandleFunc("/product-category/delete/{id}", controllers.DeleteProductCategory).Methods("DELETE")

	router.HandleFunc("/product", controllers.IndexProduct).Methods("GET")
	router.HandleFunc("/product/get", controllers.GetAllProduct).Methods("GET")
	router.HandleFunc("/product/create", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/update/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/get/{id}", controllers.GetOneProduct).Methods("GET")
	router.HandleFunc("/product/delete/{id}", controllers.DeleteProduct).Methods("DELETE")
}

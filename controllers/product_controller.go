package controllers

import (
	"crud/database"
	"crud/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

//Display page
func IndexProduct(w http.ResponseWriter, r *http.Request) {
	var index_template = template.Must(template.ParseFiles("web/product.html"))
	err := index_template.ExecuteTemplate(w, "product.html", "Liste des produit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//GetAllProduct get all Product data
func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var Products []entity.Product
	var Product entity.Product
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&Product.Id, &Product.ProductCategoryId, &Product.Name, &Product.Description, &Product.Price, &Product.Quantity)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			Products = append(Products, Product)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Products)
}

//CreateProduct creates Product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	product_category := r.FormValue("product_category")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")
	insertData, err := db.Prepare("INSERT INTO products(product_category_id, name, description, price, quantity) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
		return
	}

	insertData.Exec(product_category, name, description, price, quantity)
	println("Insert data to database")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Create succeeded")
}

//GetOneProduct get's Product with specific ID
func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()
	id := mux.Vars(r)["id"]

	selDB, err := db.Query("SELECT * FROM products WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	Product := entity.Product{}
	for selDB.Next() {
		err = selDB.Scan(&Product.Id, &Product.ProductCategoryId, &Product.Name, &Product.Description, &Product.Price, &Product.Quantity)

		if err != nil {
			panic(err.Error())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Product)

}


//UpdateProduct updates Product with respective ID
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()
	id := mux.Vars(r)["id"]
	product_category := r.FormValue("product_category")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")
	insForm, err := db.Prepare("UPDATE products SET product_category_id=?, name=?, description=?, price=?, quantity=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(product_category, name, description, price, quantity, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Update succeeded")
}

//DeletProduct delete's Product with specific ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()
	id := mux.Vars(r)["id"]

	delForm, err := db.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Delete succeeded")

}

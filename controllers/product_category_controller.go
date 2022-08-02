package controllers

import (
	"crud/database"
	"crud/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"html/template"
	"net/http"
)

//Display page
func IndexProductCategory(w http.ResponseWriter, r *http.Request) {
	var index_template = template.Must(template.ParseFiles("web/product_category.html"))
	err := index_template.ExecuteTemplate(w, "product_category.html", "Liste des categories de produit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//GetAllProductCategory get all ProductCategory data
func GetAllProductCategory(w http.ResponseWriter, r *http.Request) {
	var ProductCategories []entity.ProductCategory
	var ProductCategory entity.ProductCategory
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM product_categories")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&ProductCategory.Id, &ProductCategory.Name)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			ProductCategories = append(ProductCategories, ProductCategory)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ProductCategories)
}

//CreateProductCategory creates ProductCategory
func CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	name := r.FormValue("name")
	insertData, err := db.Prepare("INSERT INTO product_categories(name) VALUES(?)")
	if err != nil {
		log.Print(err)
		return
	}

	insertData.Exec(name)
	println("Insert data to database")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Create succeeded")
}

//GetOneProductCategory get's ProductCategory with specific ID
func GetOneProductCategory(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()
	id := mux.Vars(r)["id"]

	selDB, err := db.Query("SELECT * FROM product_categories WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	ProductCategory := entity.ProductCategory{}
	for selDB.Next() {
		err = selDB.Scan(&ProductCategory.Id, &ProductCategory.Name)

		if err != nil {
			panic(err.Error())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ProductCategory)

}

//UpdateProductCategory updates ProductCategory with respective ID
func UpdateProductCategory(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()
	id := mux.Vars(r)["id"]
	name := r.FormValue("name")
	insForm, err := db.Prepare("UPDATE product_categories SET name=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(name, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Update succeeded")
}

//DeletProductCategory delete's ProductCategory with specific ID
func DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()
	id := mux.Vars(r)["id"]

	delForm, err := db.Prepare("DELETE FROM product_categories WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Delete succeeded")

}

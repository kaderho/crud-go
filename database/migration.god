package database

import (
	"log"
	"crud/entity"
	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connexion réussie!")
	return nil
}

//Migrate create/updates database table
func MigrateProductCategory(table *entity.ProductCategory) {
	Connector.AutoMigrate(&table)
	log.Println("Migration réussie categorie de produit réussie")
}

//Migrate create/updates database table
func MigrateProduct(table *entity.Product) {
	Connector.AutoMigrate(&table)
	log.Println("Migration réussie produit réussie")
}
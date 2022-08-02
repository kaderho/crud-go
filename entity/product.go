package entity

//Person object for REST(CRUD)
type Product struct {
	Id       int    `json:"id"`
	ProductCategoryId int `json:"product_category_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price      float64 `json:"price"`
	Quantity    float64 `json:"quantity"`

}
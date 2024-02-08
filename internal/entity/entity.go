package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.NewString(),
		Name: name,
	}
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryId  string
	ImageUrl    string
}

func NewProduct(name, description, categoryId, imageUrl string, price float64) *Product {
	return &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryId:  categoryId,
		ImageUrl:    imageUrl,
	}
}

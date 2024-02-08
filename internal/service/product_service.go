package service

import (
	"github.com/danilolima05/fullcycle/goapi/internal/database"
	"github.com/danilolima05/fullcycle/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

func (cs *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (cs *ProductService) CreateProduct(name, description, categoryId, imageUrl string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryId, imageUrl, price)
	_, err := cs.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (cs *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := cs.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

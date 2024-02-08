package web

import (
	"encoding/json"
	"net/http"

	"github.com/danilolima05/fullcycle/goapi/internal/entity"
	"github.com/danilolima05/fullcycle/goapi/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{
		ProductService: productService,
	}
}

func (wch *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wch.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wch *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wch.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wch *WebProductHandler) GetProductByCategoryId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "categoryId")
	if id == "" {
		http.Error(w, "categoryId is required", http.StatusBadRequest)
		return
	}
	product, err := wch.ProductService.ProductDB.GetProductsByCategoryId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wch *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := wch.ProductService.CreateProduct(product.Name, product.Description, product.CategoryId, product.ImageUrl, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

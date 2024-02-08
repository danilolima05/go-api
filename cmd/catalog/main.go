package main

import (
	"database/sql"
	"fmt"

	"github.com/danilolima05/fullcycle/goapi/internal/database"
	"github.com/danilolima05/fullcycle/goapi/internal/service"
	"github.com/danilolima05/fullcycle/goapi/internal/web"
	"github.com/danilolima05/fullcycle/goapi/internal/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "catalog"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productDb := database.NewProductDB(db)
	createProductService := service.NewProductService(*productDb)
	productHandler := web.NewWebProductHandler(createProductService)

	categoryDb := database.NewCategoryDB(db)
	createCategoryService := service.NewCategoryService(*categoryDb)
	categoryHandler := web.NewWebCategoryHandler(createCategoryService)

	webserver := webserver.NewWebServer(":8080")

	webserver.AddGetHandler("/product", productHandler.GetProducts)
	webserver.AddGetHandler("/product/{id}", productHandler.GetProduct)
	webserver.AddGetHandler("/product/category/{categoryId}", productHandler.GetProductByCategoryId)
	webserver.AddPostHandler("/product", productHandler.CreateProduct)

	webserver.AddGetHandler("/category/{id}", categoryHandler.GetCategory)
	webserver.AddGetHandler("/category", categoryHandler.GetCategories)
	webserver.AddPostHandler("/category", categoryHandler.CreateCategory)

	fmt.Println("Server is running")
	webserver.Start()
}

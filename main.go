package main

import (
	"fmt"
	_ "postmanapi/docs"
	"postmanapi/handlers"
	"postmanapi/storage"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

func main() {
	r := gin.Default()

	_, err := storage.InitDB()
	if err != nil {
		fmt.Println("db not found")
	}

	r.GET("/users", handlers.GetUsersHandler)
	r.GET("/users/:id", handlers.GetUserHandler)
	r.GET("/admin_panel/products", handlers.GetProductsHandler)
	r.POST("/users", handlers.CreateUserHandler)
	r.PUT("/products/:id", handlers.UpdateProductHandler)
	r.DELETE("/products/:id", handlers.DeleteProductHandler)

	r.POST("/admin/login", handlers.AdminLoginHandler)
	r.POST("/admin_panel/products/add_product", handlers.AddProductsHandler)

	// Swagger setup important..(to serve swagger index ...pages)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

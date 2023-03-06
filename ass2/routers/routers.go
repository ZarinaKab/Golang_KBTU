package routers

import (
	"github.com/gin-gonic/gin"
	"ass2/authorization"
	"ass2/product_manipulation"
	"ass2/registration"
)

func Routers() {
	router := gin.Default()
	router.PUT("/productsUpdate/:id", product_manipulation.UpdateProduct)

	router.GET("/products", product_manipulation.ListOfProducts)
	router.GET("/products/:ID", product_manipulation.GetProducts)
	router.GET("/search", product_manipulation.SearchProduct)
	router.GET("/sort", product_manipulation.SortProducts)

	router.POST("/createProduct", product_manipulation.CreateProducts)

	router.DELETE("/deleteProduct/:ID", product_manipulation.DeleteProducts)

	router.POST("/register", registration.RegisterUser)
	router.POST("/login", authorization.Login)

	router.Run(":8080")
}

package routers

import (
	"github.com/gin-gonic/gin"
	"golang_ass_2/authorization"
	"golang_ass_2/product_manipulation"
	"golang_ass_2/registration"
)

func Routers() {
	router := gin.Default()
	router.PUT("/bookUpdate/:id", product_manipulation.UpdateProduct)

	router.GET("/books", product_manipulation.ListOfProducts)
	router.GET("/books/:ID", product_manipulation.GetProducts)
	router.GET("/search", product_manipulation.SearchProduct)
	router.GET("/sort", product_manipulation.SortProducts)

	router.POST("/createBook", product_manipulation.CreateProducts)

	router.DELETE("/deleteBook/:ID", product_manipulation.DeleteProducts)

	router.POST("/register", registration.RegisterUser)
	router.POST("/login", authorization.Login)

	router.Run(":8080")
}

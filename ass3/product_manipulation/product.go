package product_manipulation

import (
	"fmt"
	"net/http"

	"golang_ass_2/database"

	"github.com/gin-gonic/gin"
)

func ListOfProducts(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}

	products := []database.Books{}
	db.Find(&products)
	val.JSON(http.StatusOK, products)
}

func GetProducts(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}

	productID := val.Param("ID")
	product := database.Books{}

	db.First(&product, productID)
	if product.ID == 0 {
		val.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Books not found or not exist"})
		return
	}
	val.JSON(http.StatusAccepted, product)
}

func CreateProducts(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}
	var productt database.Books
	if err := val.ShouldBindJSON(&productt); err != nil {
		val.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&productt)
	val.JSON(http.StatusAccepted, productt)
}

func DeleteProducts(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}

	productID := val.Param("ID")
	product := database.Books{}
	db.First(&product, productID)
	if product.ID == 0 {
		val.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Books not found or not exist!"})
		return
	}
	db.Delete(&product)
	val.JSON(http.StatusAccepted, gin.H{"message": "Books successfully deleted!"})
}

func SearchProduct(val *gin.Context) {
	query := val.Query("title")

	products := []database.Books{}
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}
	db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", query)).Find(&products)

	val.JSON(http.StatusAccepted, products)
}
func UpdateProduct(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}
	productID := val.Param("ID")
	existingProduct := database.Books{}
	db.First(&existingProduct, productID)
	if existingProduct.ID == 0 {
		val.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Books not found"})
		return
	}
	var updatedProductData database.Books
	if err := val.ShouldBindJSON(&updatedProductData); err != nil {
		val.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updatedProductData.Title != "" {
		existingProduct.Title = updatedProductData.Title
	}
	if updatedProductData.Cost != 0 {
		existingProduct.Cost = updatedProductData.Cost
	}
	db.Save(&existingProduct)
	val.JSON(http.StatusOK, existingProduct)
}

func SortProducts(val *gin.Context) {
	var products = []database.Books{}
	sortBy := val.Query("sortBy")
	sortOrder := val.Query("sortOrder")
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}
	query := db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
	query.Find(&products)

	val.JSON(http.StatusAccepted, products)
}

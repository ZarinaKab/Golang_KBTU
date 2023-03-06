package product_manipulation

import (
	"fmt"
	"net/http"

	"ass2/database"

	"github.com/gin-gonic/gin"
)

func ListOfProducts(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}

	products := []database.Product{}
	db.Find(&products)
	val.JSON(http.StatusOK, products)
}

func GetProducts(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}

	productID := val.Param("ID")
	product := database.Product{}

	db.First(&product, productID)
	if product.ID == 0 {
		val.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found or not exist"})
		return
	}
	val.JSON(http.StatusAccepted, product)
}

func CreateProducts(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}
	var productt database.Product
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
	product := database.Product{}
	db.First(&product, productID)
	if product.ID == 0 {
		val.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found or not exist!"})
		return
	}
	db.Delete(&product)
	val.JSON(http.StatusAccepted, gin.H{"message": "Product successfully deleted!"})
}

func SearchProduct(val *gin.Context) {
	query := val.Query("name")

	products := []database.Product{}
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}
	db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", query)).Find(&products)

	val.JSON(http.StatusAccepted, products)
}
func UpdateProduct(val *gin.Context) {
	db, err := database.GETDB()
	if err != nil {
		panic(err)
	}
	productID := val.Param("ID")
	existingProduct := database.Product{}
	db.First(&existingProduct, productID)
	if existingProduct.ID == 0 {
		val.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	var updatedProductData database.Product
	if err := val.ShouldBindJSON(&updatedProductData); err != nil {
		val.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updatedProductData.Name != "" {
		existingProduct.Name = updatedProductData.Name
	}
	if updatedProductData.Price != 0 {
		existingProduct.Price = updatedProductData.Price
	}
	if updatedProductData.Rating != 0 {
		existingProduct.Rating = updatedProductData.Rating
	}
	db.Save(&existingProduct)
	val.JSON(http.StatusOK, existingProduct)
}

func SortProducts(val *gin.Context) {
	var products = []database.Product{}
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

package authorization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"ass2/database"
)

func Login(val *gin.Context) {
	var enter struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := val.ShouldBindJSON(&enter); err != nil {
		val.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Get the user by name
	user := database.User{}
	if err := user.Username(enter.Name); err != nil {
		val.JSON(http.StatusNotFound, gin.H{"error": "Wrong username"})
		return
	}

	//Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(enter.Password)); err != nil {
		val.JSON(http.StatusNotFound, gin.H{"error": "Wrong password"})
		return
	}

	val.JSON(http.StatusAccepted, gin.H{"message": "Welcome back"})
}

package registration

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang_ass_2/database"
	"net/http"
)

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return bytes, err
}

func RegisterUser(val *gin.Context) {
	var enter struct {
		Userame     string `json:"username" binding:"required"`
		Password 	string `json:"password" binding:"required"`
	}

	if err := val.ShouldBindJSON(&enter); err != nil {
		val.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(enter.Password)
	if err != nil {
		val.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := database.User{
		ID:       0,
		Username:     enter.Userame,
		Password: string(hashedPassword),
	}

	if err := user.CreateUser(); err != nil {
		val.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	val.JSON(http.StatusCreated, gin.H{"message": "Congratulations! User created."})
}

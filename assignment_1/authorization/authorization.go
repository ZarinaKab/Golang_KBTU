package authorization

import (
	"example.com/packages/Golang_KBTU/assignment_1/packages"
	"golang.org/x/crypto/bcrypt"
)

type Users struct{
	*packages.DatabaseUsers
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (users *Users) SignIn(user, password string) string{
	token := "token"
	for _, u := range users.Users{
		if u.Name == user && u.Password == password{
			return token
		}
	}
	return token
}
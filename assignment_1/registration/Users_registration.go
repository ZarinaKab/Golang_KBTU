package registration

import (
	"example.com/packages/Golang_KBTU/assignment_1/packages"
	"golang.org/x/crypto/bcrypt"
)

type Registration struct {
	Name       string
	SecondName string
	Age        int
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

type dbUsers struct {
	*packages.DatabaseUsers
}

func (d *dbUsers) Register(name string, password string) {
	pass, err := HashPassword(password)
	if err != nil {
		return
	}
	d.Users = append(d.Users, packages.User{Name: name, Password: pass})
}
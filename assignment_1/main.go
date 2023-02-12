package main

import (
	"fmt"
	"example.com/packages/Golang_KBTU/assignment_1/packages"
	"golang.org/x/crypto/bcrypt"
	// "github.com/ZarinaKab/Golang_KBTU/tree/main/assignment_1/packages"
)

type Registration struct {
	Name       string
	SecondName string
	Age        int
}

func main(){
	pas := "secret"
	has, err := HashPassword(pas)
	match := CheckPasswordHash(pas, has)
	if err != nil{
		fmt.Println("err")
	} else if !match{
		fmt.Println("not match")
	} else {
		fmt.Println(pas, has)
	}
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

type dbUsers struct {
	*packages.DatabaseUsers
}

func (d *dbUsers) Register(name string, password string) {
	pass, err := HashPassword(password)
	if err != nil{
		return
	}
	d.Users = append(d.Users, packages.User{Name: name, Password: pass})
}
package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (c *Collection) SignUp(username, password string) {
	ok := false
	c.UserIterator++
	// if c.GetUser(username) == true {
	for _, elem := range c.Users {
		// fmt.Print(elem.Username, " ")
		if elem.Username == username {
			ok = true
			// break
		}
	}
	user := User{c.UserIterator, username, password}
	if ok == true {
		fmt.Println("User:", user.Username, "already exists.")
		// break
	} else {
		// c.UserSaveData()
		c.Users = append(c.Users, user)
		fmt.Println("User:", user.Username, "was created successfully.")
		// break
	}
}

func (c *Collection) SignIn(username, password string) bool {
	for _, col := range c.Users {
		if col.Username == username && col.Password == password && c.GetUser(username) == true{
			fmt.Println("User:", username, "authorized.")
			return true
		}
	}
	// fmt.Println("User:", username, "failed the authorization.")
	return false
}

func (c *Collection) GetUser(username string) bool {
	// var result []User
	for _, user := range c.Users {
		if user.Username == username {
			// result = append(result, user)
			return true
		}
	}
	return false
	// return result
}

func (c *Collection) UserTakeData() {
	// var collect []string
	file, err := os.Open("user.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		u := strings.Split(line, ", ")
		n, err := strconv.Atoi(u[0])
		if err != nil {
			fmt.Println(err)
		}
		user := User{n, u[1], u[2]}
		c.Users = append(c.Users, user)
		c.UserIterator = n
	}
}

func (c *Collection) UserSaveData() {
	if err := os.Truncate("user.txt", 0); err != nil {
		fmt.Println(err)
	}
	file, err := os.OpenFile("user.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, user := range c.Users {
		// c.Users = append(c.Users, user.Username)
		if _, err := file.WriteString(strconv.Itoa(user.Id) + ", " + user.Username + ", " + user.Password + "\n"); err != nil {
			fmt.Println(err)
		}
	}
}

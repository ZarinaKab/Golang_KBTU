package main

import (
	"fmt"

	functions "example.com/packages/Golang_KBTU/functions"
)

func main() {
	sms := functions.Collection{
		Users:        make([]functions.User, 0),
		Items:        make([]functions.Item, 0),
		ItemIterator: 0,
		UserIterator: 0,
	}
	var name, password string
	sms.UserTakeData()
	sms.ItemTakeData()
	var operation int
	isAuthorized := false
	for true {
		fmt.Print("1 - add user\n2 - authorize\n3 - manage items\n4 - exit the program\n5 - log out\n")
		fmt.Scanln(&operation)
		if operation == 1 {
			fmt.Print("Enter a username: ")
			fmt.Scanln(&name)
			fmt.Print("Enter a password: ")
			fmt.Scanln(&password)
			sms.SignUp(name, password)
			sms.SignIn(name, password)
		} else if operation == 2 {
			fmt.Println("Type existing user to authorize")
			fmt.Print("Enter a username: ")
			fmt.Scanln(&name)
			fmt.Print("Enter a password: ")
			fmt.Scanln(&password)
			// sms.SignIn(name, password)
			if sms.SignIn(name, password) {
				// fmt.Println("Alright, you can manage the items!")
				isAuthorized = true
			} else {
				fmt.Println("The data you have wrote is invalid! Try again.")
			}
		} else if operation == 3 {
			var itemOper int
			var itemName string
			var rating, price float64
			if isAuthorized {
				fmt.Println("1 - add item\n2 - filter item by price\n3 - filter item by rating\n4 - search item by name\n5 - set rating for item")
				fmt.Print("Choose operation for item: ")
				fmt.Scanln(&itemOper)
				if itemOper == 1 {
					fmt.Println("Write item you want to add")
					fmt.Print("Write name: ")
					fmt.Scanln(&itemName)
					fmt.Print("Write price: ")
					fmt.Scanln(&price)
					fmt.Print("Write rating: ")
					fmt.Scanln(&rating)
					// sms.UserSaveData()
					sms.ItemPush(itemName, price, rating)
					sms.ItemSaveData()
				} else if itemOper == 2 {
					fmt.Print("Write item price you want to sort: ")
					fmt.Scanln(&price)
					for _, element := range sms.FilterItemsByPrice(price) {
						fmt.Println("Item:", element.ItemName, ", Price:", element.Price)
					}
					fmt.Println(sms.FilterItemsByPrice(price))
				} else if itemOper == 3 {
					fmt.Print("Write item rating you want to sort: ")
					fmt.Scanln(&rating)
					for _, element := range sms.FilterItemsByRating(rating) {
						fmt.Println("ItemName:", element.ItemName, ", Price:", element.Price, ", Rating:", element.Rating)
					}
					// fmt.Println(sms.FilterItemsByRating(rating))
				} else if itemOper == 4 {
					fmt.Print("Write item name you want to find: ")
					fmt.Scanln(&itemName)
					for _, element := range sms.SearchItemsByName(itemName) {
						fmt.Println("ItemName:", element.ItemName, ", Price:", element.Price, ", Rating:", element.Rating)
					}
					// fmt.Println(sms.SearchItemsByName(itemName))
				} else if itemOper == 5 {
					var itemId int
					fmt.Println("To set rating of existing item, you have to write id and rating.")
					fmt.Print("item id: ")
					fmt.Scanln(&itemId)
					fmt.Print("rating of item: ")
					fmt.Scanln(&rating)
					sms.SetRating(rating, itemId)
				} else {
					fmt.Println("You wrote incorrect operation for item managing!")
				}
			} else {
				fmt.Println("Firstly, you have to authorize. This operation is not available for you!")
			}
		} else if operation == 4 {
			fmt.Print("Are you sure?(yes or no): ")
			var isSure string
			fmt.Scanln(&isSure)
			if isSure == "yes" {
				sms.UserSaveData()
				// sms.ItemSaveData()
				fmt.Println("Come back again!!!")
				break
			} else {
				continue
			}
		} else if operation == 5 {
			isAuthorized = false
			fmt.Println("You have just logged out.")
		} else {
			fmt.Println("Operation doesn't exist, try again.")
		}
	}
}
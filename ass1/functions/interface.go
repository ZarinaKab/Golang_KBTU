package functions

type UserManager interface {
	SignUp(username, password string)
	SignIn(username, password string) bool
	GetUser(username string) bool
}

type ItemManager interface {
	ItemPush(name string, price, rating float64)
	SearchItemsByName(name string) []Item
	FilterItemsByPrice(price float64) []Item
	FilterItemsByRating(rating float64) []Item
	SetRating(rating float64, id int)
	ItemTakeData()
	ItemSaveData()
}

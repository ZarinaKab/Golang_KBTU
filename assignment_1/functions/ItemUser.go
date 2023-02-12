package functions

type User struct {
	Id       int
	Username string
	Password string
}

type Item struct {
	Id       int
	ItemName string
	Price    float64
	Rating   float64
}

type Collection struct {
	UserIterator, ItemIterator int
	Users                      []User
	Items                      []Item
}
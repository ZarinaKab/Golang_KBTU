package packages

type Item struct {
	Name   string
	Price  float64
	Rating float64
}

type DatabaseItems struct {
	Items []Item
}
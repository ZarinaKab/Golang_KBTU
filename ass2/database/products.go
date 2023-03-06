package database

type Product struct {
	ID     uint   `json:"ID" gorm:"primaryKey"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
	Price  int    `json:"price"`
}

func (AboutProduct *Product) ProductID(id uint) error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.First(AboutProduct, id)
	return result.Error
}

func (AboutProduct *Product) Create() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Create(AboutProduct)
	return result.Error
}

func (AboutProduct *Product) Update() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}
	result := db.Save(AboutProduct)
	return result.Error
}

func (AboutProduct *Product) Delete() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Delete(AboutProduct)
	return result.Error
}

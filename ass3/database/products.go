package database

type Books struct {
	ID            uint      `json:"ID" gorm:"primaryKey"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Cost          int       `json:"cost"`
}

func (AboutBooks *Books) BooksID(id uint) error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.First(AboutBooks, id)
	return result.Error
}

func (AboutBooks *Books) Create() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Create(AboutBooks)
	return result.Error
}

func (AboutBooks *Books) Update() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}
	result := db.Save(AboutBooks)
	return result.Error
}

func (AboutBooks *Books) Delete() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Delete(AboutBooks)
	return result.Error
}

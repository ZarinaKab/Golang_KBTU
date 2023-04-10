package database

type User struct {
	ID       		uint   		`json:"ID" gorm:"primaryKey"`
	Username     	string 		`json:"username"`
	Password 		string 		`json:"-"`
}

func (AboutUser *User) Name(username string) error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Where("username = ?", username).First(AboutUser)
	return result.Error
}

func (AboutUser *User) CreateUser() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Create(AboutUser)
	return result.Error
}

func (AboutUser *User) Update() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Save(AboutUser)
	return result.Error
}

func (AboutUser *User) Delete() error {
	db, err := GETDB()
	if err != nil {
		panic(err)
	}

	result := db.Delete(AboutUser)

	return result.Error
}

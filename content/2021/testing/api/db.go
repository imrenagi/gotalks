package api

import "gorm.io/gorm"

type User struct {
	Name string
}

//START,OMIT
type UserFetcher struct {
	DB *gorm.DB // Source of non-determinism // HL
}

func (u UserFetcher) FindByID(ID string) (*User, error) {
	var obj User
	tx := u.DB.Where("id = ?", ID).First(&obj)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &obj, nil
}

//STOP,OMIT

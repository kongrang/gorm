package userModel

import "gorm.io/gorm"

type UserModelHelper struct {
	DB *gorm.DB
}

func (u *UserModelHelper) Insert(user *User) error {
	result := u.DB.Debug().Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserModelHelper) FindById(id string) (*User, error) {
	user := User{}
	result := u.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

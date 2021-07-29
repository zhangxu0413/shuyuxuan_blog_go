package Models

import (
	"fmt"
	"shuyuxuan_blog_go/Types"
)



type iUser interface {
	AddUser(user Types.User) error
	DeleteUser(ID int) error
	updateUser(user Types.User) error
	getUser(ID int) ([]Types.User, error)
	getUserList() ([]Types.User, error)
}

type UserModel struct {}

func (m UserModel) AddUser(user Types.User) error {
	var err error
	DB.AutoMigrate(&Types.User{})
	DB.Create(&user)
	return err
}

func (m UserModel) UpdateUser(user Types.User) error {
	var err error
	if err := DB.Save(&user).Error; err != nil {
		fmt.Errorf("Save user failed, err: %+v\n", err)
	}
	return err
}

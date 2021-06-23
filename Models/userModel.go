package Models

import (
	"fmt"
	Mysql "shuyuxuan_blog_go/Databases"
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

var DB = Mysql.DB

func (m UserModel) AddUser(user Types.User) error {
	var err error
	fmt.Printf("%v", user)
	//if !DB.HasTable(&Types.User{}) {
	//	if err = DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Types.User{}).Error; err != nil {
	//		panic(err)
	//	}
	//}
	DB.AutoMigrate(&Types.User {})
	//DB.Create(user)
	return err
}

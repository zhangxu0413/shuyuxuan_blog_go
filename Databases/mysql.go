package Mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB
type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}
func Init()  {
	var err error
	DB, err = gorm.Open("mysql", "zhangshuyu:zx644398058@tcp(118.25.228.199:3306)/shuyuxuan")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
	defer DB.Close()

	if !DB.HasTable(&Like{}) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}
	DB.AutoMigrate(&Like {})
	l1 := &Like{2, "192.168.150.18","Test", "测试",1231,time.Now() }
	DB.Create(&l1)
}
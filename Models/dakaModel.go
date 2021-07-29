package Models

import (
	"fmt"
	"shuyuxuan_blog_go/Types"
)

type DakaModel struct{}

// AddRecord 新增健身打卡记录
func (b DakaModel) AddRecord(record Types.DakaRecord) error {
	var err error
	fmt.Printf("%v",DB.HasTable(&Types.DakaRecord{})  )
	if !DB.HasTable(&Types.DakaRecord{}) {
		DB.Set("gorm:table_options","ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Types.DakaRecord{})
	}
	DB.AutoMigrate(&Types.DakaRecord{})
	err = DB.Create(&record).Error
	return err
}

func (b DakaModel) GetList() []Types.DakaRecord {
	var records []Types.DakaRecord
	DB.Find(&records)
	return records
}

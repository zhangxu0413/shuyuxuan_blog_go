package Services

import (
	"shuyuxuan_blog_go/Models"
	"shuyuxuan_blog_go/Types"
	"time"
)


var dakaModel Models.DakaModel

type DakaService struct {}

func (d DakaService ) AddRecord(req Types.DakaReq) error  {
	var record Types.DakaRecord
	record.Title = req.Title
	record.Content = req.Content
	record.Time = req.Time
	record.Photo = req.Photo
	t := time.Now()
	record.CreatTime = t.Format("2006-01-02 15:04:05")
	var err error
	err = dakaModel.AddRecord(record)
	return err
}

func (d DakaService) GetList() []Types.DakaReq  {
	var list []Types.DakaReq
	records := dakaModel.GetList()
	for i:= 0;i<len(records);i++ {
		list[i].Title = records[i].Title
		list[i].Content = records[i].Content
		list[i].Time = records[i].Time
	}
	return list
}
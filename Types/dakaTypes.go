package Types

type DakaReq struct {
	Title      	string  `json:"title"`
	Time      	string  `json:"time"`
	Content		string	`json:"content"`
	Photo		string	`json:"photo"`
}

type DakaRecord struct {
	ID			int    `gorm:"primary_key;AUTO_INCREMENT"`
	Time 	  	string	`gorm:"type:varchar(20);not null;"`
	Title     	string `gorm:"type:varchar(256);not null;"`
	Content   	string `gorm:"type:varchar(256);not null;"`
	Photo     	string `gorm:"type:varchar(256);"`
	CreatTime 	string `gorm:"type:varchar(20);not null;"`
}

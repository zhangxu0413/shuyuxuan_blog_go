package Types

type UserReq struct {
	Id        	int     `json:"id"`
	Name      	string  `json:"name"`
	Mail      	string  `json:"mail"`
	Phone		string  `json:"phone"`
	Description	string  `json:"description"`
}

type User struct {
	ID        	int     `gorm:"primary_key;AUTO_INCREMENT"`
	Name      	string  `gorm:"type:varchar(20);not null;index:name_idx"`
	Mail      	string  `gorm:"type:varchar(256);not null;"`
	Phone		string  `gorm:"type:varchar(128);not null;index:phone_idx"`
	Description	string  `gorm:"type:varchar(128);not null"`
	Status		uint	`gorm:"default:0"`
	CreatedAt 	string  `gorm:"type:varchar(128)"`
}

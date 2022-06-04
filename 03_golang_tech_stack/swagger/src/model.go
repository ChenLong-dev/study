package src

import (
	"gorm.io/gorm"
)

//this model represent a database table
// {"title":"chenlong", "des": "alexchen", "status":"study golong"}
type Post struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);" json:"title" example:"title" binding:"required"`
	Des    string `gorm:"type:varchar(100);" json:"des" example:"desc" binding:"required"`
	Status string `gorm:"type:varchar(200);" json:"status" example:"Active"`
}

type Response struct {
	Msg  string
	Data interface{}
}


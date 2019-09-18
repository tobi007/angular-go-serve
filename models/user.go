package models

import (
	"time"
)

type User struct {
	Email    	string 		`gorm:"primary_key";json:"email"`
	Phone    	string 		`json:"phone"`
	Company    	string 		`json:"company"`
	Password  	string 		`json:"password"`
	CreatedAt 	time.Time 	`json:"-"`
	UpdatedAt 	time.Time 	`json:"-"`
	DeletedAt 	*time.Time	`json:"-"`
}

func (User) TableName() string {
	return "ThirdPartyUser"
}

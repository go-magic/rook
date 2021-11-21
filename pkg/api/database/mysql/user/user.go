package user

import "time"

type User struct {
	ID       int `gorm:"primary_key"`
	UserName string
	PickName string
	CreateAt time.Time
}

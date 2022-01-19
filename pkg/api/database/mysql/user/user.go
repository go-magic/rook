package user

import (
	"errors"
	"github.com/go-magic/rook/pkg/api/database/mysql"
	"time"
)

type User struct {
	ID       uint64 `gorm:"primary_key"`
	UserName string
	PickName string
	Passwd   string
	CreateAt time.Time
}

func GetUserByUserId(userId uint64) (*User, error) {
	user := &User{ID: userId}
	db := mysql.GetMysqlInstance().GetDB()
	if db == nil {
		return nil, errors.New("db invalid")
	}
	if err := db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{UserName: username}
	if err := mysql.GetMysqlInstance().GetDB().First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

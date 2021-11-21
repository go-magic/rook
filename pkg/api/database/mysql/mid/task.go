package mid

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	ID       int    `gorm:"primary_key"`
	MsgID    int    `gorm:"msgId"`
	SubTask  string `gorm:"type:varchar(1000);not null"`
	CreateAt time.Time
}

type Result struct {
}

func (t *Task) Insert(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Task) Update(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Task) Delete(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Task) Select(db *gorm.DB) error {
	return db.Select(t, t).Error
}

func (t *Task) Execute(db *gorm.DB, sql string) error {
	return db.Select(t, t).Error
}

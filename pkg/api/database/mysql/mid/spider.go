package mid

import "github.com/jinzhu/gorm"

type Spider struct {
}

func (s *Spider) Insert(db *gorm.DB) error {
	return db.Create(s).Error
}

func (s *Spider) Update(db *gorm.DB) error {
	return db.Create(s).Error
}

func (s *Spider) Delete(db *gorm.DB) error {
	return db.Create(s).Error
}

func (s *Spider) Select(db *gorm.DB) error {
	return db.Select(s, s).Error
}

func (s *Spider) Execute(db *gorm.DB, sql string) error {
	return db.Select(s, s).Error
}

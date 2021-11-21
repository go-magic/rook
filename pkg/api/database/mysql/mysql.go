package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysql *Mysql

type Mysql struct {
	db *gorm.DB
}

func InitMysql(mysqlAddr string) error {
	mysql = &Mysql{}
	db, err := gorm.Open("mysql", mysqlAddr)
	if err != nil {
		return err
	}
	mysql.db = db
	db.SingularTable(true)
	return nil
}

func GetMysqlInstance() *Mysql {
	return mysql
}

/*
初始化表,由业务自己实现
*/
func (m Mysql) InitTables(tables ...interface{}) error {
	return m.db.AutoMigrate(tables...).Error
}

func (m Mysql) Close() error {
	return m.db.Close()
}

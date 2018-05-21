package models

import (
	"fmt"
	"github.com/jinzhu/gorm"

	// mysql db driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// NewMysqlDB - mysql database
func NewMysqlDB(user, password, dbname string) *DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname))
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(false)

	return &DB{db}
}

package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(db *gorm.DB) *Dao {
	return &Dao{db}
}

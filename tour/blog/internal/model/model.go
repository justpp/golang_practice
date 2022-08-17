package model

import (
	"fmt"
	"giao/tour/blog/global"
	"giao/tour/blog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(settings *setting.DatabaseSettings) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(settings.DBType, fmt.Sprintf(s,
		settings.UserName,
		settings.Password,
		settings.Host,
		settings.DBName,
		settings.Charset,
		settings.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConnects)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConnects)
	return db, nil
}

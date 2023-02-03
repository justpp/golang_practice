package model

import (
	"fmt"
	"giao/tour/blog/global"
	"giao/tour/blog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimestampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimestampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConnects)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConnects)
	return db, nil
}
func updateTimestampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreateOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if updateTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if updateTimeField.IsBlank {
				_ = updateTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimestampForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if updateField, ok := scope.FieldByName("ModifiedOn"); ok {
			_ = updateField.Set(time.Now().Unix())
		}
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deleteField, hasDeletedField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDel := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedField && hasIsDel {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

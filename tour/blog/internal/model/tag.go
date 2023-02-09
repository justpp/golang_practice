package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageSize >= 0 && pageOffset >= 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, err
}

func (t Tag) Create(db *gorm.DB) error {
	res := db.Create(&t)

	return res.Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Model(&Tag{}).Where("id = ? and is_del = 0", t.ID).Update(t).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = 0", t.ID).Delete(&t).Error
}

func (t Tag) Find(db *gorm.DB, id uint32) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ? and is_del = 0", id).First(&tag).Error
	return &tag, err
}

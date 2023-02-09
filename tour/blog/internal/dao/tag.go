package dao

import (
	"fmt"
	model "giao/tour/blog/internal/model"
	"giao/tour/blog/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	fmt.Println("GetTagList ", tag.TableName())
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) Find(id uint32) (*model.Tag, error) {
	var tag model.Tag
	fmt.Println("Find ", tag.TableName())
	fmt.Println("Dao model", tag)
	return tag.Find(d.engine, id)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	}
	return tag.Update(d.engine)
}

func (d *Dao) Del(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

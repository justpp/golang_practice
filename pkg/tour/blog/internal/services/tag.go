package services

import (
	"fmt"
	"giao/pkg/tour/blog/internal/model"
	"giao/pkg/tour/blog/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ListTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `json:"created_by" binding:"required,min=2,max=100"`
}

type UpdateTagRequest struct {
	Id         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by"  json:"modified_by" binding:"required,min=2,max=100"`
}

type DelTagRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

func (s *Service) CountTag(param *CountTagRequest) (int, error) {
	return s.dao.CountTag(param.Name, param.State)
}

func (s *Service) ListTag(param *ListTagRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s *Service) TagById(id uint32) (*model.Tag, error) {
	fmt.Println("service", id)
	return s.dao.Find(id)
}

func (s *Service) CreateTag(param *CreateTagRequest) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (s *Service) UpdateTag(param *UpdateTagRequest) error {
	return s.dao.UpdateTag(param.Id, param.Name, param.State, param.ModifiedBy)
}

// @Summary sdff
func (s *Service) DelTag(param *DelTagRequest) error {
	return s.dao.Del(param.Id)
}

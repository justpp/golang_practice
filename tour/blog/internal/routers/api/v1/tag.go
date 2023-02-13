package v1

import (
	"giao/tour/blog/global"
	"giao/tour/blog/internal/services"
	"giao/tour/blog/pkg/app"
	"giao/tour/blog/pkg/convert"
	"giao/tour/blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) List(c *gin.Context) {
	param := services.ListTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := services.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&services.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.ListTag(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

func (t Tag) Create(c *gin.Context) {
	param := services.CreateTagRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := services.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (t Tag) Get(c *gin.Context) {
	response := app.NewResponse(c)
	Id := convert.StrTo(c.Param("id")).MustUInt32()
	svc := services.New(c.Request.Context())
	tag, err := svc.TagById(Id)
	if err != nil {
		global.Logger.Errorf("svc.TagById err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagFail)
		return
	}
	response.ToResponse(tag)
}
func (t Tag) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	params := services.DelTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	svc := services.New(c.Request.Context())
	err := svc.DelTag(&params)
	if err != nil {
		global.Logger.Errorf("svc.DelTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagFail)
		return
	}
	response.ToResponse(gin.H{})
}
func (t Tag) Update(c *gin.Context) {
	response := app.NewResponse(c)
	param := services.UpdateTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := services.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
}

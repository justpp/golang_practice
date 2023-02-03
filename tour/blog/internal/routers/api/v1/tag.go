package v1

import (
	"giao/tour/blog/global"
	"giao/tour/blog/internal/services"
	"giao/tour/blog/pkg/app"
	"giao/tour/blog/pkg/errorcode"
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
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := services.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&services.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errorcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.ListTag(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errorcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}
func (t Tag) Get(c *gin.Context) {
}
func (t Tag) Create(*gin.Context) {}
func (t Tag) Update(*gin.Context) {}
func (t Tag) Delete(*gin.Context) {}

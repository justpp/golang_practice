package v1

import (
	"giao/tour/blog/global"
	"giao/tour/blog/pkg/app"
	"giao/tour/blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (a Article) Get(c *gin.Context) {
	global.Logger.WithContext(c).Info("嘿嘿")
	global.Logger.Info("hihi")
	app.NewResponse(c).ToResponse("234234")
	return
}
func (a Article) Create(*gin.Context) {}
func (a Article) Update(*gin.Context) {}
func (a Article) Delete(*gin.Context) {}

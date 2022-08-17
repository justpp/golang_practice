package v1

import (
	"giao/tour/blog/pkg/app"
	"giao/tour/blog/pkg/errorcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) List(*gin.Context) {}
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errorcode.ServerError)
	return
}
func (a Article) Create(*gin.Context) {}
func (a Article) Update(*gin.Context) {}
func (a Article) Delete(*gin.Context) {}

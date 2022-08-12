package v1

import "github.com/gin-gonic/gin"

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) List(*gin.Context)   {}
func (a Article) Get(*gin.Context)    {}
func (a Article) Create(*gin.Context) {}
func (a Article) Update(*gin.Context) {}
func (a Article) Delete(*gin.Context) {}

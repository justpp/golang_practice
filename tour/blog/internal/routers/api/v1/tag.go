package v1

import "github.com/gin-gonic/gin"

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) List(*gin.Context)   {}
func (t Tag) Get(*gin.Context)    {}
func (t Tag) Create(*gin.Context) {}
func (t Tag) Update(*gin.Context) {}
func (t Tag) Delete(*gin.Context) {}

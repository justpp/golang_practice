package v1

import (
	"fmt"
	"giao/tour/blog/global"
	"github.com/gin-gonic/gin"
	"log"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) List(*gin.Context) {
	log.Printf("9999999999999%s", "2344444444444444444")
	fmt.Println("99999999999999999")
	global.Logger.WithCallersFrames().Info("234234")
}
func (t Tag) Get(*gin.Context)    {}
func (t Tag) Create(*gin.Context) {}
func (t Tag) Update(*gin.Context) {}
func (t Tag) Delete(*gin.Context) {}

package api

import (
	"fmt"
	"giao/tour/blog/global"
	"giao/tour/blog/internal/services"
	"giao/tour/blog/pkg/app"
	"giao/tour/blog/pkg/convert"
	"giao/tour/blog/pkg/errcode"
	"giao/tour/blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fmt.Println("uploadFile", file, fileHeader, err)

	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	fmt.Println("fileType", fileType)
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := services.New(c)
	uploadFile, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		return
	}
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{"file_access_url": uploadFile.AccessUrl})
}

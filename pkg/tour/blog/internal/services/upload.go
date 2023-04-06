package services

import (
	"errors"
	"giao/pkg/tour/blog/global"
	"giao/pkg/tour/blog/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	filename := upload.GetFileName(fileHeader.Filename)
	// 检查后缀类型
	if !upload.CheckContainExt(fileType, filename) {
		return nil, errors.New("file suffix is not allowed")
	}

	// 检查大小
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	// 创建文件夹
	savePath := upload.GetSaveDir()
	if upload.CheckSavePath(savePath) {
		err := upload.CreateSavePath(savePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create dir:" + err.Error())
		}
	}

	// 检查权限
	if upload.CheckPermission(savePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := savePath + "/" + filename
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + filename
	return &FileInfo{Name: filename, AccessUrl: accessUrl}, nil

}

package service

import (
	"admin-api/common/config"
	"admin-api/common/result"
	"admin-api/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"time"
)

type IUploadService interface {
	Upload(c *gin.Context)
}

type UploadServiceImpl struct{}

func (u UploadServiceImpl) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FILEUPLOADERROR),
			result.ApiCode.GetMessage(result.ApiCode.FILEUPLOADERROR))
	}
	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s",
		config.ImageSetting.UploadDir,
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%04d", now.Day()))
	util.CreateDir(filePath)
	fullPath := filePath + "/" + fileName
	c.SaveUploadedFile(file, fullPath)
	result.Success(c, fullPath)
}

var uploadService = UploadServiceImpl{}

func UploadService() IUploadService {
	return &uploadService
}

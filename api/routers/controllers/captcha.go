package controllers

import (
	"admin-api/api/service"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

// 验证码
// @Summary 验证码接口
// @Produce json
// @Description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptMake()
	result.Success(c, map[string]interface{}{"idKey": id, "image": base64Image})
}

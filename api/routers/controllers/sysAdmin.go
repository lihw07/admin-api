package controllers

import (
	"admin-api/api/model"
	"admin-api/api/service"
	"github.com/gin-gonic/gin"
)

// Login
// @Summary 用户登录接口
// @Produce json
// @Description 用户登录接口
// @Param data body model.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
func Login(c *gin.Context) {
	var dto model.LoginDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().Login(c, dto)
}

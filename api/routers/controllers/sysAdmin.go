package controllers

import (
	"admin-api/api/dao"
	"admin-api/api/model"
	"admin-api/api/service"
	"admin-api/common/result"
	"admin-api/pkg/jwt"
	"admin-api/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

type SysAdminServiceImpl struct{}

// 登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto model.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}
	// 验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired),
			result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}

	// 校验验证码
	verifyRes := service.CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		result.Failed(c, int(result.ApiCode.CAPTCHANOTTRUE),
			result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}

	// 校验
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE),
			result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}

	const status int = 2
	if sysAdmin.Status == status {
		result.Failed(c, int(result.ApiCode.STATUSISENABLE),
			result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}
	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	// 左侧菜单列表
	var leftMenuVo []model.LeftMenuVo
	leftMenuList := dao.QueryLeftMenuList(sysAdmin.ID)
	for _, value := range leftMenuList {
		menuSvoList := dao.QueryMenuVoList(sysAdmin.ID, value.Id)
		item := model.LeftMenuVo{}
		item.MenuSvoList = menuSvoList
		item.Id = value.Id
		item.MenuName = value.MenuName
		item.Icon = value.Icon
		item.Url = value.Url
		leftMenuVo = append(leftMenuVo, item)
	}

	// 权限列表
	permissionList := dao.QueryPermissionList(sysAdmin.ID)
	var stringList = make([]string, 0)
	for _, value := range permissionList {
		stringList = append(stringList, value.Value)
	}
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdmin, "leftMenuList": leftMenuVo, "permissionList": stringList})
}

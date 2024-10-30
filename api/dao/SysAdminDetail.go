package dao

import (
	"admin-api/api/model"
	. "admin-api/pkg/db"
)

// SysAdminDetail 用户详情
func SysAdminDetail(dto model.LoginDto) (sysAdmin model.SysAdmin) {
	username := dto.Username
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

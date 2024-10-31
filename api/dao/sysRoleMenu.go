package dao

import (
	"admin-api/api/model"
	. "admin-api/pkg/db"
	"admin-api/pkg/util"
	"time"
)

// 根据菜单名称查询
func GetSysMenuByName(menuName string) (sysMenu model.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}

// 查询新增选项列表
func QuerySysMenuVoList() (sysMenuVo []model.SysMenuVo) {
	Db.Table("sys_menu").Select("id, menu_name AS label,parent_id").Scan(&sysMenuVo)
	return sysMenuVo
}

// 新增菜单
func CreateSysMenu(addSysMenu model.SysMenu) bool {
	sysMenuByName := GetSysMenuByName(addSysMenu.MenuName)
	if sysMenuByName.ID != 0 {
		return false
	}
	// 目录
	if addSysMenu.MenuType == 1 {
		sysMenu := model.SysMenu{
			ParentId:   0,
			MenuName:   addSysMenu.MenuName,
			Icon:       addSysMenu.Icon,
			MenuType:   addSysMenu.MenuType,
			Url:        addSysMenu.Url,
			MenuStatus: addSysMenu.MenuStatus,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 2 {
		sysMenu := model.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			Icon:       addSysMenu.Icon,
			MenuType:   addSysMenu.MenuType,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			Url:        addSysMenu.Url,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 3 {
		sysMenu := model.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			MenuType:   addSysMenu.MenuType,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	}
	return false
}

// 根据id查询菜单详情
func GetSysMenu(Id int) (sysMenu model.SysMenu) {
	Db.First(&sysMenu, Id)
	return sysMenu
}

// 修改菜单
func UpdateSysMenu(menu model.SysMenu) (sysMenu model.SysMenu) {
	Db.First(&sysMenu, menu.ID)
	sysMenu.ParentId = menu.ParentId
	sysMenu.MenuName = menu.MenuName
	sysMenu.Icon = menu.Icon
	sysMenu.Value = menu.Value
	sysMenu.MenuType = menu.MenuType
	sysMenu.Url = menu.Url
	sysMenu.MenuStatus = menu.MenuStatus
	sysMenu.Sort = menu.Sort
	Db.Save(&sysMenu)
	return sysMenu
}
func GetSysRoleMenu(id uint) (sysRoleMenu model.SysRoleMenu) {
	Db.Where("menu_id = ?", id).First(&sysRoleMenu)
	return sysRoleMenu
}

// 删除菜单
func DeleteSysMenu(dto model.SysMenuIdDto) bool {
	// 菜单已分配角色不能删除
	sysRoleMenu := GetSysRoleMenu(dto.Id)
	if sysRoleMenu.MenuId > 0 {
		return false
	}
	Db.Where("parent_id = ?", dto.Id).Delete(&model.SysMenu{})
	Db.Delete(&model.SysMenu{}, dto.Id)
	return true
}

// 查询菜单列表
func GetSysMenuList(MenuName string, MenuStatus string) (sysMenu []*model.SysMenu) {
	curDb := Db.Table("sys_menu").Order("sort")
	if MenuName != "" {
		curDb = curDb.Where("menu_name = ?", MenuName)
	}
	if MenuStatus != "" {
		curDb = curDb.Where("menu_status = ?", MenuStatus)
	}
	curDb.Find(&sysMenu)
	return sysMenu
}

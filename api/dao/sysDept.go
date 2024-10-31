package dao

import (
	"admin-api/api/model"
	. "admin-api/pkg/db"
	"admin-api/pkg/util"
	"time"
)

// 根据部门名称查询
func GetSysDeptByName(deptName string) (sysDept model.SysDept) {
	Db.Where("dept_name = ?", deptName).First(&sysDept)
	return sysDept
}

// 新增部门
func CreateSysDept(sysDept model.SysDept) bool {
	sysDeptByName := GetSysDeptByName(sysDept.DeptName)
	if sysDeptByName.ID > 0 {
		return false
	}
	if sysDept.DeptType == 1 {
		sysDept := model.SysDept{
			DeptStatus: sysDept.DeptStatus,
			ParentId:   0,
			DeptType:   sysDept.DeptType,
			DeptName:   sysDept.DeptName,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysDept)
		return true
	} else {
		sysDept := model.SysDept{
			DeptStatus: sysDept.DeptStatus,
			ParentId:   sysDept.ParentId,
			DeptType:   sysDept.DeptType,
			DeptName:   sysDept.DeptName,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysDept)
		return true
	}
	return false
}

// 部门下拉列表
func QuerySysDeptVoList() (sysDeptVo []model.SysDeptVo) {
	Db.Table("sys_dept").Select("id, dept_name AS label,parent_id").Scan(&sysDeptVo)
	return sysDeptVo
}

// 根据id查询部门
func GetSysDeptById(Id int) (sysDept model.SysDept) {
	Db.First(&sysDept, Id)
	return sysDept
}

// 修改部门
func UpdateSysDept(dept model.SysDept) (sysDept model.SysDept) {
	Db.First(&sysDept, dept.ID)
	sysDept.ParentId = dept.ParentId
	sysDept.DeptType = dept.DeptType
	sysDept.DeptName = dept.DeptName
	sysDept.DeptStatus = dept.DeptStatus
	Db.Save(&sysDept)
	return sysDept
}

// 查询部门是否有人
func GetSysAdminDept(id int) (sysAdmin model.SysAdmin) {
	Db.Where("dept_id = ?", id).First(&sysAdmin)
	return sysAdmin
}

// 删除部门
func DeleteSysDeptById(dto model.SysDeptIdDto) bool {
	sysAdmin := GetSysAdminDept(dto.Id)
	if sysAdmin.ID > 0 {
		return false
	}
	Db.Where("parent_id = ?", dto.Id).Delete(&model.SysDept{})
	Db.Delete(&model.SysDept{}, dto.Id)
	return true
}

// 查询部门列表
func GetSysDeptList(DeptName string, DeptStatus string) (sysDept []model.SysDept) {
	curDb := Db.Table("sys_dept")
	if DeptName != "" {
		curDb = curDb.Where("dept_name = ?", DeptName)
	}
	if DeptStatus != "" {
		curDb = curDb.Where("dept_status = ?", DeptStatus)
	}
	curDb.Find(&sysDept)
	return sysDept
}

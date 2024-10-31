package service

import (
	"admin-api/api/dao"
	"admin-api/api/model"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

// 接口定义
type ISysDeptService interface {
	CreateSysDept(c *gin.Context, sysDept model.SysDept)
	QuerySysDeptVoList(c *gin.Context)
	GetSysDeptById(c *gin.Context, Id int)
	UpdateSysDept(c *gin.Context, dept model.SysDept)
	DeleteSysDeptById(c *gin.Context, dto model.SysDeptIdDto)
	GetSysDeptList(c *gin.Context, DeptName string, DeptStatus string)
}
type SysDeptServiceImpl struct{}

// 查询部门列表
func (s SysDeptServiceImpl) GetSysDeptList(c *gin.Context, DeptName string,
	DeptStatus string) {
	result.Success(c, dao.GetSysDeptList(DeptName, DeptStatus))
}

// 删除部门
func (s SysDeptServiceImpl) DeleteSysDeptById(c *gin.Context, dto model.SysDeptIdDto) {
	bool := dao.DeleteSysDeptById(dto)
	if !bool {
		result.Failed(c, int(result.ApiCode.DEPTISDISTRIBUTE),
			result.ApiCode.GetMessage(result.ApiCode.DEPTISDISTRIBUTE))
		return
	}
	result.Success(c, true)
}

// 修改部门
func (s SysDeptServiceImpl) UpdateSysDept(c *gin.Context, dept model.SysDept) {
	sysDept := dao.UpdateSysDept(dept)
	result.Success(c, sysDept)
}

// 根据id查询部门
func (s SysDeptServiceImpl) GetSysDeptById(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysDeptById(Id))
}

// 部门下拉列表
func (s SysDeptServiceImpl) QuerySysDeptVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysDeptVoList())
}

// 新增部门
func (s SysDeptServiceImpl) CreateSysDept(c *gin.Context, sysDept model.SysDept) {
	bool := dao.CreateSysDept(sysDept)
	if !bool {
		result.Failed(c, int(result.ApiCode.DEPTISEXIST),
			result.ApiCode.GetMessage(result.ApiCode.DEPTISEXIST))
		return
	}
	result.Success(c, true)
}

var sysDeptService = SysDeptServiceImpl{}

func SysDeptService() ISysDeptService {
	return &sysDeptService
}

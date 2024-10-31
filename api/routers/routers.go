package routers

import (
	"admin-api/api/routers/controllers"
	"admin-api/common/config"
	"admin-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 跌机是恢复
	router.Use(gin.Recovery())
	// 跨域中间件
	router.Use(middleware.Cors())
	// 图片访问路径静态文件夹可直接访问
	router.StaticFS(config.ImageSetting.UploadDir,
		http.Dir(config.ImageSetting.UploadDir))
	// 日志log中间件
	router.Use(middleware.Logger())
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {

	router.GET("/api/captcha", controllers.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/login", controllers.Login)
	// jwt鉴权接口
	jwt := router.Group("/api", middleware.AuthMiddleware())
	{
		jwt.POST("/post/add", controllers.CreateSysPost)
		jwt.GET("/post/info", controllers.GetSysPostById)
		jwt.PUT("/post/update", controllers.UpdateSysPost)
		jwt.DELETE("/post/delete", controllers.DeleteSysPostById)
		jwt.DELETE("/post/batch/delete", controllers.BatchDeleteSysPost)
		jwt.PUT("/post/updateStatus", controllers.UpdateSysPostStatus)
		jwt.GET("/post/list", controllers.GetSysPostList)
		jwt.GET("/post/vo/list", controllers.QuerySysPostVoList)
		jwt.POST("/upload", controllers.Upload)
		jwt.GET("/api/menu/vo/list", controllers.QuerySysMenuVoList)
		jwt.POST("/api/menu/add", controllers.CreateSysMenu)
		jwt.GET("/api/menu/info", controllers.GetSysMenu)
		jwt.PUT("/api/menu/update", controllers.UpdateSysMenu)
		jwt.DELETE("/api/menu/delete", controllers.DeleteSysRoleMenu)
		jwt.GET("/api/menu/list", controllers.GetSysMenuList)
		jwt.POST("/api/role/add", controllers.CreateSysRole)
		jwt.PUT("/api/role/info", controllers.GetSysRoleById)
		jwt.PUT("/api/role/update", controllers.UpdateSysRole)
		jwt.DELETE("/api/role/delete", controllers.DeleteSysRoleById)
		jwt.PUT("/api/role/updateStatus", controllers.UpdateSysRoleStatus)
		jwt.GET("/api/role/list", controllers.GetSysRoleList)
		jwt.GET("/api/role/vo/list", controllers.QuerySysRoleVoList)
		jwt.GET("/api/role/vo/idList", controllers.QueryRoleMenuIdList)
		jwt.PUT("/api/role/assignPermissions", controllers.AssignPermissions)
	}
}

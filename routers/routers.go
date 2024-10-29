package routers

import (
	"admin-api/common/config"
	"admin-api/middleware"
	"github.com/gin-gonic/gin"
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
	// todo 后续接口url
}

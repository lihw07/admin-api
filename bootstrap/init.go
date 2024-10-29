package bootstrap

import (
	"admin-api/common/config"
	"admin-api/pkg/db"
	log "admin-api/pkg/log"
	"admin-api/pkg/redis"
	"github.com/gin-gonic/gin"
)

// Init 初始化启动
func Init(path string) {

	// 加载日志
	logger := log.Log()

	config.Init(path)
	// Debug 关闭时，切换为生产模式
	if config.SystemConfig.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化配置文件
	dependencies := []struct {
		mode    string
		factory func()
	}{
		{
			"both",
			func() {
				err := db.SetupDBLink()
				if err != nil {
					logger.Info("database config error")
				}
			},
		},
		{
			"both",
			func() {
				err := redis.SetupRedisDb()
				if err != nil {
					logger.Info("redis config error")
				}
			},
		},
	}

	for _, dependency := range dependencies {
		if dependency.mode == config.SystemConfig.Mode || dependency.mode == "both" {
			dependency.factory()
		}
	}
}

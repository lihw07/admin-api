package main

import (
	"admin-api/bootstrap"
	"admin-api/common/config"
	"admin-api/pkg/log"
	"admin-api/pkg/util"
	"admin-api/routers"
	"context"
	"flag"
	"github.com/cloudreve/Cloudreve/v3/pkg/conf"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	confPath   string
	scriptName string
)

// 初始化配置
func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "Path to the config file.")
	flag.StringVar(&scriptName, "database-script", "", "Name of database util script.")
	flag.Parse()

	bootstrap.Init(confPath)
}

func main() {
	// 加载日志
	logger := log.Log()
	api := routers.InitRouter()
	api.TrustedPlatform = conf.SystemConfig.ProxyHeader
	server := &http.Server{Handler: api}

	// 收到信号后关闭服务器
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	go shutdown(sigChan, server)

	defer func() {
		<-sigChan
	}()

	logger.Info("Listening to :", config.SystemConfig.Listen)
	server.Addr = config.SystemConfig.Listen
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Failed to listen to %q: %s", config.SystemConfig.Listen, err)
	}

	log.Log().Info("Listening to %q", conf.SystemConfig.Listen)
	server.Addr = conf.SystemConfig.Listen
	if err := server.ListenAndServe(); err != nil {
		log.Log().Error("Failed to listen to %q: %s", conf.SystemConfig.Listen, err)
	}

}

func shutdown(sigChan chan os.Signal, server *http.Server) {
	logger := log.Log()
	sig := <-sigChan
	logger.Info("Signal %s received, shutting down server...", sig)
	ctx := context.Background()
	if config.SystemConfig.GracePeriod != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(config.SystemConfig.GracePeriod)*time.Second)
		defer cancel()
	}

	// Shutdown http server
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Error("Failed to shutdown server: %s", err)
	}

	close(sigChan)
}

package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Meedu/gin-template/global"
	"github.com/Meedu/gin-template/initialize"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.MD_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.MD_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.MD_LOG.Info("server run success on ", zap.String("address", address))

	global.MD_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

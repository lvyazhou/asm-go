package server

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/interfaces/routers"
	"asm_platform/interfaces/swag"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	routers.InitRouter(r)

	conf := config.GetConfig()
	slog.Infof("start server asm-platform-api %s-%s on %d success...\n", conf.GetString("server.name"), conf.GetString("server.version"), conf.GetInt("server.port"))

	// swag初始化
	swag.Init(r, conf)

	// 服务启动1
	listenAddr := conf.GetString("server.listen")
	if err := r.Run(listenAddr); err != nil {
		slog.Panicf("error on listening to %s: %s\n", listenAddr, err)
	}
}

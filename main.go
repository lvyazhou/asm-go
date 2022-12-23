package main

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/database/es"
	"asm_platform/infrastructure/pkg/database/kafka"
	mgo "asm_platform/infrastructure/pkg/database/mongo"
	db_mysql "asm_platform/infrastructure/pkg/database/mysql"
	cache "asm_platform/infrastructure/pkg/database/redis"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/infrastructure/pkg/task"
	"asm_platform/interfaces/server"
	"flag"
	"fmt"
	"os"
)

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        ASM-TOKEN
func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: asm-platform-web -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	// 配置文件加载
	config.Init(*environment)

	// 日志文件加载
	slog.Init()

	// Mysql加载
	db_mysql.Init()
	defer db_mysql.Close()

	// redis加载
	cache.Init()
	defer cache.Close()

	// es加载
	es.Init()

	// kafka加载
	kfk.Init()

	// mgo 加载
	mgo.Init()

	// 加载task queue
	task.Init()

	// 服务启动
	server.Init()
}

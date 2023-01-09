package main

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/infrastructure/repo"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

	// 生产数据
	productMsg()
}

func productMsg() {
	fmt.Println(" ---- start product msg ----")
	// 发送kafka test
	r := repo.NewKafkaRepo("lyz", "")
	m := map[string]string{}
	key := strconv.FormatInt(time.Now().Unix(), 10)
	m[key] = "lyz888888"
	r.WriteKafkaMessage(m)
	slog.Info(" ---- end product msg ----")
}

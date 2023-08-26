package main

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/slog"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"asm_platform/infrastructure/repo"
	"flag"
	"fmt"
	"os"
	"strconv"
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
func sendEvent(topic string, str string) {
	fmt.Printf("send topic name %v start ... \n", topic)
	// 发送kafka test
	r := repo.NewKafkaRepo(topic, "")
	var ms []map[string]string
	for i := 0; i < 1; i++ {
		m := map[string]string{}
		key := strconv.FormatInt(utils_tool.GenerateUniqueId(), 10)
		m[key] = str
		ms = append(ms, m)
	}
	r.WriteKafkaMessageList(ms)
	fmt.Printf("send topic name %v finished ... \n", topic)
}

func productMsg() {
	fmt.Println(" ---- start product msg ----")
	// 发送kafka test
	r := repo.NewKafkaRepo("mss-edr-log", "")
	var ms []map[string]string

	for i := 0; i < 111; i++ {
		m := map[string]string{}
		key := strconv.FormatInt(utils_tool.GenerateUniqueId(), 10)

		m[key] = ""
		ms = append(ms, m)
	}
	slog.Infof("send kafka size %v ...", len(ms))
	r.WriteKafkaMessageList(ms)
	slog.Info(" ---- end product msg ----")
}

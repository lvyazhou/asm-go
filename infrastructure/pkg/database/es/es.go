package es

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/slog"
	"github.com/olivere/elastic/v7"
	"log"
)

var Client *elastic.Client

// Init es init
func Init() {
	if err := newEsDriver(); err != nil {
		slog.Panicf("elasticSearch error on database initialization: %s\n", err)
		return
	}
}

// es加载驱动
func newEsDriver() error {
	conf := config.GetConfig()
	url := conf.GetString("es.addr")
	user := conf.GetString("es.user")
	pwd := conf.GetString("es.pwd")
	sniff := conf.GetBool("es.sniffer_enabled")
	health := conf.GetBool("es.health_check")

	c, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(user, pwd),
		elastic.SetHealthcheck(health),
		elastic.SetSniff(sniff),
		// 设置错误日志
		//elastic.SetErrorLog(log.New(os.Stderr, "ES-ERROR ", log.LstdFlags)),
		// 设置Info日志
		//elastic.SetInfoLog(log.New(os.Stdout, "ES-INFO ", log.LstdFlags)),
		//elastic.SetTraceLog(log.New(os.Stdout, "ES-Trace-Log ", log.LstdFlags)),
	)
	if err != nil {
		log.Fatalln("Failed to create elasticSearch client")
		return err
	}

	slog.Infof("elasticSearch Connected...")
	Client = c
	return nil
}

package main

import (
	"asm_platform/infrastructure/config"
	kfk "asm_platform/infrastructure/pkg/database/kafka"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// 执行kafka consumer
	slog.Info("----start consumer msg ----")
	go consumeGroupMsg()

	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// recivie signal to exit main goroutine
	//window signal
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	// Block until we receive our signal.
	<-ch

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// if your application should wait for other services
	// to finalize based on context cancellation.
	<-ctx.Done()

	log.Println("shutting down")
}

func consumeGroupMsg() {
	consumer := kfk.KfkConsumer{
		Topic:           "lyz",
		EnablePartition: false,
		Partition:       0,
		GroupID:         "lyz_group888",
	}
	r := consumer.NewKfkConsumer()
	defer r.Close()

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}

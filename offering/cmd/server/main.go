package main

import (
	"context"
	"flag"
	"fmt"
	"offering/internal/app"
	"offering/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "configs/server/.config.json", "set config path")
	flag.Parse()

	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		fmt.Println(fmt.Errorf("fatal: init config %v", err))
		os.Exit(1)
	}

	a := app.NewApp(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	a.Run()

	<-ctx.Done()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	a.Stop(ctx)
}

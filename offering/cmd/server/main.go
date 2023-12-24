package main

import (
	"context"
	"offering/internal/app"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	a := app.NewApp(":8080")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	a.Run()

	<-ctx.Done()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	a.Stop(ctx)
}

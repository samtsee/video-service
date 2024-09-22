package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/samtsee/video-service/internal/api"
)

func main() {
	app, err := api.New()
	if err != nil {
		fmt.Println("failed to initialize app: ", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: %w", err)
	}
}

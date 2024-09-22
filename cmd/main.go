package main

import (
	"context"
	"fmt"

	"github.com/samtsee/video-service/internal/api"
)

func main() {
	app := api.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app: %w", err)
	}
}

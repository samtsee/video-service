package main

import (
	"context"
	"fmt"

	"github.com/samtsee/video-service/internal/api"
)

func main() {
	app, err := api.New()
	if err != nil {
		fmt.Println("failed to initialize app: ", err)
		return
	}

	err = app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app: %w", err)
	}
}

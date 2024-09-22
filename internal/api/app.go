package api

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/samtsee/video-service/internal/repository"
)

type App struct {
	router http.Handler
	db     repository.PostRepository
}

func New() (*App, error) {
	db, err := repository.NewPostgresRepo()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	app := &App{
		router: loadRoutes(),
		db:     db,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

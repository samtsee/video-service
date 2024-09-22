package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

	defer func() {
		if err := a.db.Disconnect(); err != nil {
			fmt.Errorf("failed to close postgres connection: %w", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}

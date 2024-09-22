package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samtsee/video-service/internal/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/posts", loadPostRoutes)

	return router
}

func loadPostRoutes(router chi.Router) {
	postHandler := &handler.Post{}

	router.Post("/", postHandler.Create)
	router.Get("/", postHandler.List)
	router.Get("/{id}", postHandler.GetByID)
	router.Put("/{id}", postHandler.UpdateByID)
	router.Delete("/{id}", postHandler.DeleteByID)
}

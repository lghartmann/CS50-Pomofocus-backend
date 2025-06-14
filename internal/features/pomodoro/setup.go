package pomodoro

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/lghartmann/CS50-Pomofocus-backend/internal/middleware"
)

func SetupRoutesAndInjection(db *sql.DB, r *chi.Mux) {
	repo := NewPomodoroRepository(db)
	service := NewPomodoroService(repo)
	handler := NewPomodoroHandler(service)

	r.With(middleware.AuthMiddleware).Get("/", handler.Search)
	r.With(middleware.AuthMiddleware).Post("/", handler.Create)
	r.With(middleware.AuthMiddleware).Delete("/{id}", handler.Inactivate)
}

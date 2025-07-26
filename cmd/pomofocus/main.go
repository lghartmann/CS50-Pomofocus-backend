package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/lghartmann/CS50-Pomofocus-backend/internal/features/pomodoro"
	custommiddleware "github.com/lghartmann/CS50-Pomofocus-backend/internal/middleware"
)

type Config struct {
	DB_URI       string
	HTTP_ADDRESS string
}

var cfg Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load envs: %v", err)
	}

	custommiddleware.JwtSecret.Value = os.Getenv("JWT_SECRET")

	cfg.DB_URI = os.Getenv("DB_URI")
	cfg.HTTP_ADDRESS = os.Getenv("HTTP_ADDRESS")
}

func main() {
	db, err := sql.Open("postgres", cfg.DB_URI)
	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}

	r := chi.NewMux()

	setup(db, r)
	log.Printf("server running at http://localhost%s", cfg.HTTP_ADDRESS)

	if err := http.ListenAndServe(cfg.HTTP_ADDRESS, r); err != nil {
		log.Fatalf("server start failed: %v", err)
	}
}

func setup(db *sql.DB, r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/pomodoro", func(r chi.Router) {
			pomodoro.SetupRoutesAndInjection(db, r)
		})
	})
}

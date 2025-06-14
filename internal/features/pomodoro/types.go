package pomodoro

import (
	"context"
	"net/http"
	"time"

	endpointtypes "github.com/lghartmann/CS50-Pomofocus-backend/pkg/types"
)

type IPomodoroRepository interface {
	GetById(id string, ctx context.Context) (PomodoroDto, error)
	Create(dto PomodoroRepositoryCreateDto, ctx context.Context) error
	Search(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error)
	Inactivate(id string, ctx context.Context) error
}

type IPomodoroService interface {
	Create(dto PomodoroCreateDto, ctx context.Context) error
	Search(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error)
	Inactivate(id string, ctx context.Context) error
}

type IPomodoroHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Inactivate(w http.ResponseWriter, r *http.Request)
}

type PomodoroCreateDto struct {
	Duration      string  `json:"duration"`
	PauseDuration string  `json:"pause_duration"`
	Effort        float32 `json:"effort"`
	Distraction   float32 `json:"distraction"`
	Productivity  float32 `json:"productivity"`
}

type PomodoroRepositoryCreateDto struct {
	UserID        string
	Duration      string
	PauseDuration string
	Effort        float32
	Distraction   float32
	Productivity  float32
}

type PomodoroDto struct {
	ID            string
	Duration      string
	PauseDuration string
	Effort        float32
	Distraction   float32
	Productivity  float32
	CreatedAt     time.Time
}

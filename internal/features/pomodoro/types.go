package pomodoro

import (
	"context"
	"net/http"
	"time"

	endpointtypes "github.com/lghartmann/CS50-Pomofocus-backend/pkg/types"
)

type IPomodoroRepository interface {
	GetById(id string, ctx context.Context) (PomodoroDto, error)
	SearchDashboard(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error)
	Search(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error)
	Create(dto PomodoroRepositoryCreateDto, ctx context.Context) error
	Update(id string, dto PomodoroUpdateDto, ctx context.Context) error
	Inactivate(id string, ctx context.Context) error
}

type IPomodoroService interface {
	GetById(id string, ctx context.Context) (PomodoroDto, error)
	SearchDashboard(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error)
	Search(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error)
	Create(dto PomodoroCreateDto, ctx context.Context) error
	Update(id string, dto PomodoroUpdateDto, ctx context.Context) error
	Inactivate(id string, ctx context.Context) error
}

type IPomodoroHandler interface {
	GetById(w http.ResponseWriter, r *http.Request)
	SearchDashboard(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Inactivate(w http.ResponseWriter, r *http.Request)
}

type PomodoroCreateDto struct {
	Duration      string  `json:"duration"`
	PauseDuration string  `json:"pause_duration"`
	Effort        float32 `json:"effort"`
	Distraction   float32 `json:"distraction"`
	Productivity  float32 `json:"productivity"`
}

type PomodoroUpdateDto struct {
	Duration      *string  `json:"duration,omitempty"`
	PauseDuration *string  `json:"pause_duration,omitempty"`
	Effort        *float32 `json:"effort,omitempty"`
	Distraction   *float32 `json:"distraction,omitempty"`
	Productivity  *float32 `json:"productivity,omitempty"`
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
	ID            string     `json:"id"`
	Duration      string     `json:"duration"`
	PauseDuration string     `json:"pauseDuration"`
	Effort        float32    `json:"effort"`
	Distraction   float32    `json:"distraction"`
	Productivity  float32    `json:"productivity"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

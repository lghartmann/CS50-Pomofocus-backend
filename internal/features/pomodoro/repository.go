package pomodoro

import (
	"context"
	"database/sql"
	"time"

	"github.com/lghartmann/CS50-Pomofocus-backend/internal/middleware"
	"github.com/lghartmann/CS50-Pomofocus-backend/pkg"
	endpointtypes "github.com/lghartmann/CS50-Pomofocus-backend/pkg/types"
)

type PomodoroRepository struct {
	db *sql.DB
}

func NewPomodoroRepository(db *sql.DB) IPomodoroRepository {
	return &PomodoroRepository{db: db}
}

func (p *PomodoroRepository) GetById(id string, ctx context.Context) (PomodoroDto, error) {
	userId, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return PomodoroDto{}, sql.ErrNoRows
	}

	query := "SELECT id, duration, pause_duration, effort, distraction, productivity, created_at FROM pomodoro WHERE user_id = $1 AND id = $2 AND deleted_at IS NULL LIMIT 1;"

	var pomo PomodoroDto

	err := p.db.QueryRowContext(ctx, query, userId, id).Scan(
		&pomo.ID,
		&pomo.Duration,
		&pomo.PauseDuration,
		&pomo.Effort,
		&pomo.Distraction,
		&pomo.Productivity,
		&pomo.CreatedAt,
	)

	return pomo, err
}

func (p *PomodoroRepository) Search(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error) {
	start := pkg.GetStartFromOptions(ctx)
	offset := pkg.GetOffsetFromOptions(ctx)
	userId, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return endpointtypes.SearchResponse[PomodoroDto]{}, sql.ErrNoRows
	}

	query := "SELECT id, duration, pause_duration, effort, distraction, productivity, created_at FROM pomodoro WHERE user_id = $1 AND deleted_at IS NULL OFFSET $2 LIMIT $3;"

	var res []PomodoroDto

	rows, err := p.db.QueryContext(ctx, query, userId, start, offset)
	if err != nil {
		return endpointtypes.SearchResponse[PomodoroDto]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var dto PomodoroDto
		err := rows.Scan(
			&dto.ID,
			&dto.Duration,
			&dto.PauseDuration,
			&dto.Effort,
			&dto.Distraction,
			&dto.Productivity,
			&dto.CreatedAt,
		)
		if err != nil {
			return endpointtypes.SearchResponse[PomodoroDto]{}, err
		}
		res = append(res, dto)
	}

	return endpointtypes.SearchResponse[PomodoroDto]{Data: res}, nil
}

func (p *PomodoroRepository) Create(dto PomodoroRepositoryCreateDto, ctx context.Context) error {
	query := "INSERT INTO pomodoro (user_id, duration, pause_duration, effort, distraction, productivity) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := p.db.Exec(query, dto.UserID, dto.Duration, dto.PauseDuration, dto.Effort, dto.Distraction, dto.Productivity)

	return err
}

func (p *PomodoroRepository) Inactivate(id string, ctx context.Context) error {
	userId, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return sql.ErrNoRows
	}

	date := time.Now()

	query := "UPDATE pomodoro SET deleted_at = $1 WHERE deleted_at IS NULL and user_id = $2;"

	_, err := p.db.Exec(query, date, userId)

	return err
}

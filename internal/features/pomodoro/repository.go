package pomodoro

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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

	query := "SELECT id, duration, pause_duration, effort, distraction, productivity, created_at, updated_at FROM pomodoro WHERE user_id = $1 AND id = $2 AND deleted_at IS NULL LIMIT 1;"

	var pomo PomodoroDto

	err := p.db.QueryRowContext(ctx, query, userId, id).Scan(
		&pomo.ID,
		&pomo.Duration,
		&pomo.PauseDuration,
		&pomo.Effort,
		&pomo.Distraction,
		&pomo.Productivity,
		&pomo.CreatedAt,
		&pomo.UpdatedAt,
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

	countQuery := "SELECT COUNT(*) FROM pomodoro WHERE user_id = $1 AND deleted_at IS NULL;"
	var totalCount int
	err := p.db.QueryRowContext(ctx, countQuery, userId).Scan(&totalCount)
	if err != nil {
		return endpointtypes.SearchResponse[PomodoroDto]{}, err
	}

	query := "SELECT id, duration, pause_duration, effort, distraction, productivity, created_at, updated_at FROM pomodoro WHERE user_id = $1 AND deleted_at IS NULL OFFSET $2 LIMIT $3;"

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
			&dto.UpdatedAt,
		)
		if err != nil {
			return endpointtypes.SearchResponse[PomodoroDto]{}, err
		}
		res = append(res, dto)
	}

	return endpointtypes.SearchResponse[PomodoroDto]{Data: res, Count: totalCount}, nil
}

func (p *PomodoroRepository) Create(dto PomodoroRepositoryCreateDto, ctx context.Context) error {
	query := "INSERT INTO pomodoro (user_id, duration, pause_duration, effort, distraction, productivity) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := p.db.Exec(query, dto.UserID, dto.Duration, dto.PauseDuration, dto.Effort, dto.Distraction, dto.Productivity)

	return err
}

func (p *PomodoroRepository) Update(id string, dto PomodoroUpdateDto, ctx context.Context) error {
	userId, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return sql.ErrNoRows
	}

	var setParts []string
	var args []any
	argIndex := 1

	if dto.Duration != nil {
		setParts = append(setParts, fmt.Sprintf("duration = $%d", argIndex))
		args = append(args, *dto.Duration)
		argIndex++
	}
	if dto.PauseDuration != nil {
		setParts = append(setParts, fmt.Sprintf("pause_duration = $%d", argIndex))
		args = append(args, *dto.PauseDuration)
		argIndex++
	}
	if dto.Effort != nil {
		setParts = append(setParts, fmt.Sprintf("effort = $%d", argIndex))
		args = append(args, *dto.Effort)
		argIndex++
	}
	if dto.Distraction != nil {
		setParts = append(setParts, fmt.Sprintf("distraction = $%d", argIndex))
		args = append(args, *dto.Distraction)
		argIndex++
	}
	if dto.Productivity != nil {
		setParts = append(setParts, fmt.Sprintf("productivity = $%d", argIndex))
		args = append(args, *dto.Productivity)
		argIndex++
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++

	args = append(args, userId, id)

	query := fmt.Sprintf("UPDATE pomodoro SET %s WHERE user_id = $%d AND id = $%d AND deleted_at IS NULL;",
		strings.Join(setParts, ", "), argIndex, argIndex+1)

	_, err := p.db.ExecContext(ctx, query, args...)
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

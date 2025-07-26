package pomodoro

import (
	"context"
	"fmt"

	"github.com/lghartmann/CS50-Pomofocus-backend/internal/middleware"
	endpointtypes "github.com/lghartmann/CS50-Pomofocus-backend/pkg/types"
)

type PomodoroService struct {
	repository IPomodoroRepository
}

func NewPomodoroService(repository IPomodoroRepository) IPomodoroService {
	return &PomodoroService{repository: repository}
}

func (p *PomodoroService) GetById(id string, ctx context.Context) (PomodoroDto, error) {
	pomodoro, err := p.repository.GetById(id, ctx)
	if err != nil {
		return PomodoroDto{}, fmt.Errorf("pomodoro not found: %v", err)
	}

	return pomodoro, nil
}

func (p *PomodoroService) SearchDashboard(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error) {
	return p.repository.SearchDashboard(ctx)
}

func (p *PomodoroService) Search(ctx context.Context) (endpointtypes.SearchResponse[PomodoroDto], error) {
	return p.repository.Search(ctx)
}

func (p *PomodoroService) Create(dto PomodoroCreateDto, ctx context.Context) error {
	userId, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("user ID not found in context")
	}

	err := p.repository.Create(PomodoroRepositoryCreateDto{
		UserID:        userId,
		Duration:      dto.Duration,
		PauseDuration: dto.PauseDuration,
		Effort:        dto.Effort,
		Distraction:   dto.Distraction,
		Productivity:  dto.Productivity,
	}, ctx)
	if err != nil {
		return fmt.Errorf("error inserting pomodoro in database: %v", err)
	}

	return nil
}

func (p *PomodoroService) Update(id string, dto PomodoroUpdateDto, ctx context.Context) error {
	_, err := p.GetById(id, ctx)
	if err != nil {
		return err
	}

	err = p.repository.Update(id, dto, ctx)
	if err != nil {
		return fmt.Errorf("error updating pomodoro in database: %v", err)
	}

	return nil
}

func (p *PomodoroService) Inactivate(id string, ctx context.Context) error {
	_, err := p.repository.GetById(id, ctx)
	if err != nil {
		return fmt.Errorf("no available pomodoro to delete: %v", err)
	}

	err = p.repository.Inactivate(id, ctx)
	if err != nil {
		return fmt.Errorf("unable to delete pomodoro: %v", err)
	}

	return nil
}

package pomodoro

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PomodoroHandler struct {
	service IPomodoroService
}

func NewPomodoroHandler(service IPomodoroService) IPomodoroHandler {
	return &PomodoroHandler{service: service}
}

func (p *PomodoroHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto PomodoroCreateDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := p.service.Create(dto, r.Context())
	if err != nil {
		http.Error(w, fmt.Errorf("unable to create pomodoro: %v", err).Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *PomodoroHandler) Inactivate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := p.service.Inactivate(id, r.Context())
	if err != nil {
		http.Error(w, fmt.Errorf("unable to delete pomodoro: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (p *PomodoroHandler) Search(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	offsetStr := r.URL.Query().Get("offset")

	start, _ := strconv.Atoi(startStr)
	offset, _ := strconv.Atoi(offsetStr)

	ctx := context.WithValue(r.Context(), "start", start)
	ctx = context.WithValue(ctx, "offset", offset)

	res, err := p.service.Search(ctx)
	if err != nil {
		http.Error(w, "error searching for pomodoros", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

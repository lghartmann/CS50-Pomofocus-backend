package pkg

import (
	"context"
	"time"
)

const start string = "start"
const offset string = "offset"
const start_date string = "start_date"
const end_date string = "end_date"

func GetStartFromOptions(ctx context.Context) int {
	s := ctx.Value(start)
	if v, ok := s.(int); ok {
		return v
	}
	return 0
}

func GetOffsetFromOptions(ctx context.Context) int {
	s := ctx.Value(offset)
	if v, ok := s.(int); ok {
		return v
	}
	return 10
}

func GetStartDateFromOptions(ctx context.Context) string {
	fallBack := time.Now().AddDate(0, 0, -7).Format(time.RFC3339)
	s := ctx.Value(start_date)
	str, ok := s.(string)
	if !ok {
		return fallBack
	}

	if v, err := time.Parse(time.RFC3339, str); err == nil {
		return v.Format(time.RFC3339)
	}

	return fallBack
}

func GetEndDateFromOptions(ctx context.Context) string {
	fallBack := time.Now().Format(time.RFC3339)
	s := ctx.Value(end_date)
	str, ok := s.(string)
	if !ok {
		return fallBack
	}

	if v, err := time.Parse(time.RFC3339, str); err == nil {
		return v.Format(time.RFC3339)
	}

	return fallBack
}

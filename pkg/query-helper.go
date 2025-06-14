package pkg

import (
	"context"
)

const start string = "start"
const offset string = "offset"

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

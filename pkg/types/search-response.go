package endpointtypes

type SearchResponse[T any] struct {
	Data  []T
	Count int
}

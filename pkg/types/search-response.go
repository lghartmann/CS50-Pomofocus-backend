package endpointtypes

type SearchResponse[T any] struct {
	Data  []T `json:"data"`
	Count int `json:"count"`
}

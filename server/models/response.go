package models

type ResponseBody[T any] struct {
	Error bool `json:"error"`
	Data  T    `json:"data" validate:"omitempty"`
}

type PaginatedResponse[T any] struct {
	Records []T `json:"records,omitempty"`
	Skip    int `json:"skip"`
	Take    int `json:"take"`
	Count   int `json:"count"`
}

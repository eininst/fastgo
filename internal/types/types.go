package types

type Page[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
	Extra any   `json:"extra"`
}

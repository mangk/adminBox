package util

func TypePtr[T any](v T) *T { return &v }

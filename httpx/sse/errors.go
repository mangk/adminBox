package sse

import "errors"

var (
	ErrAlreadyStarted = errors.New("sse stream already started")
)

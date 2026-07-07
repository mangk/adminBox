package sse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Event represents a single SSE event.
type Event struct {
	Event string      `json:"event,omitempty"`
	ID    string      `json:"id,omitempty"`
	Data  interface{} `json:"data"`
	Retry int         `json:"retry,omitempty"`
}

// Stream is a Server-Sent Events stream controller.
type Stream struct {
	ctx       *gin.Context
	writer    io.Writer
	flusher   http.Flusher
	started   bool
	mu        sync.Mutex
	closed    bool
	heartbeat *time.Ticker
	done      chan struct{}
}

// NewStream creates a new SSE stream for the given context.
func NewStream(ctx *gin.Context) *Stream {
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("X-Accel-Buffering", "no")

	return &Stream{
		ctx:     ctx,
		writer:  ctx.Writer,
		flusher: ctx.Writer.(http.Flusher),
		done:    make(chan struct{}),
	}
}

// Start marks the stream as started. Returns ErrAlreadyStarted if called more than once.
func (s *Stream) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.started {
		return ErrAlreadyStarted
	}
	s.started = true
	return nil
}

// Send writes an event to the stream.
func (s *Stream) Send(event Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("sse stream already closed")
	}

	if !s.started {
		if err := s.Start(); err != nil {
			return err
		}
	}

	var builder strings.Builder

	if event.ID != "" {
		builder.WriteString("id: " + event.ID + "\n")
	}

	if event.Event != "" {
		builder.WriteString("event: " + event.Event + "\n")
	}

	if event.Retry > 0 {
		builder.WriteString(fmt.Sprintf("retry: %d\n", event.Retry))
	}

	dataBytes, err := json.Marshal(event.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal sse data: %w", err)
	}

	dataStr := string(dataBytes)
	lines := strings.Split(dataStr, "\n")
	for _, line := range lines {
		builder.WriteString("data: " + line + "\n")
	}

	builder.WriteString("\n")

	if _, err := s.writer.Write([]byte(builder.String())); err != nil {
		return fmt.Errorf("failed to write sse data: %w", err)
	}

	s.flusher.Flush()
	return nil
}

// SendJSON sends data as a JSON SSE event of type "message".
func (s *Stream) SendJSON(data interface{}) error {
	return s.Send(Event{
		Event: "message",
		Data:  data,
	})
}

// SendEvent sends a typed event.
func (s *Stream) SendEvent(eventType string, data interface{}) error {
	return s.Send(Event{
		Event: eventType,
		Data:  data,
	})
}

// StartHeartbeat starts a background goroutine that sends keep-alive comments.
func (s *Stream) StartHeartbeat(interval time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.heartbeat != nil {
		s.heartbeat.Stop()
	}

	s.heartbeat = time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-s.heartbeat.C:
				s.mu.Lock()
				if !s.closed {
					s.writer.Write([]byte(": heartbeat\n\n"))
					s.flusher.Flush()
				}
				s.mu.Unlock()
			case <-s.done:
				return
			}
		}
	}()
}

// Close closes the stream and stops the heartbeat.
func (s *Stream) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil
	}

	s.closed = true
	if s.heartbeat != nil {
		s.heartbeat.Stop()
	}
	close(s.done)
	return nil
}

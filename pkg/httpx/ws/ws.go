package ws

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const BufSize = 1024

// Message represents a WebSocket message.
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
	ID   string      `json:"id,omitempty"`
	Time int64       `json:"time,omitempty"`
}

// Conn is a WebSocket connection wrapper.
type Conn struct {
	conn     *websocket.Conn
	mu       sync.Mutex
	closed   bool
	ctx      context.Context
	cancel   context.CancelFunc
	msgChan  chan *Message
	errChan  chan error
	pingChan chan struct{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  BufSize,
	WriteBufferSize: BufSize,
}

// NewConn creates a WebSocket connection by upgrading the HTTP request.
func NewConn(ctx *gin.Context) (*Conn, error) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade to websocket: %w", err)
	}

	wc := &Conn{
		conn:     conn,
		closed:   false,
		msgChan:  make(chan *Message, BufSize),
		errChan:  make(chan error, 1),
		pingChan: make(chan struct{}, 1),
	}

	conn.SetPongHandler(func(string) error {
		select {
		case wc.pingChan <- struct{}{}:
		default:
		}
		return nil
	})

	go wc.readLoop()

	return wc, nil
}

func (c *Conn) readLoop() {
	defer c.Close()

	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.errChan <- err
			}
			return
		}

		if msg.Type == "ping" {
			c.Send(&Message{
				Type: "pong",
				Time: time.Now().Unix(),
			})
			continue
		}

		select {
		case c.msgChan <- &msg:
		default:
			log.Printf("websocket message buffer full, dropping message")
		}
	}
}

// Send writes a message to the connection.
func (c *Conn) Send(msg *Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return ErrClosed
	}

	if msg.Time == 0 {
		msg.Time = time.Now().Unix()
	}

	return c.conn.WriteJSON(msg)
}

// SendJSON sends data as a "message"-typed event.
func (c *Conn) SendJSON(data interface{}) error {
	return c.Send(&Message{
		Type: "message",
		Data: data,
	})
}

// SendEvent sends a typed event.
func (c *Conn) SendEvent(eventType string, data interface{}) error {
	return c.Send(&Message{
		Type: eventType,
		Data: data,
	})
}

// SendError sends an error event.
func (c *Conn) SendError(errMsg string) error {
	return c.Send(&Message{
		Type: "error",
		Data: gin.H{"error": errMsg},
	})
}

// Messages returns a read-only channel of incoming messages.
func (c *Conn) Messages() <-chan *Message {
	return c.msgChan
}

// Close closes the connection.
func (c *Conn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}

	c.closed = true
	close(c.msgChan)
	close(c.errChan)
	close(c.pingChan)
	return c.conn.Close()
}

// IsClosed reports whether the connection is closed.
func (c *Conn) IsClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

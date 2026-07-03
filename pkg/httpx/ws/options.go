package ws

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mangk/adminBox/pkg/httpx"
)

// AsWebSocket 设置为 WebSocket 连接
func AsWebSocket(handler func(*Conn) error) httpx.RespOption {
	return func(r *httpx.Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			conn, err := NewConn(ctx)
			if err != nil {
				ctx.String(500, "WebSocket upgrade failed: %v", err)
				return
			}
			defer conn.Close()

			if err := handler(conn); err != nil {
				conn.SendError(err.Error())
			}
		})
	}
}

// AsWebSocketWithPing 带 Ping 的 WebSocket
func AsWebSocketWithPing(handler func(*Conn) error) httpx.RespOption {
	return func(r *httpx.Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			conn, err := NewConn(ctx)
			if err != nil {
				ctx.String(500, "WebSocket upgrade failed: %v", err)
				return
			}
			defer conn.Close()

			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()

			errChan := make(chan error, 1)

			go func() {
				errChan <- handler(conn)
			}()

			select {
			case err := <-errChan:
				if err != nil {
					conn.SendError(err.Error())
				}
			case <-ticker.C:
				conn.Send(&Message{Type: "ping"})
			}
		})
	}
}

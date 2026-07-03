package sse

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mangk/adminBox/pkg/httpx"
)

// AsSSE 设置为 SSE 流式响应
func AsSSE(handler func(*Stream) error) httpx.RespOption {
	return func(r *httpx.Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			stream := NewStream(ctx)
			defer stream.Close()

			stream.StartHeartbeat(30 * time.Second)

			if err := handler(stream); err != nil {
				stream.Send(Event{
					Event: "error",
					Data:  gin.H{"error": err.Error()},
				})
			}
		})
	}
}

// AsSSEEvent 发送单个或多个 SSE 事件
func AsSSEEvent(events ...Event) httpx.RespOption {
	return func(r *httpx.Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			stream := NewStream(ctx)
			defer stream.Close()

			for _, e := range events {
				if err := stream.Send(e); err != nil {
					return
				}
			}
		})
	}
}

// AsSSEWithHeartbeat 带心跳的 SSE
func AsSSEWithHeartbeat(heartbeat time.Duration, handler func(*Stream) error) httpx.RespOption {
	return func(r *httpx.Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			stream := NewStream(ctx)
			defer stream.Close()

			if heartbeat > 0 {
				stream.StartHeartbeat(heartbeat)
			}

			if err := handler(stream); err != nil {
				stream.Send(Event{
					Event: "error",
					Data:  gin.H{"error": err.Error()},
				})
			}
		})
	}
}

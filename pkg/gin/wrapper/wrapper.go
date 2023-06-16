package wrapper

import (
	"github.com/davveo/coupon-server/pkg/gin/code"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type MessageError struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

func Reply(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Message{
		Code: code.OK,
		Data: data,
	})
}

func ReplyErrWithStatusCode(ctx *gin.Context, statusCode, code int, hints ...string) {
	msg := &MessageError{
		Code: code,
	}
	if len(hints) > 0 {
		msg.Error = strings.Join(hints, ", ")
	}
	ctx.JSON(statusCode, msg)
}

func ReplyOK(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Message{
		Code: code.OK,
	})
}

func ReplyErr(ctx *gin.Context, code int, hints ...string) {
	msg := &MessageError{
		Code: code,
	}
	if len(hints) > 0 {
		msg.Error = strings.Join(hints, ", ")
	}
	ctx.JSON(http.StatusOK, msg)
}

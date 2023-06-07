package logger

import (
	"bytes"
	"encoding/json"
	"github.com/davveo/market-coupon/pkg/logger"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		if c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPost {
			// 请求体
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = c.GetRawData()
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				m := make(map[string]interface{})
				_ = json.Unmarshal(bodyBytes, &m)
				delete(m, "password")
				delete(m, "old_password")
				delete(m, "new_password")
				jsonBytes, _ := json.Marshal(&m)
				logger.Infof(c, "request url: %s request body: %s", c.Request.RequestURI, string(jsonBytes))
			}
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		//处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.Request.Host

		logger.Infof(c, "response-info: %3d | %13v | %15s | %s | %s | %s|",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			blw.body.String(),
		)
	}
}

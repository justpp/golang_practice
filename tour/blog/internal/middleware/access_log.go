package middleware

import (
	"bytes"
	"giao/tour/blog/global"
	"giao/tour/blog/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	write, err := w.body.Write(p)
	if err != nil {
		return write, err
	}
	return w.ResponseWriter.Write(p)
}

func (w AccessLogWriter) WriteString(p string) (int, error) {
	w.body.WriteString(p)
	return w.ResponseWriter.WriteString(p)
}

func AccessLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		teeReader := io.TeeReader(c.Request.Body, &buf)
		request, _ := io.ReadAll(teeReader)
		c.Request.Body = io.NopCloser(&buf)

		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now()
		c.Next()
		endTime := time.Now()
		fields := logger.Fields{
			"request":  string(request),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %s, end_time: %s",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}

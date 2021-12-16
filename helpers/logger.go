package helpers

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	ioWriter, _ := InitialLogFile()
	gin.DefaultWriter = io.MultiWriter(ioWriter, os.Stdout)
}

func LoggerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		f, err := InitialLogFile()

		if err != nil {
			ctx.Next()
			return
		}

		handler := gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf(
					"%s | %d | %s | %s %s | %s | %s | %s\n",
					params.TimeStamp.Format(time.RFC1123),
					params.StatusCode,
					params.ClientIP,
					params.Method,
					params.Path,
					params.Latency,
					params.Request.UserAgent(),
					params.ErrorMessage,
				)
			},
			Output: io.MultiWriter(f, os.Stdout),
		})

		handler(ctx)
	}
}

func InitialLogFile() (io.Writer, error) {
	timeNow := time.Now().Format("2006-01-02")

	filename := "LOG-" + timeNow + ".log"

	ioWriter, err := os.OpenFile(path.Join("./logs", filename), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		return nil, err
	}

	return ioWriter, nil
}

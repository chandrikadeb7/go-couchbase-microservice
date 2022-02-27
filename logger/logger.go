package logger

import (
	"github.com/labstack/echo/v4/middleware"
)

func EchoLoggerConfig() middleware.LoggerConfig {
	logMessageTemplate := `{"@timestamp":"${time_rfc3339}", "method":"${method}", "uri":"${uri}", "status":"${status}", "latency":"${latency}"," error":"${error}", "userAgent":"${user_agent}"}` + "\n"
	return middleware.LoggerConfig{
		Format: logMessageTemplate,
	}
}

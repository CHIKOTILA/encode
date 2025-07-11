package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var RequestLoggerMiddleware = middleware.LoggerWithConfig(middleware.LoggerConfig{
	Format: "method=${method}, uri=${uri}, status=${status}\n",
})

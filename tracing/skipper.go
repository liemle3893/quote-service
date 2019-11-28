package tracing

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewPathSkipper create new skipper for <p>path</p>
func NewPathSkipper(path string) middleware.Skipper {
	return func(c echo.Context) bool {
		skip := "/metrics" == c.Path()
		return skip
	}
}

package routing

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler interface {
}

type Middleware interface {
}

type Router struct {
	handler    Handler
	middleware Middleware
}

func New(options ...func(*Router)) *echo.Router {
	router := &Router{}
	for _, option := range options {
		option(router)
	}

	return nil
}

func WithHandler(handler Handler) func(*Router) {
	return func(router *Router) {
		router.handler = handler
	}
}

func WithMiddleware(middleware Middleware) func(*Router) {
	return func(router *Router) {
		router.middleware = middleware
	}
}

func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Set CORS headers
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Response().Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

			// Handle OPTIONS method
			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	}
}
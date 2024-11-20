package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

const (
	readTimeout  = 3 * time.Second
	writeTimeout = 10 * time.Second
)

func New(options ...func(*http.Server)) *http.Server {
	server := new(http.Server)

	for _, option := range options {
		option(server)
	}

	return server
}

func WithHost(host string) func(*http.Server) {
	return func(server *http.Server) {
		server.Addr = host
	}
}

func WithHandler(e *echo.Echo) func(*http.Server) {
	return func(server *http.Server) {
		server.Handler = e
	}
}

func WithDefaultTimeouts() func(*http.Server) {
	return func(server *http.Server) {
		server.ReadTimeout = readTimeout
		server.WriteTimeout = writeTimeout
	}
}

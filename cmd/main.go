package main

import (
	configs "auth/config"
	"auth/handler"
	"auth/logging"
	"auth/middleware"
	"auth/repository/postgres"
	"auth/repository/redis"
	routing "auth/router"
	"auth/server"
	"auth/service"
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const (
	configPath = "./config"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	err := run(ctx)
	if err != nil {
		log.Printf("failed to run app: %s", err)

		return
	}
}

func run(ctx context.Context) error {
	// TODO: add swagger
	// TODO: introduce CI/CD
	// TODO: https and certificates

	// TODO: error handling refactoring
	// TODO: validation errors better output

	// TODO: deploy in cloud (AWS)
	// TODO: hook up cloud logging
	// TODO: move to k8s
	// TODO: introduce terraform

	var logger logging.Logger

	logger, err := logging.NewDefaultLogger()
	if err != nil {
		log.Fatalf("Failed to init logger: %s", err)
	}

	config, err := configs.LoadConfig(configPath)
	if err != nil {
		return err
	}

	db, err := postgres.OpenDB(config.PostgresDb)
	if err != nil {
		logger.Error("failed to connect to Postgres: " + err.Error())

		return err
	}

	repo := postgres.NewRepository(
		postgres.WithDB(db),
	)

	redisDb, err := redis.OpenDB(config.RedisDb)
	if err != nil {
		logger.Error("failed to connect to Redis: " + err.Error())
		return err
	}

	redisRepo := redis.NewRepository(
		redis.WithDB(redisDb),
	)

	userService := service.NewUserService(repo)
	oAuthService := service.NewOAuthService(repo)
	twoFAService := service.NewTwoFAService(repo)
	tokenService := service.NewTokenService(repo, redisRepo)

	middleware := middleware.NewMiddleware()

	handler := handler.NewHandler(
		logger,
		handler.WithUserService(userService),
		handler.WithOAuthService(oAuthService),
		handler.WithTwoFAService(twoFAService),
		handler.WithTokenService(tokenService),
	)

	router := routing.New(
		routing.WithHandler(handler),
		routing.WithMiddleware(middleware),
	)

	srv := server.New(
		server.WithHandler(router),
		server.WithHost(":8080"),
		server.WithDefaultTimeouts(),
	)

	go func() {
		logger.Info("starting http server", logging.Any("port", ":8080"))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to server HTTP server: " + err.Error())

			return
		}
	}()

	<-ctx.Done()

	shutdownTimeout := 5
	shutdownCtx, cancel := context.WithTimeout(ctx, time.Duration(shutdownTimeout)*time.Second)
	defer cancel()

	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error("failed to gracefully shutdown: " + err.Error())

		return err
	}

	logger.Info("Gracefully stopped the service")

	return nil
}

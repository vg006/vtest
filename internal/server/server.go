package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/vg006/vtest/internal/logger"
	shut "github.com/vg006/vtest/internal/shutdown"
)

type Server struct {
	srv   *echo.Echo
	wg    sync.WaitGroup
	logr  *logger.Logger
	Port  string
	srvCh chan bool
	errCh chan error
}

func New(port string) *Server {
	return &Server{
		srv:   echo.New(),
		Port:  port,
		logr:  logger.New(100),
		srvCh: make(chan bool, 1),
		errCh: make(chan error, 1),
	}
}

func (s *Server) RegisterRoutes() {
	s.srv.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api := s.srv.Group("/api")

	api.POST("/details", DetailsHandler)
	api.POST("/routes", RoutesHandler)
}

func (s *Server) Exec() {
	s.logr.Exec()
	s.RegisterRoutes()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.logr.LogInfo("Starting the server on port " + s.Port)
		err := s.srv.Start(s.Port)
		if err != nil {
			s.errCh <- err
		}
		close(s.errCh)
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		select {
		case <-s.srvCh:
			s.logr.LogInfo("Shutdown signal received, shutting down the server...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := s.srv.Shutdown(ctx); err != nil {
				s.logr.LogError("Error during server shutdown: " + err.Error())
			} else {
				s.logr.LogDone("Server shut down gracefully.")
			}
		case err := <-s.errCh:
			if err != nil && err != http.ErrServerClosed {
				s.logr.LogError("Server error: " + err.Error())
				shut.GracefulExit()
			}
		}
	}()
}

func (s *Server) Exit() {
	s.logr.LogInfo("Shutting down server...")
	close(s.srvCh)
	s.wg.Wait()
	s.logr.Exit("Gracefully shutdown server...")
}

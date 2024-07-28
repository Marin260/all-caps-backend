package server

import (
	"net/http"

	"github.com/Marin260/all-caps-backend/internal/api/authhandler"
	"github.com/Marin260/all-caps-backend/internal/api/healthhandlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	authhandler.MountAuthRoutes(r)
	healthhandlers.MountHealthRoutes(r)

	return r
}

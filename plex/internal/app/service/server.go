package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/dabiggm0e/plextrakt/common/monitoring"
	"github.com/dabiggm0e/plextrakt/plex/cmd"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	cfg *cmd.Config
	r   *chi.Mux
	h   *Handler
}

func NewServer(cfg *cmd.Config, h *Handler) *Server {
	return &Server{cfg: cfg, h: h}
}

func (s *Server) Close() {

}

func (s *Server) Start() {
	logrus.Infof("Starting HTTP server on %v", ":"+s.cfg.Port)
	err := http.ListenAndServe(":"+s.cfg.Port, s.r)
	if err != nil {
		logrus.WithError(err).Fatal("error starting HTTP server")
	}
}

func (s *Server) SetupRoutes() {

	s.r = chi.NewRouter()
	s.r.Use(middleware.RequestID)
	s.r.Use(middleware.RealIP)
	s.r.Use(middleware.Logger)
	s.r.Use(middleware.Recoverer)
	s.r.Use(middleware.Timeout(time.Minute))

	for _, route := range routes {
		signature := fmt.Sprintf("%v %v", route.Method, route.Pattern)

		switch route.Method {
		case "POST":
			s.r.Route(route.Pattern, func(r chi.Router) {
				//	r.With(monitoring.Trace(route.Name)).
				r.With(monitoring.Monitor(s.cfg.Name, route.Name, signature)).
					Post("/", s.h.ParsePlexEvent)
			})
		}

	}

}

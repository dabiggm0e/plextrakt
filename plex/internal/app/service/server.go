package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/dabiggm0e/plextrakt/plex/cmd"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/goblog/common/monitoring"
	"github.com/goblog/common/tracing"
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
	/*
	   	for _, route := range routes {
	   		s.r.Route("/plex/events", func(r chi.Router) {
	   			r.With(Trace("Parse_Plex_Event")).
	   				With(Monitor(s.cfg.Name, "ParsePlexEvent", "POST /plex/events")).
	           Post.
	   		})
	   	}
	*/
}

func Monitor(serviceName, routeName, signature string) func(http.Handler) http.Handler {
	summaryVec := monitoring.BuildSummaryVec(serviceName, routeName, signature)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			start := time.Now()
			next.ServeHTTP(rw, req)
			duration := time.Since(start)

			// Store duration of request
			summaryVec.WithLabelValues("duration").Observe(duration.Seconds())

			// Store size of response, if possible.
			size, err := strconv.Atoi(rw.Header().Get("Content-Length"))
			if err == nil {
				summaryVec.WithLabelValues("size").Observe(float64(size))
			}
		})
	}
}

func Trace(opName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			logrus.Infof("starting span for %v", opName)
			span := tracing.StartHTTPTrace(req, opName)
			ctx := tracing.UpdateContext(req.Context(), span)
			next.ServeHTTP(rw, req.WithContext(ctx))

			span.Finish()
			logrus.Infof("finished span for %v", opName)
		})
	}
}

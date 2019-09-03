package monitoring

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/callistaenterprise/goblog/common/router"
	"github.com/callistaenterprise/goblog/common/tracing"
	"github.com/prometheus/client_golang/prometheus"
)

func BuildSummaryVec(serviceName, metricName, metricHelp string) *prometheus.SummaryVec {
	summaryVec := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: serviceName,
			Name:      metricName,
			Help:      metricHelp,
		},
		[]string{"service"},
	)
	prometheus.Register(summaryVec)
	return summaryVec
}

// WithMonitoring optionally adds a middleware that stores request duration and response size into the supplied
// summaryVec
func WithMonitoring(next http.Handler, route router.Route, summary *prometheus.SummaryVec) http.Handler {

	// Just return the next handler if route shouldn't be monitored
	if !route.Monitor {
		return next
	}

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(rw, req)
		duration := time.Since(start)

		// Store duration of request
		summary.WithLabelValues("duration").Observe(duration.Seconds())

		// Store size of response, if possible.
		size, err := strconv.Atoi(rw.Header().Get("Content-Length"))
		if err == nil {
			summary.WithLabelValues("size").Observe(float64(size))
		}
	})
}

func Monitor(serviceName, routeName, signature string) func(http.Handler) http.Handler {
	summaryVec := BuildSummaryVec(serviceName, routeName, signature)

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
			logrus.Infof("req=%v opName=%v\n", req, opName)
			span := tracing.StartHTTPTrace(req, opName) //FIXME TODO
			ctx := tracing.UpdateContext(req.Context(), span)
			next.ServeHTTP(rw, req.WithContext(ctx))

			span.Finish()
			logrus.Infof("finished span for %v", opName)
		})
	}
}

package monitor

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsServer creates a router with an endpoint for prometheus on api/v1/metrics
func MetricsServer() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/api/v1/metrics").Handler(promhttp.Handler())

	return router
}

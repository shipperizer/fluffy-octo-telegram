package monitor

import (
	"github.com/gorilla/mux"
)

func MetricsServer() *mux.Router {
	router := mux.NewRouter()
	NewMonitorAPI().Routes(router)

	return router
}

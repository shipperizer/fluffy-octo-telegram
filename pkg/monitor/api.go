package monitor

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/shipperizer/fluffy-octo-telegram/pkg/api"
)

type monitorAPI struct{}

func (a *monitorAPI) Routes(router *mux.Router) {
	router.PathPrefix("/api/v0/metrics").Handler(promhttp.Handler())
}

func NewMonitorAPI() api.API {
	return &monitorAPI{}
}

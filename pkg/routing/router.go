package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shipperizer/fluffy-octo-telegram/pkg/monitor"
	log "github.com/sirupsen/logrus"
)

func NewRouter(jwkPubPath, jwkPrivPath string) (http.Handler, error) {
	jwks, key, err := LoadJWKs(jwkPubPath, jwkPrivPath)

	if err != nil {
		return nil, err
	}

	log.Debugf("jwks: %v", jwks)
	log.Debugf("key: %v", key)

	router := mux.NewRouter()

	monitor.NewMetricsAPI().(router)
	jwks.NewJWKAPI(jwks).Routes(router)

	return router, nil
}

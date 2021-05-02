package routing

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/shipperizer/fluffy-octo-telegram/pkg/jwks"
	"github.com/shipperizer/fluffy-octo-telegram/pkg/monitor"
)

func NewRouter(jwkPubPath, jwkPrivPath string) (http.Handler, error) {
	jwk, key, err := jwks.LoadJWKs(jwkPubPath, jwkPrivPath)

	if err != nil {
		return nil, err
	}

	log.Debugf("jwk: %v", jwk)
	log.Debugf("key: %v", key)

	router := mux.NewRouter()

	monitor.NewMonitorAPI().Routes(router)
	jwks.NewJWKAPI(jwk).Routes(router)

	return router, nil
}

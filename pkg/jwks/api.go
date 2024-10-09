package jwks

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwk"

	"github.com/shipperizer/fluffy-octo-telegram/pkg/api"
)

type jwkAPI struct {
	jwkSet jwk.Set
}

func (a *jwkAPI) Routes(router *mux.Router) {
	router.HandleFunc("/.well-known/jwks.json", a.ListJWKs).Methods(http.MethodGet, http.MethodOptions)
}

func (a *jwkAPI) ListJWKs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a.jwkSet)
}

func NewJWKAPI(jwkPub jwk.Set) api.API {
	return &jwkAPI{
		jwkSet: jwkPub,
	}
}

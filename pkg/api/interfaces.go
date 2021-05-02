package api

import (
	"github.com/gorilla/mux"
)

type API interface {
	Routes(*mux.Router)
}

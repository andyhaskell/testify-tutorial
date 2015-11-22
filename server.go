package testifytutorial

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	rt := mux.NewRouter()
	rt.Path("/").Methods("GET").HandlerFunc(getLocations)
	rt.Path("/add-location").Methods("POST").HandlerFunc(addLocation)
	return rt
}

func init() {
	http.Handle("/", initRouter())
}

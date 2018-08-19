package server

import (
	"net/http"

	h "recipes/handler"
	m "recipes/middleware"

	"github.com/gorilla/mux"
)

func Router(dep *Dependencies) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", h.PingHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(h.NotFoundHandler)

	router.HandleFunc("/v1/recipe", m.Authenticate(h.CreateRecipeHandler(dep))).Methods("POST").HeadersRegexp("Content-Type", "application/json")
	router.HandleFunc("/v1/recipes/{id}", h.GetRecipeHandler(dep)).Methods("GET").HeadersRegexp("Content-Type", "application/json")
	router.HandleFunc("/v1/recipes", h.ListRecipesHandler(dep)).Methods("GET").HeadersRegexp("Content-Type", "application/json")
	router.HandleFunc("/v1/recipes/{id}", m.Authenticate(h.UpdateRecipeHandler(dep))).Methods("PUT").HeadersRegexp("Content-Type", "application/json")
	router.HandleFunc("/v1/recipes/{id}", m.Authenticate(h.DeleteRecipeHandler(dep))).Methods("DELETE").HeadersRegexp("Content-Type", "application/json")
	router.HandleFunc("/v1/recipes/{id}/rating", h.UpdateRecipeRatingHandler(dep)).Methods("POST").HeadersRegexp("Content-Type", "application/json")
	router.HandleFunc("/v1/recipe", h.SearchRecipeHandler(dep)).Methods("GET").HeadersRegexp("Content-Type", "application/json")
	return router
}

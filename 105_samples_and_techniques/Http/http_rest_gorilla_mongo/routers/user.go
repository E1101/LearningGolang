package routers

import (
	"github.com/gorilla/mux"

	"GolangTraining/105_samples_and_techniques/http_rest_gorilla_mongo/controllers"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/users", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")

	return router
}

package routes

import (
	"000-crud-web-applications/backend/controllers"
	"000-crud-web-applications/backend/utils"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	return router
}

func SetRoutes(router *mux.Router) *mux.Router {
	router = SetUserRoutes(router)

	// Add authentication middleware here
	router.Use(utils.AuthenticationMiddleware)

	return router
}

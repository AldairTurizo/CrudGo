package routes

import (
	"github.com/AldairTurizo/CrudGo/controllers"
	"github.com/gorilla/mux"
)

// SetContactsRoutes agrega las rutas de clientos
func SetContactsRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/clients/{id}", controllers.GetContact).Methods("GET")
	subRouter.HandleFunc("/clients", controllers.GetContacts).Methods("GET")
	subRouter.HandleFunc("/clients", controllers.StoreContact).Methods("POST")
	subRouter.HandleFunc("/clients/{id}", controllers.UpdateContact).Methods("PUT")
	subRouter.HandleFunc("/clients/{id}", controllers.DeleteContact).Methods("DELETE")
}

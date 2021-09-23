package routes
import (
"github.com/AldairTurizo/GoCrud/controllers"
"github.com/gorilla/mux"
)
// SetContactsRoutes agrega las rutas de contactos
func SetContactsRoutes(r *mux.Router) {
subRouter := r.PathPrefix("/api").Subrouter()
subRouter.HandleFunc("/contacts/{id}", controllers.GetContact).Methods("GET")
subRouter.HandleFunc("/contacts", controllers.GetContacts).Methods("GET")
subRouter.HandleFunc("/contacts", controllers.StoreContact).Methods("POST")
subRouter.HandleFunc("/contacts/{id}", controllers.UpdateContact).Methods("PUT")
subRouter.HandleFunc("/contacts/{id}", controllers.DeleteContact).Methods("DELETE")
}
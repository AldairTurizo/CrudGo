package main

import (
	"log"
	"net/http"

	"github.com/AldairTurizo/CrudGo/routes"
	"github.com/AldairTurizo/CrudGo/utils"
	"github.com/gorilla/mux"
)

func main() {
	// Migracion de la base de datos
	utils.MigrateDB()
	// router para el manejo de las rutas de la aplicacion
	r := mux.NewRouter()
	// se agregan las rutas de contactos
	routes.SetContactsRoutes(r)
	// generacion de un nuevo servidor, especificamos el puerto y las rutas
	srv := http.Server{
		Addr:    ":4000",
		Handler: r,
	}
	log.Println("Running on port 4000")
	// se ejecuta el servidor
	log.Println(srv.ListenAndServe())
}

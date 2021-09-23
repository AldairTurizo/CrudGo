package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AldairTurizo/CrudGo/models"
	"github.com/AldairTurizo/CrudGo/utils"
	"github.com/gorilla/mux"
)

// GetContact obtiene un cliento por su ID
func GetContact(w http.ResponseWriter, r *http.Request) {
	// Estructura vacia donde se gurdarán los datos
	client := models.Client{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM clients WHERE ID = ?
	db.Find(&client, id)
	// Se comprueba que exista el registro
	if client.ID > 0 {
		// Se codifican los datos a formato JSON
		j, _ := json.Marshal(client)
		// Se envian los datos
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		// Si no existe se envia un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

// GetContacts obtiene todos los clientos
func GetContacts(w http.ResponseWriter, r *http.Request) {
	// Slice (array) donde se guardaran los datos
	clients := []models.Client{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM clients
	db.Find(&clients)
	// Se codifican los datos a formato JSON
	j, _ := json.Marshal(clients)
	// Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)
}

// StoreContact guarda un nuevo cliento
func StoreContact(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se gurdaran los datos del body
	client := models.Client{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se decodifican los datos del body a la estructura client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	// Se guardan los datos en la DB
	err = db.Create(&client).Error
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}
	// Se codifica el nuevo registro y se devuelve
	j, _ := json.Marshal(client)
	utils.SendResponse(w, http.StatusCreated, j)
}

// UpdateContact modifica los datos de un cliento por su ID
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	// Estructuras donde se almacenaran los datos
	clientFind := models.Client{}
	clientData := models.Client{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se buscan los datos
	db.Find(&clientFind, id)
	if clientFind.ID > 0 {
		// Si existe el registro se decodifican los datos del body
		err := json.NewDecoder(r.Body).Decode(&clientData)
		if err != nil {
			// Sí hay algun error en los datos se devolvera un error 400
			utils.SendErr(w, http.StatusBadRequest)
			return
		}
		// Se modifican los datos
		db.Model(&clientFind).Updates(clientData)
		// Se codifica el registro modificado y se devuelve
		j, _ := json.Marshal(clientFind)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

// DeleteContact elimina un cliento por ID
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se guardara el registo buscado
	client := models.Client{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se busca el cliento
	db.Find(&client, id)
	if client.ID > 0 {
		// Sí existe, se borra y se envia contenido vacio
		db.Delete(client)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

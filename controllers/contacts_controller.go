package controllers

import (
"encoding/json"
"fmt"
"net/http"
"github.com/AldairTurizo/GoCrud/models"
"github.com/AldairTurizo/GoCrud/utils"
"github.com/gorilla/mux"
)

// GetContact obtiene un contacto por su ID
func GetContact(w http.ResponseWriter, r *http.Request) {
	// Estructura vacia donde se gurdarán los datos
	contact := models.Contact{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM contacts WHERE ID = ?
	db.Find(&contact, id)
	// Se comprueba que exista el registro
	if contact.ID > 0 {
	// Se codifican los datos a formato JSON
	j, _ := json.Marshal(contact)
	// Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)
	} else {
	// Si no existe se envia un error 404
	utils.SendErr(w, http.StatusNotFound)
	}
	}

	// GetContacts obtiene todos los contactos
func GetContacts(w http.ResponseWriter, r *http.Request) {
	// Slice (array) donde se guardaran los datos
	contacts := []models.Contact{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM contacts
	db.Find(&contacts)
	// Se codifican los datos a formato JSON
	j, _ := json.Marshal(contacts)
	// Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)
	}

	// StoreContact guarda un nuevo contacto
func StoreContact(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se gurdaran los datos del body
	contact := models.Contact{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se decodifican los datos del body a la estructura contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
	// Sí hay algun error en los datos se devolvera un error 400
	fmt.Println(err)
	utils.SendErr(w, http.StatusBadRequest)
	return
	}
	// Se guardan los datos en la DB
	err = db.Create(&contact).Error
	if err != nil {
	// Sí hay algun error al guardar los datos se devolvera un error 500
	fmt.Println(err)
	utils.SendErr(w, http.StatusInternalServerError)
	return
	}
	// Se codifica el nuevo registro y se devuelve
	j, _ := json.Marshal(contact)
	utils.SendResponse(w, http.StatusCreated, j)
	}

	// UpdateContact modifica los datos de un contacto por su ID
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	// Estructuras donde se almacenaran los datos
	contactFind := models.Contact{}
	contactData := models.Contact{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se buscan los datos
	db.Find(&contactFind, id)
	if contactFind.ID > 0 {
	// Si existe el registro se decodifican los datos del body
	err := json.NewDecoder(r.Body).Decode(&contactData)
	if err != nil {
	// Sí hay algun error en los datos se devolvera un error 400
	utils.SendErr(w, http.StatusBadRequest)
	return
	}
	// Se modifican los datos
	db.Model(&contactFind).Updates(contactData)
	// Se codifica el registro modificado y se devuelve
	j, _ := json.Marshal(contactFind)
	utils.SendResponse(w, http.StatusOK, j)
	} else {
	// Sí no existe el registro especificado se devuelde un error 404
	utils.SendErr(w, http.StatusNotFound)
	}
	}

	// DeleteContact elimina un contacto por ID
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se guardara el registo buscado
	contact := models.Contact{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se busca el contacto
	db.Find(&contact, id)
	if contact.ID > 0 {
	// Sí existe, se borra y se envia contenido vacio
	db.Delete(contact)
	utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
	// Sí no existe el registro especificado se devuelde un error 404
	utils.SendErr(w, http.StatusNotFound)
	}
	}
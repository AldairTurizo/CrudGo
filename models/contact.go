package models

import "github.com/jinzhu/gorm"

// Contact modelo para clientos
type Client struct {
	gorm.Model
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Telefono  string `json:"telefono"`
	Barrio    string `json:"barrio"`
	Direccion string `json:"direccion"`
	Email     string `json:"email"`
	UrlImages  string `json:"url_images"`
}

package controller

import (
	"golantha/templates"
	"net/http"
)

func ConnexionHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "create", nil)
}

func AventurierHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "aventurier", nil)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "404", nil)
}

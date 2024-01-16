package main

import (
	"golantha/root"
	"golantha/sauvegarde"
	"golantha/templates"
	"html/template"
	"net/http"
)

var Temp *template.Template

func main() {
	sauvegarde.PrintColorResult("blue", "http://localhost:8080/create \n")
	templates.Inittemplate()
	root.InitServe()
	http.ListenAndServe("localhost:8080", nil)
}

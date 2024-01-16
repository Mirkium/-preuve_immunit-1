package root

import (
	"golantha/controller"
	"golantha/templates"
	"net/http"
)

func InitServe() {
	FileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", FileServer))
	http.HandleFunc("/create", controller.ConnexionHandler)
	http.HandleFunc("/aventurier", controller.AventurierHandler)
	http.HandleFunc("/404", controller.ErrorHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		templates.Temp.ExecuteTemplate(w, "404", nil)
	})

}

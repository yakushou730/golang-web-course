package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yakushou730/golang-web-course/views"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
	homeView        *views.View
	contactView     *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is FAQ page.</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>This is 404 page.</h1>")
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}

	r := httprouter.New()
	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)

	var h http.Handler = http.HandlerFunc(notFound)
	r.NotFound = h
	http.ListenAndServe(":3000", r)
}

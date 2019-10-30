package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:yakushou730@gmail.com\">"+
		"yakushou730@gmail.com</a>.")
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
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)

	var h http.Handler = http.HandlerFunc(notFound)
	r.NotFound = h
	http.ListenAndServe(":3000", r)
}

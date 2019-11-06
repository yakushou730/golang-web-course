package main

import (
	"net/http"

	"github.com/yakushou730/golang-web-course/controllers"

	"github.com/gorilla/mux"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()
	galleriesC := controllers.NewGalleries()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/galleries/new", galleriesC.New).Methods("GET")

	http.ListenAndServe(":3000", r)
}

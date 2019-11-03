package controllers

import "github.com/yakushou730/golang-web-course/views"

type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
		Faq:     views.NewView("bootstrap_v2", "views/static/faq.gohtml"),
	}
}

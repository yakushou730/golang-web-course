package controllers

import (
	"net/http"

	"github.com/yakushou730/golang-web-course/views"
)

type Galleries struct {
	NewView *views.View
}

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
	}
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	g.NewView.Render(w, nil)
}

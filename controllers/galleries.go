package controllers

import (
	"github.com/yakushou730/golang-web-course/models"

	"github.com/yakushou730/golang-web-course/views"
)

type Galleries struct {
	New *views.View
	gs  models.GalleryService
}

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}

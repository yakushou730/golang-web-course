package views

import (
	"html/template"
	"path/filepath"
)

type View struct {
	Template *template.Template
	Layout   string
}

var (
	LayoutDir  string = "views/layouts/"
	TempateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TempateExt)
	if err != nil {
		panic(err)
	}
	return files
}

package views

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/gorilla/csrf"

	"github.com/yakushou730/golang-web-course/context"
)

type View struct {
	Template *template.Template
	Layout   string
}

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	// We are changing how we create our templates, calling
	// New("") to give us a template that we can add a function to
	// before finally passing in files to parse as part of the template.
	t, err := template.New("").Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			// If this is called without being replace with a proper implementation
			// returning an error as the second argument will cause our template
			// package to return an error when executed.
			return "", errors.New("csrfField is not implemented")
		},
		"pathEscape": func(s string) string {
			return url.PathEscape(s)
		},
	}).ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		// We need to do this so we can access the data in a var
		// with the type Data.
		vd = d
	default:
		// If the data IS NOT of the type Data, we create one
		// and set the data to the Yield field like before.
		vd = Data{
			Yield: data,
		}
	}
	// Lookup and set the user to the User field
	vd.User = context.User(r.Context())
	var buf bytes.Buffer
	// We need to create the csrfField using the current http request.
	csrfField := csrf.TemplateField(r)
	tpl := v.Template.Funcs(template.FuncMap{
		// We can also change the return type of our function, since we no longer
		// need to worry about errors.
		"csrfField": func() template.HTML {
			// We can then create this closure that returns the csrfField for
			// any templates that need access to it.
			return csrfField
		},
	})
	// Then we continue to execute the template just like before.
	err := tpl.ExecuteTemplate(&buf, v.Layout, vd)
	if err != nil {
		http.Error(w, "Something went wrong. If the problem "+
			"persists, please email yakushou730@gmail.com",
			http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}

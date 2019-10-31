package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name  string
		Email string
	}{
		Name:  "<script>alert('Howdy!');</script>",
		Email: "yakushou730@gmail.com",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

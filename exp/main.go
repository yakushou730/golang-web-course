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
		Name      string
		Email     string
		Number    int32
		Decimal   float32
		TestMap   map[string]int32
		Condition bool
	}{
		Name:    "<script>alert('Howdy!');</script>",
		Email:   "yakushou730@gmail.com",
		Number:  10,
		Decimal: 123.456,
		TestMap: map[string]int32{
			"Test1": 123,
			"Test2": 456,
		},
		Condition: false,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

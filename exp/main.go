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

	type Data struct {
		Name           string
		emailExtension string
	}

	data := Data{"Joshua Hauhio", "@gmail.com"}
	data.Name = "<script>alert('Howdy!');</script>"

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

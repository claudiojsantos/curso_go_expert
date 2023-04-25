package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", TemplateHandler)
	http.ListenAndServe(":8080", nil)
}

func TemplateHandler(res http.ResponseWriter, req *http.Request) {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(res, Cursos{
		{"Go", 40},
		{"Java", 40},
		{"Python", 40},
	})

	if err != nil {
		panic(err)
	}
}

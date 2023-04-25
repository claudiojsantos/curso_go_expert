package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	template := template.New("CursoTemplate")
	template, _ = template.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := template.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}

}

package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// Take a filePath as argument, return 2 arguments (types Template and error)
func Parse(filepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil { // if there is an error return an empty Template and an error
		return Template{}, fmt.Errorf("Parsing template: %w", err) // Wrap error with extra info
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

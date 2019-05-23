package main

import (
	"html/template"
	"path/filepath"
	"time"
)

const (
	templatesDir = "templates"
)

// parseTemplate applies a given file to the base template.
func parseTemplate(filename string) *template.Template {
	funcMap := template.FuncMap{
		"datef": func(date time.Time) string {
			return date.Format("Monday, January 2, 2006")
		},
	}

	return template.Must(
		template.New("base.html").Funcs(funcMap).ParseFiles(
			filepath.Join(templatesDir, "base.html"),
			filepath.Join(templatesDir, filename),
		),
	)
}

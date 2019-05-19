package main

import (
	"html/template"
	"path/filepath"
)

const (
	templatesDir = "templates"
)

// parseTemplate applies a given file to the base template.
func parseTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(
		filepath.Join(templatesDir, "base.html"),
		filepath.Join(templatesDir, filename),
	))
}

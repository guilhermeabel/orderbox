package main

import (
	"io/fs"
	"path/filepath"
	"text/template"
	"time"

	"github.com/guilhermeabel/orderbox/cmd/ui"
	"github.com/guilhermeabel/orderbox/internal/models"
)

type templateData struct {
	Order           *models.Order
	Orders          []*models.Order
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

var functions = template.FuncMap{
	"parseDate": parseDate,
}

func parseDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.html",
			"html/components/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

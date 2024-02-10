package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/guilhermeabel/orderbox/internal/models"
)

type templateData struct {
	Order       *models.Order
	Orders      []*models.Order
	CurrentYear int
	Form        any
}

var functions = template.FuncMap{
	"parseDate": parseDate,
}

func parseDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("../ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("../ui/html/components/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

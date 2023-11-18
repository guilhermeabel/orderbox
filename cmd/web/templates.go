package main

import (
	"path/filepath"
	"text/template"

	"github.com/guilhermeabel/orderbox/internal/models"
)

type templateData struct {
	Order       *models.Order
	Orders      []*models.Order
	CurrentYear int
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles("../ui/html/base.html")
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

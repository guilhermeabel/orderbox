package main

import (
	"embed"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

//go:embed public/*
var assetsFS embed.FS

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.FS(assetsFS))))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/order/view/:id", dynamic.ThenFunc(app.viewOrder))
	router.Handler(http.MethodGet, "/order/create", dynamic.ThenFunc(app.createOrder))
	router.Handler(http.MethodPost, "/order/create", dynamic.ThenFunc(app.createOrderPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}

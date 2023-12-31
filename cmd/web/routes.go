package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("../ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/order/view/:id", app.viewOrder)
	router.HandlerFunc(http.MethodPost, "/order/create", app.createOrderPost)

	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/guilhermeabel/orderbox/internal/models"
	"github.com/guilhermeabel/orderbox/internal/validator"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	orders, err := app.orders.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Orders = orders

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) viewOrder(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	order, err := app.orders.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	flash := app.sessionManager.PopString(r.Context(), "flash")

	data := app.newTemplateData(r)
	data.Order = order

	data.Flash = flash

	app.render(w, http.StatusOK, "order.html", data)
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = createOrderForm{}

	app.render(w, http.StatusOK, "create.html", data)
}

func (app *application) createOrderPost(w http.ResponseWriter, r *http.Request) {
	var form createOrderForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	expires := time.Now().AddDate(0, 0, 1)

	id, err := app.orders.Insert(form.Title, form.Content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Order successfully created")

	http.Redirect(w, r, fmt.Sprintf("/order/view/%d", id), http.StatusSeeOther)
}

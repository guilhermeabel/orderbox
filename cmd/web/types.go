package main

import "github.com/guilhermeabel/orderbox/internal/validator"

type createOrderForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	validator.Validator `form:"-"`
}

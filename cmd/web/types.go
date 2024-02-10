package main

import "github.com/guilhermeabel/orderbox/internal/validator"

type createOrderForm struct {
	Title   string
	Content string
	validator.Validator
}

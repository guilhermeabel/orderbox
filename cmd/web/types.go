package main

import "github.com/guilhermeabel/orderbox/internal/validator"

type createOrderForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	validator.Validator `form:"-"`
}

type userSignupForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

type userLoginForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

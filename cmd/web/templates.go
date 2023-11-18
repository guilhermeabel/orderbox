package main

import "github.com/guilhermeabel/orderbox/internal/models"

type templateData struct {
	Order  *models.Order
	Orders []*models.Order
}

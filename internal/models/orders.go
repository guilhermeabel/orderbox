package models

import (
	"database/sql"
	"time"
)

type Order struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type OrderModel struct {
	DB *sql.DB
}

func (m *OrderModel) Insert(title, content string, expires time.Time) (int, error) {
	stmt := `INSERT INTO orders (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), ?)`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (m *OrderModel) Get(id int) (*Order, error) {
	return nil, nil
}

func (m *OrderModel) Latest() ([]*Order, error) {
	return nil, nil
}

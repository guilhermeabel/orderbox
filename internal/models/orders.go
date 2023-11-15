package models

import (
	"database/sql"
	"errors"
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
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *OrderModel) Get(id int) (*Order, error) {
	stmt := `SELECT id, title, content, created, expires FROM orders
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	o := &Order{}

	err := row.Scan(&o.ID, &o.Title, &o.Content, &o.Created, &o.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}

		return nil, err
	}

	return o, nil
}

func (m *OrderModel) Latest() ([]*Order, error) {
	return nil, nil
}

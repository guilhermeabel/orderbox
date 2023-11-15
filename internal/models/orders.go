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
	stmt := `SELECT id, title, content, created, expires FROM orders
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []*Order{}

	for rows.Next() {
		o := &Order{}

		err = rows.Scan(&o.ID, &o.Title, &o.Content, &o.Created, &o.Expires)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

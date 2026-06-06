package repository

import (
	"database/sql"
	"fmt"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

func (r *SQLiteRepo) SaveOrder(customer string, products string, total float64, status string) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		customer, products, total, status,
	)
	if err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}
	return nil
}

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"example/solid/internal/notifications"
	"example/solid/internal/repository"
	"example/solid/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			customer TEXT NOT NULL,
			products TEXT NOT NULL,
			total REAL NOT NULL,
			status TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewSQLiteRepo(db)

	// Демонстрация с EmailSender
	emailNotifier := notifications.NewEmailSender()
	orderServiceEmail := service.NewOrderService(repo, emailNotifier)

	err = orderServiceEmail.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}

	// Демонстрация с SMSSender
	smsNotifier := notifications.NewSMSSender()
	orderServiceSMS := service.NewOrderService(repo, smsNotifier)

	err = orderServiceSMS.CreateOrder("Мария", []string{"orange", "grape"}, 15.0)
	if err != nil {
		log.Fatal(err)
	}
}

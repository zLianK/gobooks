package main

import (
	"context"
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		panic(err)
	}

	bookService := service.NewBookService(db)
	bookHandlers := web.NewBookHandlers(bookService)

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookById)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}

func createTable(db *sql.DB) error {
	_, err := db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS books (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			genre TEXT NOT NULL
		);`,
	)
	return err
}

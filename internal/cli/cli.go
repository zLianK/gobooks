package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCli struct {
	service *service.BookService
}

func NewBookCli(service *service.BookService) *BookCli {
	return &BookCli{service: service}
}

func (cli *BookCli) Run() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gobooks <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("usage: gobooks search <book title>")
			return
		}
		cli.searchBooks(os.Args[2])
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("usage: gobooks simulate <book id> <book id> <book id> ...")
			return
		}
		cli.simulateReading(os.Args[2:])

	}
}

func (cli *BookCli) searchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)
	if err != nil {
		fmt.Println("error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("no books found")
		return
	}

	fmt.Printf("%d books found\n", len(books))
	for _, book := range books {
		fmt.Printf(
			"Id: %d, Title: %s, Author: %s, Genre: %s\n",
			book.ID, book.Title, book.Author, book.Genre,
		)
	}
}

func (cli *BookCli) simulateReading(idsStr []string) {
	var ids []int
	for _, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("invalid book id:", idStr)
			continue
		}
		ids = append(ids, id)
	}
	responses := cli.service.SimulateMultipleReadings(ids, 2*time.Second)
	for _, response := range responses {
		fmt.Println(response)
	}
}

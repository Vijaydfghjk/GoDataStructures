package main

import (
	"fmt"
	"time"
)

type Book struct {
	ISBN     string
	Title    string
	Author   string
	NumPages int
	Borrowed bool
	DueDate  time.Time
}

// Member represents a library member.
type Member struct {
	ID            int
	Name          string
	BooksBorrowed []Book
}

// Library represents a library.
type Library struct {
	Books   []Book
	Members []Member
}

func (lib *Library) BorrowBook(memberID int, isbn string) error {

	// Find the member

	var mem *Member

	for i, _ := range lib.Members {

		if lib.Members[i].ID == memberID {

			mem = &lib.Members[i]
		}

	}

	if mem == nil {

		return fmt.Errorf("member with ID %d not found", memberID)
	}

	// Find book

	var book *Book

	for i, _ := range lib.Books {

		if lib.Books[i].ISBN == isbn {

			book = &lib.Books[i]
		}
	}

	if book == nil {

		return fmt.Errorf("book with ISBN %s is already borrowed", isbn)
	}

	book.Borrowed = true
	book.DueDate = time.Now().AddDate(0, 0, 14)

	mem.BooksBorrowed = append(mem.BooksBorrowed, *book)
	return nil
}

func main() {

	lib := Library{

		Books: []Book{

			{ISBN: "4233", Title: "The Go Programming Language", Author: "Alan A. A. Donovan & Brian W. Kernighan", NumPages: 400, Borrowed: false},
			{ISBN: "7245", Title: "Design Patterns", Author: "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides", NumPages: 416, Borrowed: false},
		},

		Members: []Member{

			{ID: 1, Name: "Ram", BooksBorrowed: nil},
			{ID: 2, Name: "Manoji", BooksBorrowed: nil},
			{ID: 3, Name: "King", BooksBorrowed: nil},
		},
	}

	err := lib.BorrowBook(1, "4233")
	if err != nil {

		fmt.Println("Error is :", err)
	} else {

		fmt.Println("Book borrowed successfully!")
	}

	// Print borrowed books of a member

	for _, v := range lib.Members {

		//fmt.Printf("%s's borrowed books:\n", v.Name)
		//fmt.Println("val is", v.BooksBorrowed)

		for _, book := range v.BooksBorrowed {

			fmt.Printf("Title :%s\nISBN :%s\nAuther :%s", book.Title, book.ISBN, book.Author)
		}
	}
}

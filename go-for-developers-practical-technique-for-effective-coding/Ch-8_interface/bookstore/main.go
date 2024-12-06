package main

import (
	bsk "bookstore/src"
	"fmt"
)

func main() {
	fmt.Println("BookStore Interface Simulation ...")

	// crate a new book store
	store := bsk.NewBookStore()
	fmt.Println(store)

	// add books in the book sotre
	book1 := bsk.Book{ID: "101", Title: "Book1", Author: "Ailly"}
	book2 := bsk.Book{ID: "102", Title: "Book2", Author: "Billy"}

	// create a list of books
	bookList := []*bsk.Book{
		&book1,
		&book2,
	}

	//for _, book := range bookList {
	//	fmt.Println(*book)
	//}

	store.AddListOfBooks(bookList)

	// list all books in the store
	store.ListBooks()

	// Delete a book from the list
	err := store.RemoveBook("103")
	if err != nil {
		fmt.Println(err)
	}

	// list all books in the store after the book deletion
	store.ListBooks()

}

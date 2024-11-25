package main

import (
	"fmt"
)

func main() {
	fmt.Println("Rough Playground for the Book store application design")
	// create a book and add this to the bookstore
	book := Book{
		ID:           "B001",
		Title:        "The Catcher in the Rye",
		Author:       "J.D. Salinger",
		IsBorrowed:   false,
		BorrowedBy:   "",
		LifeCycle:    0,
		MaxLifeCycle: 10,
	}

	bookStore := BookStore{}
	bookStore.AddBook(&book)

	// print the book details
	fmt.Printf("Book: %+v\n", book)

	// print all books in the bookstore
	bookStore.PrintBooks()

	// update the book details
	book.UpdateBookDetails("To Kill a Mockingbird", "Harper Lee")
	fmt.Printf("Updated Book: %+v\n", book)

	// print the updated book details
	fmt.Println("Printing all books after updating book details... \n")
	bookStore.PrintBooks()

}

// create a Book struct and BookStore struct havign list of bvooks and a methods to add a book anbd a method to iterate over the books to print the index and values
type Book struct {
	ID           string
	Title        string
	Author       string
	IsBorrowed   bool
	BorrowedBy   string
	LifeCycle    int
	MaxLifeCycle int
}

type BookStore struct {
	Books []*Book // instead of create its own copy of the Book, refere to the original Book struct which created the books
	// so that any changes made to the Book struct will be reflected in the BookStore
}

func (bs *BookStore) AddBook(book *Book) {
	bs.Books = append(bs.Books, book)
}

func (bs *BookStore) PrintBooks() {
	fmt.Println("All Books in the Bookstore: ")
	for i, v := range bs.Books {
		fmt.Printf("Index: %d, Value: %+v\n", i, v)
	}
}

// method to chage the Book credetials with pointer receiver
func (b *Book) UpdateBookDetails(newTitle, newAuthor string) {
	b.Title = newTitle
	b.Author = newAuthor
}

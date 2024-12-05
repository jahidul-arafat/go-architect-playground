package main

import (
	"fmt"
	kbk "k8s-bookstore-operator/pkg/resources"
)

func main() {
	// this is the Entry point of the application
	fmt.Println("k8s BookStore Operator v0.1")

	// Create 2 books
	book1 := &kbk.Book{ID: "101", Title: "Book 1"}
	book2 := &kbk.Book{ID: "102", Title: "Book 2"}

	// create a book list
	bookList := []*kbk.Book{book1, book2}
	// create book3 and add it to the book list
	book3 := &kbk.Book{ID: "103", Title: "Book 3"}
	bookList = append(bookList, book3)

	for _, book := range bookList {
		book.Title = "New Title"        // mutate the book title
		fmt.Printf("Book: %+v\n", book) // dereference the book struct to print its values
	}

	fmt.Println("\n\nBook1: ", book1)

	// Create 2 librarians
	librarian1 := &kbk.Librarian{ID: "201", Name: "Librarian 1"}
	librarian2 := &kbk.Librarian{ID: "202", Name: "Librarian 2"}

	//// Create a bookstore and add the books and librarians
	// reference type
	store := &kbk.BookStore{
		ID:            "BKS_101",
		Name:          "Bookstore 1",
		Location:      "Location 1",
		BookList:      []*kbk.Book{book1, book2},
		LibrarianList: []*kbk.Librarian{librarian1, librarian2},
	}

	fmt.Printf("Bookstore: %+v\n", store)

	// Add a Book via AddBook()
	AddBook(store, "104", "Book4")
	store.ID = "BKS_101_NEW"

	fmt.Printf("\n\nBookstore[updated]: %+v\n", store)

	// create a copy of the bookstore
	store2 := store
	store2.ID = "BKS_101_store2"
	fmt.Printf("\n\nBookstore[updated/store]: %+v\n", store)
	fmt.Printf("\n\nBookstore[updated/store2]: %+v\n", store2)

	// change the librarian1 name in the bookstore
	store.LibrarianList[0].Name = "New Librarian 1"

	fmt.Printf("\n\nBookstore after changing librarian name: %+v\n", store)

	fmt.Println("\nOriginal Librarian Name: ", librarian1.Name)
}

func AddBook(store *kbk.BookStore, bookID, bookTitle string) {
	// create the book
	book := &kbk.Book{ID: bookID, Title: bookTitle}
	store.BookList = append(store.BookList, book)

}

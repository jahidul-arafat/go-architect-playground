package src

import "fmt"

// NewBookStore create a new book store
func NewBookStore() *Store {
	return &Store{
		books: []*Book{},
	}
}

// AddABook Add a Book to the Book Store
// Method for Struct <Store>
// pointer receiver to avoid local copy
// AddBook is an interface
func (store *Store) AddABook(book *Book) {
	store.books = append(store.books, book) // appending the book to the booklist
	fmt.Printf("Book Added: %s by %s\n", book.Title, book.Author)
}

// AddListOfBooks  Add a list of Books to the Book Store
func (store *Store) AddListOfBooks(bookList []*Book) {
	for _, book := range bookList {
		store.books = append(store.books, book)
		fmt.Printf("[into List]: Book Added: %s by %s\n", book.Title, book.Author)

	}
}

// RemoveBook implicit implementation of RemoveBook() interface
// remove a book from the book store
// search the book by bookID
func (store *Store) RemoveBook(bookID string) error {
	for i, book := range store.books {
		if bookID == book.ID {
			store.books = append(store.books[:i], store.books[i+1:]...) // remove the book at i_th position
			fmt.Printf("Book Removed [%s]:%s\n", book.ID, book.Title)
			return nil
		}
	}
	return fmt.Errorf("	--> book with ID %s not found\n", bookID)
}

// ListBooks list all books in the bookstore
func (store *Store) ListBooks() {
	fmt.Println("Listing all books in Store ...")
	for _, book := range store.books {
		fmt.Printf("- [%s] %s by %s\n", book.ID, book.Title, book.Author)
	}
}

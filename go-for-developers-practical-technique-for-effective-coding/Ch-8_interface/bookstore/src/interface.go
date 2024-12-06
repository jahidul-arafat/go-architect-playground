package src

// list all the interfaces which will be later implemented by different object types
// Note: interface only contains the method signatures, not any field

// this interface defines the behavior for managing books
type InventoryManager interface {
	// define all the methods
	AddABook(book *Book)            // returns NULL
	AddListOfBooks(books []*Book)   // returns NULL
	RemoveBook(bookID string) error // returns error if any
	ListBooks()                     // list all books in the BookStore
}

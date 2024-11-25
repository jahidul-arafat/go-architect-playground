package bookstore

// Step-1: Create custom data types for the application

// Book Create type <Book> struct;
// represents - books in the bookstore
type Book struct {
	ID           string // Unique identifier for the books
	Title        string
	Author       string
	IsBorrowed   bool   // default "false"
	IsDonated    bool   // default "false"
	BorrowedBy   string // default ""
	LifeCycle    int    // Lifecycle tracks how many times the book has been used
	MaxLifeCycle int    // Maximum allowed lifecycle
}

// BookStore - Create type <BookStore> struct
// represents - a storage for all the books - []Book
type BookStore struct {
	BookList []Book // slice of books // <index,value> // must use for range to iterage over the item's index and value
}

// Librarian - Create type <Librarian>
// represents a librarian who manage the bookstore
type Librarian struct {
	ID                 string // Unique identified for teh librarian
	Name               string
	ProcessedBookCount int // Track the number of books processed by the librarian
	SearchFrequency    int // Tracks how many times this librarian is being searched
}

// Student Create type <Student>
// represents - a student who borrows books
type Student struct {
	Name              string // Unique identifier for the student
	ID                string
	BorrowedBookCount int // Tracks the number of books current
	SearchFrequency   int // Tracks how many times this student is being searched
}

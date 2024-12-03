package resources

// Book - Book Type
// Custom Defined Struct object (aka Object Type) for the project
type Book struct {
	ID         string
	Title      string
	IsBorrowed bool
	BorrowedBy string
}

// BookStore - BookStore type which will have slice of books
type BookStore struct {
	ID            string      // a bookstore should have an unique identifier
	Name          string      // name of the bookstore
	Location      string      // location of the bookstore
	BookList      []Book      // slice of Books (aka list of all books in the bookstore)
	LibrarianList []Librarian // slice of Librarian (aka list of all librarians currently working in the bookstore)
}

// Librarian - Librarian type
// A BookStore can have a number of Librarians
type Librarian struct {
	ID                 string
	Name               string
	ProcessedBookCount int // number of books processed by the Librarian
}

// Student - Student Type
// Books from the BookStores will be borrowed by the Students
type Student struct {
	ID            string
	Name          string
	BooksBorrowed int // number of books borrowed by the students
}

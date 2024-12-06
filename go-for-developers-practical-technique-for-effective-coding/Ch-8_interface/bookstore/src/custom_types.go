package src

// Book define the book object type
type Book struct {
	ID     string
	Title  string
	Author string
}

type Store struct {
	books []*Book // a slice off all pointers to the Book objects
	// avoid making a local copy of the Book objects
}

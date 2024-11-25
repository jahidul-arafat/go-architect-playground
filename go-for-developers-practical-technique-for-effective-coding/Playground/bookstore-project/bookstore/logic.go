package bookstore

import "fmt" // go fmt implictely derefences pointers when printing structs

// Method-01: AddBook
// For: type <Librarian>
// Librarian to add a new book to the BookStore
// the BookStore must be in the global score; thus must be used as "call by reference (pointer)"
// l *Librarian -> pointer receiver, so if l.Name="Changed Name", then this changes will be globally available

func (l *Librarian) AddBook(store *BookStore, bookID, title, author string, maxLifeCycle int) {
	// create a new book with the book credentials
	newBook := Book{
		ID:           bookID,
		Title:        title,
		Author:       author,
		MaxLifeCycle: maxLifeCycle,
	}

	// add this new book in the BookStore - is a slice of Book
	// type BookStore struct {Books []Book}
	// deferencing of BookStore will be automatically hanled by Go at background
	store.BookList = append(store.BookList, newBook)

	// make a global update of the Books Count processed by the Librarian
	l.ProcessedBookCount++ // Incrementes librarian's processed book count
	fmt.Printf("Librarian %s (ID: %s) added the book 'Title: %s' by Author '%s' (BookID: %s)\n",
		l.Name, l.ID, newBook.Title, newBook.Author, newBook.ID)
}

// BorrowBook Method-02: BorrowBook
// By: Student
// Student will borrow a book by BookID
// pointer receiver
func (s *Student) BorrowBook(store *BookStore, bookID string) {
	for i := range store.BookList {
		// get the book pointer of that specific index // make sure you are not creating any local copy of any Book Struct
		book := &store.BookList[i] // get the pointer to the original Book struct so any changes made here persist even after the function score exists
		// Go's syntactic sugar is applied here
		// Go will automatically derefence the struct pointers when accessing their fields, making the code more cleaner and more concise
		// book.ID       // Automatically dereferences `book` to access the `ID` field
		//(*book).ID    // Explicit dereference to access the `ID` field

		// if (*book).ID == bookID && !(*book).IsBorrowed
		if book.ID == bookID && !book.IsBorrowed {
			book.IsBorrowed = true
			book.BorrowedBy = s.ID
			book.LifeCycle++
			s.BorrowedBookCount++
			fmt.Printf("Student %s borrowed the book %s (BookID: %s)\n",
				s.Name, book.Title, book.ID)
			return // return from the function // exist the function early //
			// without this return the loop continue iterating through all the books in the BookStore even through the requried book is already processed
		}
	}
	fmt.Printf("BookID %s is not available to borrow\n", bookID)
}

// ReturnBook - Methods03
// allows a student o return a borrowed book by BookID
func (s *Student) ReturnBook(store *BookStore, bookID string) {
	fmt.Printf("Returning books for Student ID: %s\n", s.ID)
	for i := range store.BookList {
		book := &store.BookList[i] // get the memory address of the Book struct which has created that book
		// get the pointer to the original Book struct so any changes made here persist even after the function score exists
		// if (*book).ID == bookID && (*book).IsBorrowed && (*book).BorrowedBy
		// book.ID       // Automatically dereferences `book` to access the `ID` field
		//(*book).ID    // Explicit dereference to access the `ID` field
		if book.ID == bookID && book.IsBorrowed && book.BorrowedBy == s.ID { //(*s).ID - explicit deferencing
			book.IsBorrowed = false
			book.BorrowedBy = ""
			s.BorrowedBookCount-- // student returned the book
			fmt.Printf("Student %s (Student ID: %s) has returned the books %s (BookID: %s)\n",
				s.Name, s.ID, book.Title, book.ID)

			// since the book is returned and processed make sure to exit from the function
			return // early exit from the function since the book already processed
		}
	}
	fmt.Printf("Student: %s cant return the BookID: %s\n", s.Name, bookID)

}

// DonateOldBooks Method04
// DonateOldBooks - donate books that have exceeded their lifecycle
// Librarian will donate the books reached their lifecycle
// pointer receiver
func (store *BookStore) DonateOldBooks() []*Book {
	var tmpBookStore []Book      // this will store all the books not reaching their lifecycle
	var bookDonationList []*Book // this will store all teh books reaches their lifecycle and thereby donated by the librarian
	for i := range store.BookList {
		// currentBook -> is actually a currentBook pointer
		currentBook := &store.BookList[i] // get the memory address of the Book struct which holds the actual currentBook content
		if currentBook.LifeCycle <= currentBook.MaxLifeCycle {
			tmpBookStore = append(tmpBookStore, *currentBook) // dereferenced to store the currentBook in the []Book{} slice
		} else {
			currentBook.IsDonated = true // since the lifecycle reached, change the currentBook's Donation status to "true"
			bookDonationList = append(bookDonationList, currentBook)
		}
	}
	// this will globally impact the book list since I am replacing the original booklist of the store withe tmpBookStore which are not in the donation list
	//store.BookList = tmpBookStore // global impact ----FLAG----

	// return the donated books
	return bookDonationList
}

// ListBooks Function-01
// ListBooks - print all the books in the bookstore
// pointer receiver --- store *BookStore ---
func (store *BookStore) ListBooks() {
	fmt.Println("\n\n------ Bookstore Inventory -------")
	for _, book := range store.BookList {
		bookAvailabilityStatus := "Available"
		if book.IsBorrowed {
			bookAvailabilityStatus = fmt.Sprintf("Borrowed by %s", book.BorrowedBy)
		}
		fmt.Printf("BookID: %s, Title: %s, Author: %s, Status: %s, LifeCycle:%d/%d, Donation Status: %v\n",
			book.ID, book.Title, book.Author, bookAvailabilityStatus, book.LifeCycle, book.MaxLifeCycle, book.IsDonated)
	}
}

// ListDonatedBooks Function-02
// ListAllDonatedBooks - print all the donated books
// make sure to avoid making another copy of the "donatedBookList"
// that's why used "passed by reference" to avoid excessive runtime copying
func ListDonatedBooks(donatedBookList []*Book) { //avoid copying the entire book slice locally // instead create a list of pointers to the Books
	// this will make sure changes made within the Book inside is persistent globally
	fmt.Println("\n\n----Donated Book List ----------")
	for _, book := range donatedBookList {
		book.ID = book.ID + "101"
		fmt.Printf("BookID: %s, Title: %s, Author: %s, Lifecycle:%d/%d, Donation Status: %v\n",
			book.ID, book.Title, book.Author, book.LifeCycle, book.MaxLifeCycle, book.IsDonated)
	}
}

// FindStudentByID Function-03
// FindStudentByID - searches for a student by their ID; return NIL if the student doesnt exists
// try to avoid make a local copy of the student slice [] into the function
// make sure to reuse the original student slice [] through (pass by reference)-(aka using pointer)
// (pass by value) function signature would be costly here in terms of both memory and processing time.
// []*Student - a slice of pointer to students
// *[]Student - a pointer to a slice of student
func FindStudentByID(studentList []*Student, id string) *Student {
	// iterating through the slice will give you an index,value
	// since index is not required, that's why we used an empty placeholder (_)
	for _, student := range studentList {
		if student.ID == id { // syntactic sugar applied here; go auto dereferencing (*student).ID
			student.SearchFrequency++
			return student
		}
	}
	return nil // pointer could be nil, that why when the return type is *Student, we can return a nil;
	// if return type was "Student", we either has to return *student (dereferenced) to get the Student struct or an empty Student struct
}

// FindLibrarianByID Function-04
// FindLibrarianByID - searches for a Librarian by their ID; return NIL when the Librarian is not found
func FindLibrarianByID(librarianList []*Librarian, id string) *Librarian {
	for _, currentLibrarian := range librarianList {
		if currentLibrarian.ID == id { // go auto dereferencing // (*librarian).ID==id
			currentLibrarian.SearchFrequency++
			return currentLibrarian
		}
	}
	return nil // return Librarian{}// a pointer can be nil, thats why if that librarian is not found by ID, we can return NIL
	// However, instead of *Librarian return type, if we used "Librarian" return type, then you must have to return an empty Librarian{} struct
}

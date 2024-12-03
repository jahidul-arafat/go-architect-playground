package main

import (
	bs "bookstore-project/bookstore" // <modulename>/<packagename>
	"fmt"
)

func main() {
	fmt.Println("Book Inventory Simulation for better understanding of Call by value and Call By References... ")

	// Create a book store
	// store:=&BookStore{} -> []Book
	// why used &BookStore{} here instead of *BookStore{} ? when the BookStore{} struct has an underneath slice []Book in it?
	/*
		1. BookList (the slice) inside the BookStore is a reference type, but this does not eliminate the need for & when creating a BookStore pointer.
		2. The & operator creates a pointer to the BookStore struct, enabling efficient updates to the BookList and other fields within the struct.
		3.  Without the &, you would work with a copy of the BookStore struct, leading to unintended behavior when modifying fields.
	*/
	store := &bs.BookStore{} // create a BookStore and immediately make a label for this during runtime and store the label into the variable <store>
	// &packagename.BookStore{} // create a pointer to the newly allocated BookStore{}
	fmt.Printf("bookStore Type: %[1]T, Values: %#[1]v\n", store)

	// create an array of Librarians
	// make sure to avoid any local copy, instead referring to the Single Source of Truth of Librarian struct
	// librarians := []*Librarian  // a slice of pointers to Librairan struct
	librarianList := []*bs.Librarian{ // *packagename.Librarian
		{ID: "L1", Name: "Alice"}, // at memory location: 0x140001280c0
		{ID: "L2", Name: "Bob"},   // at memory location: 0x140001280f0
	}

	newLibrarian := &bs.Librarian{ID: "L3", Name: "New Librarian"}
	librarianList = append(librarianList, newLibrarian)
	fmt.Printf("librarians [Type]:%[1]T, Value: %#[1]v\n", librarianList)

	// create an array of Students
	// make sure to avoid any local copy, instead referrign to the Single Source of Truth of <Student> struct
	// create a slice of pointers to the Student struct
	studentList := []*bs.Student{
		{ID: "S1", Name: "Cilly"}, // at memory location: 0x14000192120
		{ID: "S2", Name: "Killy"}, // at memory location: 0x14000192150
	}

	// create a single student
	singleStudent := &bs.Student{ID: "S3", Name: "Single Student"} // create a student and immediate make a label fof this during runtime
	studentList = append(studentList, singleStudent)

	// print all the students

	fmt.Printf("Students [Type]: %[1]T, Value: %#[1]v\n", studentList)

	// let librarian add a few books in the store
	// Method Signature
	// func (l *Librarian) AddBook(store *BookStore, bookID, title, author string, maxLifeCycle int)
	(*librarianList[0]).AddBook(store, "BK1", "Book1", "Title1", 5) // go Semantic Suger applied
	librarianList[0].AddBook(store, "BK2", "Book2", "Title2", 3)
	librarianList[1].AddBook(store, "BK3", "Book3", "Title3", 1)
	librarianList[1].AddBook(store, "BK4", "Book4", "Title4", 1)

	// Print all the books in the Store
	// Method Signature
	// func (store *BookStore) ListBooks()
	store.ListBooks()

	// Let students borrowed the books
	studentList[0].BorrowBook(store, "BK2") // Student Cilly borrowed the book Book2 (BookID: BK2)
	studentList[1].BorrowBook(store, "BK3") // Student Killy borrowed the book Book3 (BookID: BK3)
	studentList[0].BorrowBook(store, "BK2") // not available to borrow
	studentList[1].BorrowBook(store, "BK3") // not available to borrow
	studentList[0].BorrowBook(store, "BK2") // not available to borrow
	studentList[1].BorrowBook(store, "BK3") // not available to borrow

	// List all books after borrowing
	// Now print all the books in the store and check the book availability status
	fmt.Println("\n\nAfter Borrowing Books .....")
	store.ListBooks()

	// Let few students return the books
	studentList[0].ReturnBook(store, "BK2")
	studentList[1].ReturnBook(store, "BK3")
	studentList[2].ReturnBook(store, "BK2")

	// List books after returning
	fmt.Println("\n\nAfter Returning Books .....")
	store.ListBooks()

	// Let students borrow the book BK3 which has a maxLifeCycle=1 and after that the books will be listed for Donating
	fmt.Println("\n\n Book Status to check MaxLifeCycle...")
	//(*studentList[2]).BorrowBook(store, "BK3")
	studentList[2].BorrowBook(store, "BK3")
	store.ListBooks()

	// Donation List after reaching the book MaxLifeCycle
	fmt.Println("\n\n Donated Book List...")
	donationList := store.DonateOldBooks()
	for _, book := range donationList {
		fmt.Printf("%+v\n", book)
	}
	bs.ListDonatedBooks(donationList)

	// check the changes in the donation book ID is persistent or not?
	fmt.Println("Checking if the Changes insuide the DonatedBookList printing is persistent....")
	// using for loop to check the donationList
	for _, book := range donationList {
		fmt.Printf("%+v\n", book)
	}

	// List all books after donating
	fmt.Println("\n\nAfter Donating Books.....")
	store.ListBooks()

	// Find a Librarian By ID
	fmt.Println("\n\nSearching Librarian by ID.....")
	fmt.Println("Librairan: ", *bs.FindLibrarianByID(librarianList, "L1")) // dereferencing
	fmt.Println("Librairan: ", *bs.FindLibrarianByID(librarianList, "L1")) // dereferencing
	fmt.Println("Librairan: ", bs.FindLibrarianByID(librarianList, "L5"))  // dereferencing of NIL is not allowed

	// Find a Student By ID
	fmt.Println("\n\nSearching Student by ID....")
	fmt.Println("Student: ", *bs.FindStudentByID(studentList, "S3"))
	fmt.Println("Student: ", *bs.FindStudentByID(studentList, "S3"))
	fmt.Println("Student: ", *bs.FindStudentByID(studentList, "S3"))

}

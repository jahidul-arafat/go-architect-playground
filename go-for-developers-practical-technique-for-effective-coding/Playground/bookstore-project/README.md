## Basic Architecture
```
bookstore/
├── types.go         # Contains type definitions (e.g., structs for Book, BookStore, Librarian, and Student)
├── logic.go         # Contains business logic methods (e.g., AddBook, BorrowBook, ReturnBook, DonateOldBooks)
├── main.go          # Entry point that initializes data and calls functions
└── go.mod           # Go module file
```
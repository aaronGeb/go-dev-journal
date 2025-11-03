package services

import (
	"errors"
	"library_management/models"
)

// LibraryManager interface
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

// Library struct implements LibraryManager
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// NewLibrary creates and returns a new Library instance
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// AddBook adds a new book
func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book
}

// RemoveBook removes a book by ID
func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

// BorrowBook allows a member to borrow an available book
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

// ReturnBook allows a member to return a borrowed book
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	found := false
	newBorrowed := []models.Book{}
	for _, b := range member.BorrowedBooks {
		if b.ID == bookID {
			found = true
			continue
		}
		newBorrowed = append(newBorrowed, b)
	}

	if !found {
		return errors.New("book not borrowed by this member")
	}

	member.BorrowedBooks = newBorrowed
	l.Members[memberID] = member
	book.Status = "Available"
	l.Books[bookID] = book
	return nil
}

// ListAvailableBooks returns all available books
func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

// ListBorrowedBooks returns all borrowed books of a member
func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := l.Members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}

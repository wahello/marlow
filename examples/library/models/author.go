package models

//go:generate marlowc -stdout=true

// Example file for testing.

// Author represents an author of a book.
type Author struct {
	_    bool   `marlow:"tableName=authors"`
	Name string `marlow:"column=name"`
}
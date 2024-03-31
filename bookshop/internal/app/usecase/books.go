package usecase

import (
	"context"

	"bookshop.com/internal/domain/entities"
)

type BookRepository interface {
	AddBook(ctx context.Context, book entities.Book) error
	GetBooks(ctx context.Context) ([]entities.Book, error)
}

type BooksInteractor struct {
	Repo BookRepository
}

// Añadir un libro
func (li *BooksInteractor) NewBook(ctx context.Context, title, author string, year int) error {
	book := entities.Book{Title: title, Author: author, Year: year}
	return li.Repo.AddBook(ctx, book)
}

func (li *BooksInteractor) GetBooks(ctx context.Context) ([]entities.Book, error) {
	return li.GetBooks(ctx)
}

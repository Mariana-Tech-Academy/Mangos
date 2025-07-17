package repository

import (
	"library/database"
	"library/models"

	"github.com/google/uuid"
)

type BookRepository interface {
	GetBooksByUser(id string) (*[]models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(book *models.Book) error
	CreateBook(book *models.Book) error
    DeleteBook(id uuid.UUID) error
}

type BookRepo struct {
}

func (r *BookRepo) CreateBook(book *models.Book) error {
    return database.Db.Create(book).Error
}

func (r *BookRepo) UpdateBook(book *models.Book) error {
    err := database.Db.Save(book).Error
	return err
}

func (r *BookRepo) DeleteBook(id uuid.UUID) error {
    return database.Db.Delete(&models.Book{}, id).Error
}

func (r *BookRepo) GetBooksByUser(id string) (*[]models.Book, error) {
	var user models.User
	err := database.Db.Where("ID = ?", id).First(&user).Error
	return &user.Books, err
}

func (r *BookRepo) GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	err := database.Db.Where("ID = ?", id).First(&book).Error
	return &book, err
}


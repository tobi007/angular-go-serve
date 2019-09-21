package repository

import "github.com/tobi007/angular-go-serve/models"

// UserRepo Interface
type UserRepo interface {
	Create( u *models.User) (int64, error)
	Update( u *models.User) (*models.User, error)
	RetrieveByID (id int64) (*models.User, error)
	RetrieveByEmail (email string) (*models.User, error)
	ExistByID(email string) bool
}


package ports

import (
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/domain"
)

// UserRepository is an interface for interacting with user-related data
type UserRepository interface {
	SelectAll() ([]domain.User, error)
	Insert(input domain.User) error
	Login(email string, password string) (domain.User, string, error)
	SelectById(id uint) (domain.User, error)
	SelectByEmail(email string) (domain.User, error)
	Update(id uint, input domain.User) (data domain.User, err error)
	Delete(id uint) (row int, err error)
}

// UserService is an interface for interacting with user-related business logic
type UserService interface {
	GetAll() ([]domain.User, error)
	Create(input domain.User) error
	Login(email string, password string) (domain.User, string, error)
	GetById(id uint) (domain.User, error)
	Update(id uint, input domain.User) (data domain.User, err error)
	Delete(id uint) (err error)
}

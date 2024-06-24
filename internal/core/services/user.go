package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/domain"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/pkg/encrypt"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/ports"
)

/**
 * UserService implements port.UserService interface
 * and provides an access to the user repository
 */
type UserService struct {
	repo        ports.UserRepository
	hashService encrypt.HashInterface
	validate    *validator.Validate
}

// NewUserService creates a new user service instance
func NewUserService(repo ports.UserRepository, hash encrypt.HashInterface) ports.UserService {
	return &UserService{
		repo:        repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Create implements ports.UserService.
func (u *UserService) Create(input domain.User) error {
	// validasi /logic
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("[validation] nama/email/password tidak boleh kosong")
	// }
	errValidate := u.validate.Struct(input)
	if errValidate != nil {
		return errors.New("[validation] " + errValidate.Error())
	}

	if input.Password != "" {
		// proses hash password
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errHash
		}
		input.Password = result
	}
	// memanggil method di data layer
	err := u.repo.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements ports.UserService.
func (u *UserService) Delete(id uint) (err error) {
	panic("unimplemented")
}

// GetAll implements ports.UserService.
func (u *UserService) GetAll() ([]domain.User, error) {
	panic("unimplemented")
}

// GetById implements ports.UserService.
func (u *UserService) GetById(id uint) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, errors.New("[validation] id not valid")
	}
	return u.repo.SelectById(id)
}

// Login implements ports.UserService.
func (u *UserService) Login(email string, password string) (domain.User, string, error) {
	panic("unimplemented")
}

// Update implements ports.UserService.
func (u *UserService) Update(id uint, input domain.User) (data domain.User, err error) {
	panic("unimplemented")
}

package repositories

import (
	"github.com/iffakhry/go-commerce-hexagonal/internal/adapter/storages/postgres/models"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/domain"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/ports"
	"gorm.io/gorm"
)

/**
 * UserRepository implements port.UserRepository interface
 * and provides an access to the postgres database
 */
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &UserRepository{
		db,
	}
}

// Delete implements ports.UserRepository.
func (u *UserRepository) Delete(id uint) (row int, err error) {
	panic("unimplemented")
}

// Insert implements ports.UserRepository.
func (u *UserRepository) Insert(input domain.User) error {
	var userGorm models.User = models.UserEntityToModel(input)

	// userGorm = User{
	// 	Name:      input.Name,
	// 	Email:     input.Email,
	// 	Password:  input.Password,
	// 	Phone:     input.Phone,
	// 	Address:   input.Address,
	// 	StoreName: input.StoreName,
	// }
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Login implements ports.UserRepository.
func (u *UserRepository) Login(email string, password string) (domain.User, string, error) {
	panic("unimplemented")
}

// SelectAll implements ports.UserRepository.
func (u *UserRepository) SelectAll() ([]domain.User, error) {
	panic("unimplemented")
}

// SelectByEmail implements ports.UserRepository.
func (u *UserRepository) SelectByEmail(email string) (domain.User, error) {
	panic("unimplemented")
}

// SelectById implements ports.UserRepository.
func (u *UserRepository) SelectById(id uint) (domain.User, error) {
	panic("unimplemented")
}

// Update implements ports.UserRepository.
func (u *UserRepository) Update(id uint, input domain.User) (data domain.User, err error) {
	panic("unimplemented")
}

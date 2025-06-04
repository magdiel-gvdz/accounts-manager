package ports

import "github.com/Magdiel-GVdz/accounts-manager/internal/domain"

// UserRepository defines the expected behaviour from a repository
// providing persistence for Users.
type UserRepository interface {
	Create(user *domain.User) error
	Get(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
	List() ([]domain.User, error)
	GetByUsername(username string) (*domain.User, error)
}

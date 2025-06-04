package application

import (
	"github.com/Magdiel-GVdz/accounts-manager/internal/domain"
	"github.com/Magdiel-GVdz/accounts-manager/internal/domain/ports"
	"golang.org/x/crypto/bcrypt"
)

// UserService holds business logic for users.
type UserService struct {
	repo ports.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(r ports.UserRepository) *UserService {
	return &UserService{repo: r}
}

// Create registers a new user hashing its password.
func (s *UserService) Create(u *domain.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return s.repo.Create(u)
}

// Get retrieves a user by id.
func (s *UserService) Get(id uint) (*domain.User, error) { return s.repo.Get(id) }

// Update modifies an existing user.
func (s *UserService) Update(u *domain.User) error { return s.repo.Update(u) }

// Delete removes a user by id.
func (s *UserService) Delete(id uint) error { return s.repo.Delete(id) }

// List returns all users.
func (s *UserService) List() ([]domain.User, error) { return s.repo.List() }

// Authenticate validates user credentials.
func (s *UserService) Authenticate(username, password string) (*domain.User, error) {
	u, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}
	return u, nil
}

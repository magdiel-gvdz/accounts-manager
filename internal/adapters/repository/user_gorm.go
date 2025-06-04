package repository

import (
	"github.com/Magdiel-GVdz/accounts-manager/internal/domain"
	"gorm.io/gorm"
)

// GormUserRepository implements ports.UserRepository using GORM.
type GormUserRepository struct{ db *gorm.DB }

// NewGormUserRepository creates a new repository backed by GORM.
func NewGormUserRepository(db *gorm.DB) *GormUserRepository { return &GormUserRepository{db: db} }

func (r *GormUserRepository) Create(u *domain.User) error { return r.db.Create(u).Error }

func (r *GormUserRepository) Get(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Update(u *domain.User) error { return r.db.Save(u).Error }

func (r *GormUserRepository) Delete(id uint) error { return r.db.Delete(&domain.User{}, id).Error }

func (r *GormUserRepository) List() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

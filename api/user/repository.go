package user

import (
	"github.com/Ze-Victor/taste-match-api/taste-match-api/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (b UserRepositoryImpl) Find(id int) (*entities.User, error) {
	var user *entities.User

	result := b.DB.Preload("Preferences").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r UserRepositoryImpl) FindAllWithPreferences() ([]entities.User, error) {
	var users []entities.User

	result := r.DB.Preload("Preferences").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (b UserRepositoryImpl) Update(c User) (*User, error) {
	// TODO impl this
	return nil, nil
}

func (r *UserRepositoryImpl) Create(user *entities.User) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})
}

func (b UserRepositoryImpl) Delete(c User) error {
	// TODO impl this
	return nil
}

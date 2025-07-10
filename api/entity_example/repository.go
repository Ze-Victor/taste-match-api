package entity_example

import (
	"github.com/Ze-Victor/taste-match-api/taste-match-api/entities"
	"gorm.io/gorm"
)

type EntityExampleRepositoryImpl struct {
	DB *gorm.DB
}

func NewEntityExampleRepository(db *gorm.DB) *EntityExampleRepositoryImpl {
	return &EntityExampleRepositoryImpl{DB: db}
}

func (b EntityExampleRepositoryImpl) Find() (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (r EntityExampleRepositoryImpl) FindAllWithPreferences() ([]entities.User, error) {
	var users []entities.User

	result := r.DB.Preload("Preferences").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (b EntityExampleRepositoryImpl) Update(c EntityExample) (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleRepositoryImpl) Create(c EntityExample) (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleRepositoryImpl) Delete(c EntityExample) error {
	// TODO impl this
	return nil
}

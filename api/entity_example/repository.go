package entity_example

import (
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

package entity_example

import (
	entity_example_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/entity_example/dto"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/entities"
)

type EntityExampleRepository interface {
	Find() (*EntityExample, error)
	FindAllWithPreferences() ([]entities.User, error)
	Update(c EntityExample) (*EntityExample, error)
	Create(c EntityExample) (*EntityExample, error)
	Delete(c EntityExample) error
}

type EntityExampleBusiness interface {
	Find() (*EntityExample, error)
	FindAll() ([]entity_example_dto.UserResponse, error)
	Update(c EntityExample) (*EntityExample, error)
	Create(c EntityExample) (*EntityExample, error)
	Delete(c EntityExample) error
}

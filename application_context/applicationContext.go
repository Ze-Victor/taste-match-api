package application_context

import (
	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/entity_example"
	"gorm.io/gorm"
)

type ApplicationContext struct {
	ExampleController *entity_example.EntityExampleController
	ExampleBusiness   *entity_example.EntityExampleBusiness
	ExampleRepository *entity_example.EntityExampleRepository
}

func NewApplicationContext(db *gorm.DB) *ApplicationContext {
	exampleRepository := entity_example.NewEntityExampleRepository(db)
	exampleBusiness := entity_example.NewEntityExampleBusinessImpl(exampleRepository)
	exampleController := entity_example.NewEntityExampleController(exampleBusiness)

	return &ApplicationContext{
		ExampleController: exampleController,
	}
}

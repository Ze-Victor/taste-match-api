package application_context

import (
	"gitlab.luizalabs.com/taste-match-api/api/entity_example"
)

type ApplicationContext struct {
	ExampleController *entity_example.EntityExampleController
	ExampleBusiness   *entity_example.EntityExampleBusiness
	ExampleRepository *entity_example.EntityExampleRepository
}

func NewApplicationContext() *ApplicationContext {
	exampleRepository := entity_example.NewEntityExampleRepository()
	exampleBusiness := entity_example.NewEntityExampleBusinessImpl(exampleRepository)
	exampleController := entity_example.NewEntityExampleController(exampleBusiness)

	return &ApplicationContext{
		ExampleController: exampleController,
	}
}

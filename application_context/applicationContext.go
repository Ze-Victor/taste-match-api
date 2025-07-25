package application_context

import (
	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/entity_example"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/preferences"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/user"
	"gorm.io/gorm"
)

type ApplicationContext struct {
	ExampleController *entity_example.EntityExampleController
	ExampleBusiness   *entity_example.EntityExampleBusiness
	ExampleRepository *entity_example.EntityExampleRepository

	UserController *user.UserController
	UserBusiness   *user.UserBusiness
	UserRepository *user.UserRepository

	PreferencesController *preferences.PreferencesController
	PreferencesBusiness   *preferences.PreferencesBusiness
	PreferencesRepository *preferences.PreferencesRepository
}

func NewApplicationContext(db *gorm.DB) *ApplicationContext {
	exampleRepository := entity_example.NewEntityExampleRepository(db)
	exampleBusiness := entity_example.NewEntityExampleBusinessImpl(exampleRepository)
	exampleController := entity_example.NewEntityExampleController(exampleBusiness)

	preferencesRepository := preferences.NewPreferencesRepository(db)
	preferencesBusiness := preferences.NewPreferencesBusinessImpl(preferencesRepository)
	preferencesController := preferences.NewPreferencesController(preferencesBusiness)

	userRepository := user.NewUserRepository(db)
	userBusiness := user.NewUserBusinessImpl(userRepository, preferencesRepository)
	userController := user.NewUserController(userBusiness)

	return &ApplicationContext{
		ExampleController:     exampleController,
		UserController:        userController,
		PreferencesController: preferencesController,
	}
}

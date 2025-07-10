package entity_example

import entity_example_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/entity_example/dto"

type EntityExampleBusinessImpl struct {
	entityExampleRepository EntityExampleRepository
}

func NewEntityExampleBusinessImpl(entityExampleRepository EntityExampleRepository) *EntityExampleBusinessImpl {
	return &EntityExampleBusinessImpl{entityExampleRepository}
}

func (b EntityExampleBusinessImpl) Find() (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleBusinessImpl) FindAll() ([]entity_example_dto.UserResponse, error) {
	usersFromDb, err := b.entityExampleRepository.FindAllWithPreferences()
	if err != nil {
		return nil, err
	}

	var response []entity_example_dto.UserResponse

	for _, user := range usersFromDb {
		var userPrefs []entity_example_dto.PreferenceResponse
		for _, pref := range user.Preferences {
			userPrefs = append(userPrefs, entity_example_dto.PreferenceResponse{
				Description: pref.Description,
			})
		}

		response = append(response, entity_example_dto.UserResponse{
			Name:        user.Name,
			BirthDate:   user.BirthDate.Format("2006-01-02"),
			Preferences: userPrefs,
		})
	}

	return response, nil
}

func (b EntityExampleBusinessImpl) Update(c EntityExample) (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleBusinessImpl) Create(c EntityExample) (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleBusinessImpl) Delete(c EntityExample) error {
	// TODO impl this
	return nil
}

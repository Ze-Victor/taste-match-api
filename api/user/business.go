package user

import user_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/user/dto"

type UserBusinessImpl struct {
	UserRepository UserRepository
}

func NewUserBusinessImpl(UserRepository UserRepository) *UserBusinessImpl {
	return &UserBusinessImpl{UserRepository}
}

func (b UserBusinessImpl) Find(id int) (*user_dto.UserResponse, error) {
	userFromDb, err := b.UserRepository.Find(id)
	if err != nil {
		return nil, err
	}

	var userPrefs []user_dto.PreferenceResponse
	for _, pref := range userFromDb.Preferences {
		userPrefs = append(userPrefs, user_dto.PreferenceResponse{
			Description: pref.Description,
		})
	}

	userResponse := &user_dto.UserResponse{
		Name:        userFromDb.Name,
		BirthDate:   userFromDb.BirthDate.Format("2006-01-02"),
		Preferences: userPrefs,
	}

	return userResponse, nil
}

func (b UserBusinessImpl) FindAll() ([]user_dto.UserResponse, error) {
	usersFromDb, err := b.UserRepository.FindAllWithPreferences()
	if err != nil {
		return nil, err
	}

	var response []user_dto.UserResponse

	for _, user := range usersFromDb {
		var userPrefs []user_dto.PreferenceResponse
		for _, pref := range user.Preferences {
			userPrefs = append(userPrefs, user_dto.PreferenceResponse{
				Description: pref.Description,
			})
		}

		response = append(response, user_dto.UserResponse{
			Name:        user.Name,
			BirthDate:   user.BirthDate.Format("2006-01-02"),
			Preferences: userPrefs,
		})
	}

	return response, nil
}

func (b UserBusinessImpl) Update(c User) (*User, error) {
	// TODO impl this
	return nil, nil
}

func (b UserBusinessImpl) Create(c User) (*User, error) {
	// TODO impl this
	return nil, nil
}

func (b UserBusinessImpl) Delete(c User) error {
	// TODO impl this
	return nil
}

package user

import (
	"errors"

	user_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/user/dto"
	bcrypt "golang.org/x/crypto/bcrypt"

	"github.com/Ze-Victor/taste-match-api/taste-match-api/entities"

	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/preferences"
)

type UserBusinessImpl struct {
	UserRepository       UserRepository
	PreferenceRepository preferences.PreferencesRepository
}

func NewUserBusinessImpl(UserRepository UserRepository, PreferenceRepository preferences.PreferencesRepository) *UserBusinessImpl {
	return &UserBusinessImpl{UserRepository, PreferenceRepository}
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

func (b UserBusinessImpl) CalculateMatch(userID1, userID2 int) (float64, error) {
	user1, err := b.UserRepository.Find(userID1)
	if err != nil {
		return 0, err
	}

	user2, err := b.UserRepository.Find(userID2)
	if err != nil {
		return 0, err
	}

	score := calculateMatchScore(user1.Preferences, user2.Preferences)

	return score * 100, nil
}

func (b UserBusinessImpl) Update(c User) (*User, error) {
	// TODO impl this
	return nil, nil
}

func (b UserBusinessImpl) Create(request user_dto.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var preferencesFromDB []entities.Preference
	if len(request.Preference) > 0 {
		prefs, err := b.PreferenceRepository.FindByIDs(request.Preference)
		if err != nil {
			return err
		}
		if len(prefs) != len(request.Preference) {
			return errors.New("uma ou mais preferências não foram encontradas")
		}
		preferencesFromDB = prefs
	}

	userToCreate := &entities.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    string(hashedPassword),
		Gender:      request.Gender,
		BirthDate:   request.BirthDate,
		Preferences: preferencesFromDB,
	}

	if err := b.UserRepository.Create(userToCreate); err != nil {
		return err
	}

	return nil
}

func (b UserBusinessImpl) Delete(c User) error {
	// TODO impl this
	return nil
}

func calculateMatchScore(preferencesA, preferencesB []entities.Preference) float64 {
	setA := make(map[uint]bool)
	for _, pref := range preferencesA {
		setA[pref.ID] = true
	}

	intersectionSize := 0
	for _, pref := range preferencesB {
		if setA[pref.ID] {
			intersectionSize++
		}
	}

	unionSize := len(preferencesA) + len(preferencesB) - intersectionSize

	if unionSize == 0 {
		if len(preferencesA) == 0 && len(preferencesB) == 0 {
			return 1.0
		}
		return 0.0
	}

	return float64(intersectionSize) / float64(unionSize)
}

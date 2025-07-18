package preferences

import (
	preferences_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/preferences/dto"
	"gorm.io/gorm"
)

type PreferencesRepositoryImpl struct {
	DB *gorm.DB
}

func NewPreferencesRepository(db *gorm.DB) *PreferencesRepositoryImpl {
	return &PreferencesRepositoryImpl{DB: db}
}

func (r PreferencesRepositoryImpl) FindAll() ([]preferences_dto.Preferences, error) {
	var preferences []preferences_dto.Preferences

	result := r.DB.Find(&preferences)

	if result.Error != nil {
		return nil, result.Error
	}

	return preferences, nil
}

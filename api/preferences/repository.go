package preferences

import (
	preferences_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/preferences/dto"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/entities"
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

func (r *PreferencesRepositoryImpl) FindByID(id uint) (*entities.Preference, error) {
	var preference entities.Preference
	if err := r.DB.First(&preference, id).Error; err != nil {
		return nil, err
	}
	return &preference, nil
}

func (r *PreferencesRepositoryImpl) FindByIDs(ids []uint) ([]entities.Preference, error) {
	var preferences []entities.Preference
	if err := r.DB.Where("id IN ?", ids).Find(&preferences).Error; err != nil {
		return nil, err
	}
	return preferences, nil
}

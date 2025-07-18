package preferences

import (
	preferences_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/preferences/dto"
	entities "github.com/Ze-Victor/taste-match-api/taste-match-api/entities"
)

type PreferencesRepository interface {
	FindAll() ([]preferences_dto.Preferences, error)
	FindByID(id uint) (*entities.Preference, error)
	FindByIDs(ids []uint) ([]entities.Preference, error)
}

type PreferencesBusiness interface {
	FindAll() ([]preferences_dto.Preferences, error)
	FindByID(id uint) (*entities.Preference, error)
	FindByIDs(ids []uint) ([]entities.Preference, error)
}

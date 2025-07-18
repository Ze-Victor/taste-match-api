package preferences

import (
	preferences_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/preferences/dto"
)

type PreferencesRepository interface {
	FindAll() ([]preferences_dto.Preferences, error)
}

type PreferencesBusiness interface {
	FindAll() ([]preferences_dto.Preferences, error)
}

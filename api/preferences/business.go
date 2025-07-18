package preferences

import preferences_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/preferences/dto"

type PreferencesBusinessImpl struct {
	PreferencesRepository PreferencesRepository
}

func NewPreferencesBusinessImpl(PreferencesRepository PreferencesRepository) *PreferencesBusinessImpl {
	return &PreferencesBusinessImpl{PreferencesRepository}
}

func (b PreferencesBusinessImpl) FindAll() ([]preferences_dto.Preferences, error) {
	preferencesFromDb, err := b.PreferencesRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return preferencesFromDb, nil
}

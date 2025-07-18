package user_dto

type PreferenceResponse struct {
	Description string `json:"description"`
}
type UserResponse struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	Gender      string               `json:"gender"`
	BirthDate   string               `json:"birth_date"`
	Preferences []PreferenceResponse `json:"preferences"`
}

package user_dto

type FindEntityExampleResponse struct {
	ID   int    `json:"id" example:"50"`
	Name string `json:"name" example:"fulano"`
}

type FindAllEntityExamplesResponse struct {
	Name string `json:"name" example:"fulano"`
}

type PreferenceResponse struct {
	Description string `json:"description"`
}
type UserResponse struct {
	Name        string               `json:"name"`
	BirthDate   string               `json:"birth_date"`
	Preferences []PreferenceResponse `json:"preferences"`
}

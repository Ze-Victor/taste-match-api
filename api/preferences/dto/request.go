package preferences_dto

type CreateEntityExampleRequest struct {
	Name string `json:"name" validate:"required" example:"fulano"`
}

type PatchEntityExampleRequest struct {
	Name string `json:"name" validate:"required" example:"fulano"`
}

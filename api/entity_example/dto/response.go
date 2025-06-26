package entity_example_dto

type FindEntityExampleResponse struct {
	ID   int    `json:"id" example:"50"`
	Name string `json:"name" example:"fulano"`
}

type FindAllEntityExamplesResponse struct {
	Name string `json:"name" example:"fulano"`
}

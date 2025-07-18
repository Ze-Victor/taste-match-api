package user_dto

import (
	"time"
)

type CreateUserRequest struct {
	Name       string    `json:"name" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Password   string    `json:"password" binding:"required,min=8"`
	Gender     string    `json:"gender" binding:"required,oneof=M F"`
	BirthDate  time.Time `json:"birth_date" binding:"required"`
	Preference []uint    `json:"preferences" binding:"required"`
}

type PatchEntityExampleRequest struct {
	Name string `json:"name" validate:"required" example:"fulano"`
}

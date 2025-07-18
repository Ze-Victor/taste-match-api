package user

import (
	user_dto "github.com/Ze-Victor/taste-match-api/taste-match-api/api/user/dto"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/entities"
)

type UserRepository interface {
	Find(id int) (*entities.User, error)
	FindAllWithPreferences() ([]entities.User, error)
	Update(c User) (*User, error)
	Create(c User) (*User, error)
	Delete(c User) error
}

type UserBusiness interface {
	Find(id int) (*user_dto.UserResponse, error)
	FindAll() ([]user_dto.UserResponse, error)
	CalculateMatch(userID1, userID2 int) (float64, error)
	Update(c User) (*User, error)
	Create(c User) (*User, error)
	Delete(c User) error
}

package pg

import (
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func toUserFromDB(in User) (model.User, error) {
	userID, err := model.NewUserID(in.ID)
	if err != nil {
		return model.User{}, fmt.Errorf("convert str to user id: %w", err)
	}

	role, err := model.NewRole(in.Role)
	if err != nil {
		return model.User{}, fmt.Errorf("convert str role to model: %w", err)
	}

	user := model.User{
		ID:    userID,
		Email: in.Email,
		Role:  role,

		CreatedAt: in.CreatedAt,
	}

	return user, nil
}

package dto

import "github.com/jacky-htg/brilian/models"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *CreateUserRequest) ToEntity() models.User {
	return models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

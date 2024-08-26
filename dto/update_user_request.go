package dto

import "github.com/jacky-htg/brilian/models"

type UpdateUserRequest struct {
	Name string `json:"name"`
}

func (u *UpdateUserRequest) ToEntity() models.User {
	return models.User{
		Name: u.Name,
	}
}

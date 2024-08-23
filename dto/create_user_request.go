package dto

import "github.com/jacky-htg/brilian/models"

type CreateUserRequest struct {
	Name string `json:"name"`
	Loc  string `json:"loc"`
}

func (u *CreateUserRequest) ToEntity() models.User {
	return models.User{
		Name: u.Name,
		Loc:  u.Loc,
	}
}

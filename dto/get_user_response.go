package dto

import "github.com/jacky-htg/brilian/models"

type GetUserResponse struct {
	Id    int    `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *GetUserResponse) FromEntity(e models.User) {
	u.Id = int(e.Id)
	u.Name = e.Name
	u.Email = e.Email
}

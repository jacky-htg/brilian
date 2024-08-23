package repositories

import (
	"database/sql"
	"log"

	"github.com/jacky-htg/brilian/models"
)

// User : struct of User
type UserRepository struct {
	Db         *sql.DB
	Log        *log.Logger
	UserEntity models.User
}

// ListUsers : http handler for returning list of users
func (u *UserRepository) Find(id int) error {
	const q = `SELECT id, name, loc FROM instructors WHERE id=?`
	err := u.Db.QueryRow(q, id).Scan(&u.UserEntity.Id, &u.UserEntity.Name, &u.UserEntity.Loc)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}
	return nil
}

func (u *UserRepository) Save() error {
	const q = `INSERT INTO instructors (name, loc) VALUES (?, ?)`
	rs, err := u.Db.Exec(q, u.UserEntity.Name, u.UserEntity.Loc)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}

	id, _ := rs.LastInsertId()
	u.UserEntity.Id = uint(id)
	return nil
}

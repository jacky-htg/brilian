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
	const q = `SELECT id, name, email FROM users WHERE id=?`
	err := u.Db.QueryRow(q, id).Scan(&u.UserEntity.Id, &u.UserEntity.Name, &u.UserEntity.Email)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}
	return nil
}

func (u *UserRepository) Save() error {
	const q = `INSERT INTO users (name, password, email, created_by, updated_by) VALUES (?, ?, ?, ?, ?)`
	rs, err := u.Db.Exec(q,
		u.UserEntity.Name,
		u.UserEntity.Password,
		u.UserEntity.Email,
		1,
		1,
	)
	if err != nil {
		u.Log.Println("error create users", err)
		return err
	}

	id, _ := rs.LastInsertId()
	u.UserEntity.Id = int(id)
	return nil
}

func (u *UserRepository) Update() error {
	const q = `UPDATE users SET name = ?, updated_at = NOW(), updated_by = ? WHERE id = ?`
	_, err := u.Db.Exec(q, u.UserEntity.Name, 1, u.UserEntity.Id)
	if err != nil {
		u.Log.Println("error update users", err)
		return err
	}

	return nil
}

func (u *UserRepository) Delete() error {
	const q = `DELETE FROM users WHERE id = ?`
	_, err := u.Db.Exec(q, u.UserEntity.Id)
	if err != nil {
		u.Log.Println("error delete users", err)
		return err
	}

	return nil
}

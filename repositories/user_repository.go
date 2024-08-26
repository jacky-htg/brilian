package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/jacky-htg/brilian/models"
)

// User : struct of User
type UserRepository struct {
	Db         *sql.DB
	Log        *log.Logger
	UserEntity models.User
}

func (u *UserRepository) Find(ctx context.Context, id int) error {
	switch ctx.Err() {
	case context.Canceled:
		return errors.New("request is canceled")
	case context.DeadlineExceeded:
		return errors.New("deadline is exceeded")
	default:
	}

	//time.Sleep(5 * time.Second)

	const q = `SELECT id, name, email FROM users WHERE id=?`
	stmt, err := u.Db.PrepareContext(ctx, q)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, id).Scan(&u.UserEntity.Id, &u.UserEntity.Name, &u.UserEntity.Email)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}
	return nil
}

func (u *UserRepository) Save(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return errors.New("request is canceled")
	case context.DeadlineExceeded:
		return errors.New("deadline is exceeded")
	default:
	}

	const q = `INSERT INTO users (name, password, email, created_by, updated_by) VALUES (?, ?, ?, ?, ?)`
	stmt, err := u.Db.PrepareContext(ctx, q)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}
	defer stmt.Close()

	rs, err := stmt.ExecContext(
		ctx,
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

func (u *UserRepository) Update(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return errors.New("request is canceled")
	case context.DeadlineExceeded:
		return errors.New("deadline is exceeded")
	default:
	}

	const q = `UPDATE users SET name = ?, updated_at = NOW(), updated_by = ? WHERE id = ?`
	stmt, err := u.Db.PrepareContext(ctx, q)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.UserEntity.Name, 1, u.UserEntity.Id)
	if err != nil {
		u.Log.Println("error update users", err)
		return err
	}

	return nil
}

func (u *UserRepository) Delete(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return errors.New("request is canceled")
	case context.DeadlineExceeded:
		return errors.New("deadline is exceeded")
	default:
	}

	const q = `DELETE FROM users WHERE id = ?`
	stmt, err := u.Db.PrepareContext(ctx, q)
	if err != nil {
		u.Log.Println("error get users", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.UserEntity.Id)
	if err != nil {
		u.Log.Println("error delete users", err)
		return err
	}

	return nil
}

func (u *UserRepository) List(ctx context.Context, search string) ([]models.User, error) {
	var list []models.User

	switch ctx.Err() {
	case context.Canceled:
		return list, errors.New("request is canceled")
	case context.DeadlineExceeded:
		return list, errors.New("deadline is exceeded")
	default:
	}

	sb := strings.Builder{}
	sb.WriteString(`SELECT id, name, email FROM users`)
	var args []interface{}

	if len(search) > 0 {
		sb.WriteString(` WHERE name like ?`)
		args = append(args, `%`+search+`%`)
	}
	stmt, err := u.Db.PrepareContext(ctx, sb.String())
	if err != nil {
		u.Log.Println("error get users", err)
		return list, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		u.Log.Println("error get users", err)
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			u.Log.Println("error get users", err)
			return list, err
		}
		list = append(list, user)
	}

	if rows.Err() != nil {
		u.Log.Println("error get users", err)
		return list, err
	}

	return list, nil
}

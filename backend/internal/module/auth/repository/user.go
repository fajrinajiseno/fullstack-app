package repository

import (
	"database/sql"

	"github.com/fajrinajiseno/mygolangapp/internal/entity"
)

//go:generate mockgen -source user.go -destination mock/user_mock.go -package=mock
type UserRepository interface {
	Save(u *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

type User struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *User {
	return &User{db: db}
}

func (r *User) Save(u *entity.User) error {
	_, err := r.db.Exec(`INSERT INTO users (email, password_hash)
		VALUES (?, ?)`,
		u.Email, u.PasswordHash)
	if err != nil {
		return entity.WrapError(err, entity.ErrorCodeInternal, err.Error())
	}
	return nil
}

func (r *User) GetUserByEmail(email string) (*entity.User, error) {
	row := r.db.QueryRow(`SELECT id, email, password_hash, created_at FROM users WHERE email = ?`, email)
	var u entity.User
	if err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrorNotFound("user not found")
		}
		return nil, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}
	return &u, nil
}

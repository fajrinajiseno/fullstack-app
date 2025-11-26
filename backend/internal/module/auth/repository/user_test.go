package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajrinajiseno/mygolangapp/internal/entity"
	"github.com/stretchr/testify/assert"
)

func newMockUserRepo(t *testing.T) (*User, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	repo := NewUserRepo(db)
	cleanup := func() { db.Close() }
	return repo, mock, cleanup
}

func TestGetUserByEmail_Success(t *testing.T) {
	repo, mock, cleanup := newMockUserRepo(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"id", "email", "password_hash", "created_at"}).
		AddRow(1, "alice@example.com", "$2a$hash", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, email, password_hash, created_at FROM users WHERE email = ?")).
		WithArgs("alice@example.com").
		WillReturnRows(rows)

	u, err := repo.GetUserByEmail("alice@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, 1, u.ID)
	assert.Equal(t, "alice@example.com", u.Email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unfulfilled expectations: %v", err)
	}
}

func TestGetUserByEmail_NotFound(t *testing.T) {
	repo, mock, cleanup := newMockUserRepo(t)
	defer cleanup()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, email, password_hash, created_at FROM users WHERE email = ?")).
		WithArgs("missing@example.com").
		WillReturnError(sql.ErrNoRows)

	u, err := repo.GetUserByEmail("missing@example.com")
	assert.Nil(t, u)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unfulfilled expectations: %v", err)
	}
}

func TestGetUserByEmail_DBError(t *testing.T) {
	repo, mock, cleanup := newMockUserRepo(t)
	defer cleanup()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, email, password_hash, created_at FROM users WHERE email = ?")).
		WithArgs("err@example.com").
		WillReturnError(errors.New("db fail"))

	u, err := repo.GetUserByEmail("err@example.com")
	assert.Nil(t, u)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "db error")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unfulfilled expectations: %v", err)
	}
}
func TestSave_Success(t *testing.T) {
	repo, mock, cleanup := newMockUserRepo(t)
	defer cleanup()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (email, password_hash) VALUES (?, ?)")).
		WithArgs("carol@example.com", "password").
		WillReturnResult(sqlmock.NewResult(1, 1))

	userToSave := &entity.User{
		Email:        "carol@example.com",
		PasswordHash: "password",
	}

	err := repo.Save(userToSave)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unfulfilled expectations: %v", err)
	}
}

func TestSave_DBError(t *testing.T) {
	repo, mock, cleanup := newMockUserRepo(t)
	defer cleanup()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (email, password_hash) VALUES (?, ?)")).
		WithArgs("carol@example.com", "password").
		WillReturnError(errors.New("db fail"))

	userToSave := &entity.User{
		Email:        "carol@example.com",
		PasswordHash: "password",
	}

	err := repo.Save(userToSave)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "db fail")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unfulfilled expectations: %v", err)
	}
}

package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/fajrinajiseno/mygolangapp/internal/entity"
	"github.com/fajrinajiseno/mygolangapp/internal/module/auth/repository/mock"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuth_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	password := "secret123"

	mockRepo := mock.NewMockUserRepository(ctrl)

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().
			Save(gomock.AssignableToTypeOf(&entity.User{})). // or gomock.Any()
			DoAndReturn(func(u *entity.User) error {
				assert.Equal(t, "alice@example.com", u.Email)
				err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
				assert.NoError(t, err)

				return nil
			})

		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		message, err := u.Register("alice@example.com", password, password)
		assert.NoError(t, err)
		assert.NotEmpty(t, message)
		assert.Equal(t, "Success Register", message)
	})

	t.Run("Invalid email or password", func(t *testing.T) {
		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		_, err := u.Register("", "", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid email or password")
	})

	t.Run("Invalid password", func(t *testing.T) {
		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		_, err := u.Register("alice@example.com", "asdsa", "dddd")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid password")
	})

	t.Run("Repo Error", func(t *testing.T) {
		mockRepo.EXPECT().
			Save(gomock.AssignableToTypeOf(&entity.User{})). // or gomock.Any()
			DoAndReturn(func(u *entity.User) error {
				assert.Equal(t, "alice@example.com", u.Email)
				err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
				assert.NoError(t, err)
				return errors.New("db fail")
			})

		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		_, err := u.Register("alice@example.com", password, password)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
	})
}

func TestAuth_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	password := "secret123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	user := &entity.User{
		ID:           1,
		Email:        "alice@example.com",
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	mockRepo := mock.NewMockUserRepository(ctrl)

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().
			GetUserByEmail("alice@example.com").
			Return(user, nil)

		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		tokenStr, gotUser, err := u.Login("alice@example.com", password)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenStr)
		assert.Equal(t, user, gotUser)

		parsed, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("unexpected alg")
			}
			return secret, nil
		})
		assert.NoError(t, err)
		assert.True(t, parsed.Valid)

		claims, ok := parsed.Claims.(jwt.MapClaims)
		assert.True(t, ok)
		assert.Equal(t, "1", claims["sub"])
		assert.Equal(t, "alice@example.com", claims["email"])
	})

	t.Run("Wrong Password", func(t *testing.T) {
		mockRepo.EXPECT().
			GetUserByEmail("alice@example.com").
			Return(user, nil)

		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		// wrong password
		_, _, err = u.Login("alice@example.com", "wrong-password")
		assert.Error(t, err)
		// wrapped error message contains "invalid credentials" according to usecase
		assert.Contains(t, err.Error(), "invalid credentials")
	})

	t.Run("Repo Error", func(t *testing.T) {
		mockRepo.EXPECT().
			GetUserByEmail("alice@example.com").
			Return(nil, errors.New("db fail"))

		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		_, _, err := u.Login("alice@example.com", "pw")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
	})

	t.Run("User Not Found EmptyID", func(t *testing.T) {
		mockRepo.EXPECT().
			GetUserByEmail("noone@example.com").
			Return(&entity.User{}, nil)

		secret := []byte("test-secret")
		u := NewAuthUsecase(mockRepo, secret, time.Hour)

		_, _, err := u.Login("noone@example.com", "pw")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

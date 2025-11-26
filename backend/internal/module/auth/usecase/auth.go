package usecase

import (
	"strconv"
	"time"

	"github.com/fajrinajiseno/mygolangapp/internal/entity"
	"github.com/fajrinajiseno/mygolangapp/internal/module/auth/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source auth.go -destination mock/auth_mock.go -package=mock
type AuthUsecase interface {
	Register(email string, password string, confirmPassword string) (string, error)
	Login(email string, password string) (string, *entity.User, error)
}

type Auth struct {
	repo      repository.UserRepository
	jwtSecret []byte
	ttl       time.Duration
}

func NewAuthUsecase(repo repository.UserRepository, jwtSecret []byte, ttl time.Duration) *Auth {
	return &Auth{repo: repo, jwtSecret: jwtSecret, ttl: ttl}
}

func (a *Auth) Register(email string, password string, confirmPassword string) (string, error) {
	if email == "" || password == "" || confirmPassword == "" {
		return "", entity.ErrorForbidden("invalid email or password")
	}
	if password != confirmPassword {
		return "", entity.ErrorForbidden("invalid password")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", entity.WrapError(err, entity.ErrorCodeInternal, "hash error")
	}

	user := &entity.User{
		Email:        email,
		PasswordHash: string(hashed),
	}

	if err := a.repo.Save(user); err != nil {
		return "", err
	}
	user.PasswordHash = ""
	return "Success Register", nil
}

// Login verifies email + password and returns a JWT.
func (a *Auth) Login(email string, password string) (string, *entity.User, error) {
	user, err := a.repo.GetUserByEmail(email)
	if err != nil {
		return "", nil, err
	}
	if user.ID == 0 {
		return "", nil, entity.ErrorNotFound("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, entity.WrapError(err, entity.ErrorCodeUnauthorized, "invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub":   strconv.Itoa(user.ID),
		"email": user.Email,
		"exp":   time.Now().Add(a.ttl).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", nil, entity.WrapError(err, entity.ErrorCodeUnauthorized, "invalid credentials")
	}
	return signed, user, nil
}

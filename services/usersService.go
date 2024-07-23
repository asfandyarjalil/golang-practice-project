package services

import (
	"encoding/base64"
	"net/http"

	"github.com/asfandyarjalil/golang-practice-project/models"
	"github.com/asfandyarjalil/golang-practice-project/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (us UsersService) Login(username string, password string) (string, *models.ResponseError) {
	id, responseErr := us.usersRepository.LoginUser(username, password)
	if responseErr != nil {
		return "", responseErr
	}
	if id == "" {
		return "", &models.ResponseError{
			Message: "Login failed",
			Status:  http.StatusUnauthorized,
		}
	}
	accessToken, responseErr := generateAccessToken(username)
	if responseErr != nil {
		return "", responseErr
	}

	us.usersRepository.SetAccessToken(accessToken, id)

	return accessToken, nil

}
func (us UsersService) Logout(accessToken string) *models.ResponseError {
	if accessToken == "" {
		return &models.ResponseError{
			Message: "Invalid access token",
			Status:  http.StatusBadRequest,
		}
	}

	return us.usersRepository.RemoveAccessToken(accessToken)
}

func generateAccessToken(username string) (string, *models.ResponseError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(username), bcrypt.DefaultCost)
	if err != nil {
		return "", &models.ResponseError{
			Message: "Failed to generate token",
			Status:  http.StatusInternalServerError,
		}
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}

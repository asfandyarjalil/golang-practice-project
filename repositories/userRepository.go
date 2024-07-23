package repositories

import (
	"database/sql"
	"net/http"

	"github.com/asfandyarjalil/golang-practice-project/models"
)

type UsersRepository struct {
	dbHandler *sql.DB
}

func NewUsersRepository(dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{dbHandler: dbHandler}
}

func (ur UsersRepository) LoginUser(username string, password string) (string, *models.ResponseError) {
	query := `
		SELECT id
		FROM users
		WHERE username = $1 and user_password = crypt($2, user_password)`

	rows, err := ur.dbHandler.Query(query, username, password)
	if err != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return "", &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if rows.Err() != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return id, nil
}
func (ur UsersRepository) SetAccessToken(accessToken string, id string) *models.ResponseError {
	query := `UPDATE users SET access_token = $1 WHERE id = $2`

	_, err := ur.dbHandler.Exec(query, accessToken, id)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (ur UsersRepository) RemoveAccessToken(accessToken string) *models.ResponseError {
	query := `UPDATE users SET access_token = '' WHERE access_token = $1`

	_, err := ur.dbHandler.Exec(query, accessToken)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

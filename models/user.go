package models

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"user_password"`
	Role        string `json:"user_role"`
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"user_password"`
}

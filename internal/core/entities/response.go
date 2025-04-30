package entities

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserRegisterResponse = UserLoginResponse

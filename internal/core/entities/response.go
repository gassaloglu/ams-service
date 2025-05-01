package entities

type LoginUserResponse struct {
	Token string `json:"token"`
}

type RegisterUserResponse = LoginUserResponse

type LoginEmployeeResponse = LoginUserResponse

type RegisterEmployeeResponse = RegisterUserResponse

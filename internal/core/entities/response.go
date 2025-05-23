package entities

type LoginUserResponse struct {
	Token string `json:"token"`
}

type RegisterUserResponse = LoginUserResponse

type LoginEmployeeResponse struct {
	Token      string `json:"token"`
	Permission string `json:"permission"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
}

type RegisterEmployeeResponse = RegisterUserResponse

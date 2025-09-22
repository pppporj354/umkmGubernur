package dto

type RegisterRequestDTO struct {
	FullName string `json:"fullName" binding:"required,min=2,max=120"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponseDTO struct {
	User   UserDTO   `json:"user"`
	Tokens TokensDTO `json:"tokens"`
}

type UserDTO struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type TokensDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

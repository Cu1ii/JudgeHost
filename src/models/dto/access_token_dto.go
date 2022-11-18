package dto

type AccessTokenDTO struct {
	AccessToken string
}

func NewAccessTokenDTO(token string) *AccessTokenDTO {
	return &AccessTokenDTO{token}
}

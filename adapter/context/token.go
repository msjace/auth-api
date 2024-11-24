package context

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AccessTokenRequest struct {
	AccessToken string `json:"access_token"`
}

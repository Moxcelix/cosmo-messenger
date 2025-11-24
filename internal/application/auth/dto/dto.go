package dto

type LoginData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshData struct {
	AccessToken string `json:"access_token"`
}

type ValidateData struct {
	UserID string `json:"user_id"`
}

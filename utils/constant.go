package utils

const (
	ErrInternalServer = "Internal server error"
	ErrUnauthorized   = "Unauthorized access"
	ErrNotFound       = "Resource not found"
)

type AceesTokenAndRefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

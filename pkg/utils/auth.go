package utils

type AuthHeaders struct {
	JWT    *string `json:"jwt,omitempty"`
	APIKey *string `json:"apiKey,omitempty"`
}

type ApiHeaders struct {
	ContentType   string  `json:"Content-Type"`
	Authorization *string `json:"Authorization,omitempty"`
	XAPIKey       *string `json:"X-API-KEY,omitempty"`
}

package config

type Config struct {
	Client ClientConfig
	Api    ApiConfig
}

type ClientConfig struct {
	Timeout int64
}

type ApiConfig struct {
	Network    string
	JWT        string
	APIKey     string
	SigningKey string
}

var globalConfig = &Config{
	Client: ClientConfig{
		Timeout: 10,
	},
}

func GetConfig() *Config {
	return globalConfig
}

// func NewConfig(apiKey string, network string) *Config {
// 	apiConfig := &ApiConfig{
//         Network:     "mainnet",
//         JWT:         "your_jwt_token",
//         APIKey:      "your_api_key",
//         SigningKey:  "your_signing_key",
//     }

// 	clientConfig := &ClientConfig{
// 		Timeout: 10,
// 	},

// 	var config = &Config{
// 		Client: *clientConfig,
// 		Api: *apiConfig,}

// 	return config
// }

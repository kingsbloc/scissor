package config

import (
	"os"

	"github.com/kingsbloc/scissor/docs"
)

var (
	HOST_URL      string
	CONTACT_EMAIL = "kingsleynwankwou@gmail.com"
	SCHEMES       []string
)

func InitConstants() {
	if os.Getenv("APP_ENV") == "production" {
		HOST_URL = "http://ec2-44-211-80-65.compute-1.amazonaws.com"
		SCHEMES = []string{"http"}
	} else {
		HOST_URL = os.Getenv("SERVER_URL")
		SCHEMES = []string{"http"}
	}
	docs.SwaggerInfo.Host = HOST_URL
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Title = "Scissor API"
	docs.SwaggerInfo.Schemes = SCHEMES
	docs.SwaggerInfo.BasePath = "/"
}

type contextKey int

const (
	AccountID contextKey = iota
	JWTAuthContext
)

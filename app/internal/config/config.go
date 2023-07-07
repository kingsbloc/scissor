package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Jwt    ConfJWT
	Server ConfServer
}

type ConfServer struct {
	Port         int           `env:"PORT,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}
type ConfJWT struct {
	Access_secret  string `env:"JWT_ACCESS_SECRET,required"`
	Refresh_secret string `env:"JWT_REFRESH_SECRET,required"`
	Issuer         string `env:"JWT_ISSUER,required"`
	Audience       string `env:"JWT_AUDIENCE,required"`
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}

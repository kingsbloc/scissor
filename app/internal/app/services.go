package app

import "github.com/kingsbloc/scissor/internal/services"

type MicroServices struct {
	UserService    services.UserService
	AuthService    services.AuthService
	JwtService     services.JwtService
	ShortenService services.ShortenService
	RedisService   services.RedisService
}

func NewMicroServices(userService services.UserService, authService services.AuthService, jwtService services.JwtService, shortenService services.ShortenService, redisService services.RedisService) *MicroServices {
	return &MicroServices{
		UserService:    userService,
		AuthService:    authService,
		JwtService:     jwtService,
		ShortenService: shortenService,
		RedisService:   redisService,
	}
}

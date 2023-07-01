package app

import "github.com/kingsbloc/scissor/internal/services"

type MicroServices struct {
	UserService services.UserService
}

func NewMicroServices(userService services.UserService) *MicroServices {
	return &MicroServices{
		UserService: userService,
	}
}

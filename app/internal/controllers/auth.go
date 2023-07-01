package controllers

type AuthController interface{}

type authController struct{}

func NewAuthController() AuthController {
	return &authController{}
}

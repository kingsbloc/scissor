package controllers

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/render"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/dto"
	"github.com/kingsbloc/scissor/internal/models"
	"github.com/kingsbloc/scissor/internal/utils"
)

type AuthController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	srv *app.MicroServices
}

func NewAuthController(srv *app.MicroServices) AuthController {
	return &authController{srv: srv}
}

// Create
// @Summary Create User.
// @Description Create User Account.
// @Tags Auth
// @Accept	json
// @Produce	json
// @Param requestBody body dto.AddUserDto true "Add User Dto"
// @Success 201 {object} utils.ApiResponse{data=bool}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Failure 422 {object} utils.ApiResponse{data=[]utils.ValidationError}
// @Router /api/v1/auth/signup [post]
func (con *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	var dto dto.AddUserDto
	if err := render.Bind(r, &dto); err != nil {
		render.Render(w, r, utils.ErrValidationRequest(err, "Validation Error"))
		return
	}

	var user models.User

	user.Name = dto.Name
	user.Email = dto.Email
	user.HashPassword(dto.Password)

	_, err1 := con.srv.UserService.NewUser(&dto)

	if err1 != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err1))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &utils.ApiResponse{
		Status:  http.StatusCreated,
		Message: "Account Created Successfully",
		Success: true,
	})
}

// Login
// @Summary Login User.
// @Description Login User Account.
// @Tags Auth
// @Accept	json
// @Produce	json
// @Param requestBody body dto.LoginDto true "Login Dto"
// @Success 201 {object} utils.ApiResponse{data=bool}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Failure 422 {object} utils.ApiResponse{data=[]utils.ValidationError}
// @Router /api/v1/auth/login [post]
func (con *authController) Signin(w http.ResponseWriter, r *http.Request) {
	var dto dto.LoginDto
	if err := render.Bind(r, &dto); err != nil {
		render.Render(w, r, utils.ErrValidationRequest(err, "Validation Error"))
		return
	}
	user, err1 := con.srv.AuthService.Signin(&dto)
	if err1 != nil {
		render.Render(w, r, &utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: err1.Error(),
			Success: false,
		})
		return
	}
	var jwt string
	var refreshJWT string
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		jwt1, err := con.srv.JwtService.GenerateJWT(user.Email, string(rune(user.ID)), false)
		if err != nil {
			log.Fatal(err)
		}
		jwt = jwt1
	}()
	go func() {
		defer wg.Done()
		refreshJWT1, err := con.srv.JwtService.GenerateJWT(user.Email, string(rune(user.ID)), true)
		if err != nil {
			log.Fatal(err)
		}
		refreshJWT = refreshJWT1
	}()
	wg.Wait()

	render.Render(w, r, &utils.ApiResponse{
		Success: true,
		Message: "Login Success",
		Status:  http.StatusCreated,
		Data: map[string]interface{}{
			"user":         user,
			"accessToken":  jwt,
			"refreshToken": refreshJWT,
		},
	})
}

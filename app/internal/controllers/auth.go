package controllers

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/dto"
	"github.com/kingsbloc/scissor/internal/models"
	"github.com/kingsbloc/scissor/internal/utils"
)

type AuthController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
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
// @Tags Student
// @Accept	json
// @Produce	json
// @Param requestBody body dto.AddUserDto true "Add User Dto"
// @Success 201 {object} utils.ApiResponse{data=bool}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Failure 422 {object} utils.ApiResponse{data=[]utils.ValidationError}
// @Router /students/create [post]
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

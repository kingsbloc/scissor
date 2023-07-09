package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/models"
	"github.com/kingsbloc/scissor/internal/utils"
)

type UserController interface {
	UrlHistory(w http.ResponseWriter, r *http.Request)
	DeleteShorten(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	srv *app.MicroServices
}

func NewUserController(srv *app.MicroServices) UserController {
	return &userController{srv: srv}
}

// History
// @Summary History of Shorten Url.
// @Description History of Shorten Url.
// @Tags User
// @Accept	json
// @Produce	json
// @Success 201 {object} utils.ApiResponse{data=[]dto.AddShortenDto}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Failure 422 {object} utils.ApiResponse{data=[]utils.ValidationError}
// @Router /api/v1/user/history [get]
func (con *userController) UrlHistory(w http.ResponseWriter, r *http.Request) {
	userId := con.srv.JwtService.GetJWTAuthContext(r).Get("id")

	result := con.srv.UserService.ShortenHistory(userId)

	var list []models.Shorten

	for _, v := range result {
		list = append(list, v.ToBase())
	}

	render.Render(w, r, &utils.ApiResponse{
		Status:  http.StatusOK,
		Message: "",
		Success: true,
		Data: map[string]interface{}{
			"history": list,
		},
	})
}

// Delete Shorten
// @Summary Delete Shorten Url.
// @Description Delete Shorten Url.
// @Tags User
// @Accept	json
// @Produce	json
// @Success 201 {object} utils.ApiResponse{data=bool}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Failure 422 {object} utils.ApiResponse{data=[]utils.ValidationError}
// @Router /api/v1/user/history/:id [post]
func (con *userController) DeleteShorten(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	userId := con.srv.JwtService.GetJWTAuthContext(r).Get("id")

	result, ok := con.srv.UserService.DeleteShorten(idParam, userId)
	if !ok {
		render.Render(w, r, &utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to Delete",
			Success: false,
		})
		return
	}
	go func() {
		con.srv.RedisService.Delete(result.CustomShort)
	}()
	render.Render(w, r, &utils.ApiResponse{
		Status:  http.StatusOK,
		Message: "Deleted Successfully",
		Success: true,
	})
}

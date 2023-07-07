package controllers

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/dto"
	"github.com/kingsbloc/scissor/internal/utils"
)

type ShortenController interface {
	ShortenUrl(w http.ResponseWriter, r *http.Request)
}

type shortenController struct {
	srv *app.MicroServices
}

func NewShortenController(srv *app.MicroServices) ShortenController {
	return &shortenController{srv: srv}
}

// Shorten
// @Summary Shorten Url.
// @Description New Shorten Url.
// @Tags Shorten
// @Accept	json
// @Produce	json
// @Param requestBody body dto.AddShortenDto true "Add ShortenDto"
// @Success 201 {object} utils.ApiResponse{data=bool}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Failure 422 {object} utils.ApiResponse{data=[]utils.ValidationError}
// @Router /api/v1/shorten [post]
func (con *shortenController) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var userId string
	_id := con.srv.JwtService.GetJWTAuthContext(r).Get("id")
	if len(_id) > 0 {
		userId = _id
		log.Println(userId)
	}
	var dto dto.AddShortenDto
	if err := render.Bind(r, &dto); err != nil {
		render.Render(w, r, utils.ErrValidationRequest(err, "Validation Error"))
		return
	}

}

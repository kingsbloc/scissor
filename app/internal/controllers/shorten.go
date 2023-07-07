package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/config"
	"github.com/kingsbloc/scissor/internal/dto"
	"github.com/kingsbloc/scissor/internal/utils"
)

type ShortenController interface {
	ShortenUrl(w http.ResponseWriter, r *http.Request)
	ResolveUrl(w http.ResponseWriter, r *http.Request)
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
// @Success 201 {object} utils.ApiResponse{data=dto.AddShortenDto}
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
	var body dto.AddShortenDto
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, utils.ErrValidationRequest(err, "Validation Error"))
		return
	}

	if !con.srv.ShortenService.CheckDomainError(body.Url) {
		render.Render(w, r, &utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Url not allowed",
			Success: false,
		})
		return
	}

	body.Url = con.srv.ShortenService.EnforceHTTP(body.Url)

	var id string
	if body.CustomUrl == "" {
		id = con.srv.ShortenService.GenerateNewShort()
	} else {
		id = body.CustomUrl
	}

	existCustomShort := con.srv.RedisService.Get(id)
	if existCustomShort != nil && body.CustomUrl != "" {
		render.Render(w, r, &utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Custom Short already in use",
			Success: false,
		})
		return
	}

	if body.Exp == 0 {
		body.Exp = 24
	}

	err1 := con.srv.RedisService.Set(id, body.Url, time.Duration(time.Duration(body.Exp).Hours())).Err()

	if err1 != nil {
		render.Render(w, r, &utils.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unable to Connect to Server",
			Success: false,
		})
		return
	}

	if userId != "" {
		// TODO save to user with gorotine
		go log.Println("")
	}

	render.Render(w, r, &utils.ApiResponse{
		Status:  http.StatusCreated,
		Message: "Url Shorten",
		Success: true,
		Data: map[string]interface{}{
			"url":          body.Url,
			"custom_short": config.New().Server.ServerUrl + "/" + id,
			"expiry":       body.Exp,
		},
	})

}

// Resolve
// @Summary Resolve Url.
// @Description Redirect or Resolve Url.
// @Tags Shorten
// @Accept	json
// @Produce	json
// @Param id path int true "Course ID"
// @Success 201 {object} utils.ApiResponse{data=dto.AddShortenDto}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Failure 422 {object} utils.ApiResponse{data=[]utils.ValidationError}
// @Router /api/v1/shorten/{id} [get]
func (con *shortenController) ResolveUrl(w http.ResponseWriter, r *http.Request) {}

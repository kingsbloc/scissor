package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/controllers"
)

func ShortenRoutes(r chi.Router, srv *app.MicroServices) {
	con := controllers.NewShortenController(srv)
	r.With(srv.JwtService.JWTAuth).Route("/shorten", func(r chi.Router) {
		r.Post("/", con.ShortenUrl)
	})
}

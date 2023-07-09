package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/controllers"
)

func UserRoutes(r chi.Router, srv *app.MicroServices) {
	con := controllers.NewUserController(srv)
	r.With(srv.JwtService.JWTAuth).Route("/user", func(r chi.Router) {
		r.Get("/history", con.UrlHistory)
		r.Post("/history/{id}", con.DeleteShorten)
	})
}

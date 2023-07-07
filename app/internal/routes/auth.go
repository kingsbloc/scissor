package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/controllers"
)

func AuthRoutes(r chi.Router, srv *app.MicroServices) {
	con := controllers.NewAuthController(srv)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", con.SignUp)
		r.Post("/signin", con.Signin)
	})
}

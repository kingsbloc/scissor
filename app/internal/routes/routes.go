package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/controllers"
)

func RegisterRoutes(r *chi.Mux, srv *app.MicroServices) {
	con := controllers.NewShortenController(srv)
	r.Route("/api/v1", func(r chi.Router) {
		AuthRoutes(r, srv)
		ShortenRoutes(r, srv)
		UserRoutes(r, srv)
	})
	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", con.ResolveUrl)
	})
}

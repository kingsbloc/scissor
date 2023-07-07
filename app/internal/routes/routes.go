package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kingsbloc/scissor/internal/app"
)

func RegisterRoutes(r *chi.Mux, srv *app.MicroServices) {
	r.Route("/api/v1", func(r chi.Router) {
		AuthRoutes(r, srv)
		ShortenRoutes(r, srv)
	})
}

package routes

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/kingsbloc/scissor/docs"
	"github.com/kingsbloc/scissor/internal/config"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterSwaggerRoutes(r *chi.Mux) {
	c := config.New()
	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL(c.Server.ServerUrl+"/docs/doc.json"),
	))
	// r.Route("/docs", func(r chi.Router) {
	// 	r.Get("/*", swaggerfiles.Handler.ServeHTTP)
	// })
}

package routes

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/kingsbloc/scissor/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterSwaggerRoutes(r *chi.Mux) {
	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5000/docs/doc.json"),
	))
	// r.Route("/docs", func(r chi.Router) {
	// 	r.Get("/*", swaggerfiles.Handler.ServeHTTP)
	// })
}

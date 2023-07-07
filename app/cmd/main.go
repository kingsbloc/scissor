package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/kingsbloc/scissor/internal/app"
	"github.com/kingsbloc/scissor/internal/config"
	"github.com/kingsbloc/scissor/internal/models"
	"github.com/kingsbloc/scissor/internal/repositories"
	"github.com/kingsbloc/scissor/internal/routes"
	"github.com/kingsbloc/scissor/internal/services"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitConstants()
}

// @description	This is scissor server.
// @BasePath	/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// schemes: [http, https]

func main() {
	c := config.New()

	// Init the main Router
	r := chi.NewRouter()

	// Add Middlewares
	r.Use(
		middleware.AllowContentType("application/json"),
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	r.Use(middleware.CleanPath)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	// Connect Redis
	rdb := config.ConnectRedis()

	// Connect to DB
	dbConn, dbErr := repositories.InitDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	log.Println("==== DB connected")

	// AutoMigrate Models
	repositories.AutoMigrate([]interface{}{
		&models.User{}, &models.Shorten{},
	}, dbConn)

	// Create New DAO
	dao := repositories.NewDAO(dbConn)

	// create services
	userService := services.NewUserService(dao)
	authService := services.NewAuthService(dao)
	jwtService := services.NewJwtService()
	shortenService := services.NewShortenService(dao)
	redisService := services.NewRedisService(rdb)

	// initialize Microservices
	srv := app.NewMicroServices(userService, authService, jwtService, shortenService, redisService)

	// Register Routes
	routes.RegisterSwaggerRoutes(r)
	routes.RegisterRoutes(r, srv)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}

}

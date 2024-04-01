package main

import (
	"gocaptcha/internal/server/handlers"
	"gocaptcha/internal/server/handlers/middleware"
	"gocaptcha/internal/server/services"
	"gocaptcha/pkg"

	"github.com/gofiber/fiber/v2"
)

func main() {

	ttl := pkg.New(1000, 0)

	rateLimit := middleware.RateLimit(ttl)

	serv := services.REST{}

	handler := handlers.NewHandlerRest(&serv)

	app := fiber.New()

	app.Use(rateLimit)

	app.Get("/:id", handler.GetRandomCap)

	app.Listen(":3000")
}

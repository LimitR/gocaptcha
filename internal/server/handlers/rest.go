package handlers

import (
	"bytes"
	"gocaptcha/internal/server/services"
	"image/png"

	"github.com/gofiber/fiber/v2"
)

type REST struct {
	service *services.REST
}

func NewHandlerRest(service *services.REST) *REST {
	return &REST{
		service: service,
	}
}

func (r *REST) GetRandomCap(c *fiber.Ctx) error {
	id := c.Query("id")

	image := r.service.GetCaptcha(id)

	body := &bytes.Buffer{}

	png.Encode(body, image)

	c.Set("content-type", "image/png")

	return c.Send(body.Bytes())
}

package services

import (
	"gocaptcha/internal/core"
	"image"
	"image/color"
)

type REST struct{}

func (r *REST) GetCaptcha(id string) *image.RGBA {
	rr := core.Randomizer{}

	ph := core.GeneratePhoto{
		Color: color.RGBA{200, 0, 0, 255},
		Size: core.PhotoSize{
			XPx: 200, YPx: 50,
		},
	}

	image := ph.GeneratePhotoWithLabel(rr.GetString(5))

	return image
}

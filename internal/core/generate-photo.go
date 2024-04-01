package core

import (
	"image"
	"image/color"
	"os"

	"github.com/goki/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type PhotoSize struct {
	XPx int
	YPx int
}

type GeneratePhoto struct {
	Color color.RGBA
	Size  PhotoSize
}

func (g *GeneratePhoto) GeneratePhotoWithLabel(label string) *image.RGBA {
	image := image.NewRGBA(image.Rect(0, 0, g.Size.XPx, g.Size.YPx))

	g.addLabel(image, image.Rect.Dx()/2-len(label)-70, image.Rect.Dy()/2+10, label)

	return image
}

func (g *GeneratePhoto) addLabel(img *image.RGBA, x, y int, label string) {

	fontFile, _ := os.ReadFile("./default.ttf")
	ttf, _ := truetype.Parse(fontFile)

	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(g.Color),
		Face: truetype.NewFace(ttf, &truetype.Options{
			Size: 40,
		}),
		Dot: point,
	}
	d.DrawString(label)
}

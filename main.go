package main

import (
	"image"
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func setImageToimgMatrix(img *image.RGBA, m [][]uint8) {
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(m[x][y]),
				G: uint8(m[x][y]),
				B: uint8(m[x][y]),
				A: 255,
			})
		}
	}
}

func initWindow() () {
	cfg := pixelgl.WindowConfig{
		Title:  "Render Window",
		Bounds: pixel.R(0, 0, float64(W), float64(H)),
		VSync:  true,
	}
	window, _ = pixelgl.NewWindow(cfg)
	return
}

func renderWindow() {

	// update window to image
	img := image.NewRGBA(image.Rect(0, 0, W, H))

	for {

		// Update the image
		setImageToimgMatrix(img, imgMatrix)

		// Update the picture data
		pic := pixel.PictureDataFromImage(img)

		// Render the image
		sprite := pixel.NewSprite(pic, pic.Bounds())
		sprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()))
		window.Update()

		// close the window if escape is pressed
		if window.JustPressed(pixelgl.KeyEscape) {
			break
		}

		// Wait
		time.Sleep(time.Duration(time.Duration(1/fps).Seconds()))

	}

}

func main() {
	///////////////////////////////////////
	// call our functions
	initParams()
	initWindow()
	pixelgl.Run(renderWindow)
	///////////////////////////////////////
}

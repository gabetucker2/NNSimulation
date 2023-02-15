package main

import (
	"image"
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func getIMColCoords(x, y int) (col *Color) {
	col = NewColor(imgMatrix[0][x][y], imgMatrix[1][x][y], imgMatrix[2][x][y])
	return
}

// func getIMCol(pix *Pixel) *Color {
// 	return getIMColCoords(pix.pos.X, pix.pos.Y)
// }

func setIMCol(pix *Pixel) {
	imgMatrix[0][pix.pos.X][pix.pos.Y] = pix.col.R
	imgMatrix[1][pix.pos.X][pix.pos.Y] = pix.col.G
	imgMatrix[2][pix.pos.X][pix.pos.Y] = pix.col.B
}

func updatePixelCol(pix *Pixel, col *Color, updateIM ...any) {

	pix.col.R = col.R
	pix.col.G = col.G
	pix.col.B = col.B

	if len(updateIM) == 0 || updateIM[0] == true {
		setIMCol(pix)
	}

}

func setImageToImgMatrix(img *image.RGBA, imgMatrix [][][]uint8) {
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
  			img.Set(x, y, color.RGBA {
				R: uint8(imgMatrix[0][x][y]),
				G: uint8(imgMatrix[1][x][y]),
				B: uint8(imgMatrix[2][x][y]),
				A: 255,
			})
		}
	}
}

func initWindow() () {
	cfg := pixelgl.WindowConfig{
		Title:  "Window Renderer",
		Bounds: pixel.R(0, 0, float64(windowSize.X), float64(windowSize.Y)),
		VSync:  true,
	}
	window, _ = pixelgl.NewWindow(cfg)
	return
}

func keyPresses() {
	if window.JustPressed(pixelgl.KeyEnter) {
		runCEModel()
	}
}

func renderWindow() {

	// update window to image
	img := image.NewRGBA(image.Rect(0, 0, windowSize.X, windowSize.Y))

	// do until escape is pressed
	for !window.Closed() {

		// Update the image
		setImageToImgMatrix(img, imgMatrix)

		// Update the picture data
		pic := pixel.PictureDataFromImage(img)

		// Render the image
		sprite := pixel.NewSprite(pic, pic.Bounds())
		sprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()))
		window.Update()

		// Go through rest button press checks
		keyPresses()

		// Wait
		time.Sleep(time.Duration(time.Duration(1/fps).Seconds()))

	}

}

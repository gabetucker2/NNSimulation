package main

import (
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func getPixel(x, y int) (pix *Pixel) {
	col := NewColor(imgMatrix[0][x][y], imgMatrix[1][x][y], imgMatrix[2][x][y])
	pix = NewPixel(x, y, col)
	return
}

func setIMColCoords(x, y int, col *Color) {
	imgMatrix[0][x][y] = col.r
	imgMatrix[1][x][y] = col.g
	imgMatrix[2][x][y] = col.b
}

func setIMCol(pix *Pixel) {
	setIMColCoords(pix.pos.x, pix.pos.y, pix.col)
}

func fillIM(col *Color) {
	for x := 0; x < windowSize.x; x++ {
		for y := 0; y < windowSize.y; y++ {
			setIMColCoords(x, y, col)
		}
	}
}

func updatePixelCol(pix *Pixel, col *Color, updateIM ...any) {
	
	pix.col.r = col.r
	pix.col.g = col.g
	pix.col.b = col.b
	
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
		Bounds: pixel.R(0, 0, float64(windowSize.x), float64(windowSize.y)),
		VSync:  true,
	}
	window, _ = pixelgl.NewWindow(cfg)
	return
}

func keyPresses() {
	if window.JustPressed(pixelgl.KeyRight) {
		modelUpdateRight()
	}
	if window.JustPressed(pixelgl.KeyLeft) {
		modelUpdateLeft()
	}
	if window.JustPressed(pixelgl.KeyUp) {
		modelUpdateUp()
	}
	if window.JustPressed(pixelgl.KeyDown) {
		modelUpdateDown()
	}
}

func renderWindow() {

	// update window to image
	img := image.NewRGBA(image.Rect(0, 0, windowSize.x, windowSize.y))

	// do until X is pressed
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

	}

}

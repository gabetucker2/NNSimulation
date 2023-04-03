package main

import (
	"image"

	"github.com/faiface/pixel"
)

func renderRoutine() {

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

func physicsRoutine() {

	// do until X is pressed
	for !window.Closed() {

		for i := 0; i < len(ce.effectors); i++ {
			if i < ce.anchorIdx {
				*ce.effectors[i].pos = tensionTransform(ce.effectors[i].pos, ce.effectors[i+1].pos)
			} else if i > ce.anchorIdx {
				*ce.effectors[i].pos = tensionTransform(ce.effectors[i].pos, ce.effectors[i-1].pos)
			} // else if anchor, do nothing
		}

	}

}
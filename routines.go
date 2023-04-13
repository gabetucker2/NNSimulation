package main

import (
	"time"

	"github.com/faiface/pixel"
)

func initRoutines() {

	// coroutines
	renderTicker := time.NewTicker(time.Duration(frameΔT * float64(timePrecision)))
	physicsTicker := time.NewTicker(time.Duration(physicsΔT * float64(timePrecision)))
	go func() {
		for {
			select {
			case <-renderTicker.C:
				renderRoutine()
			case <-physicsTicker.C:
				physicsRoutine()
			}
		}
	}()

	// main routine (coroutines run concurrently with this loop)
	for !window.Closed() {
		keyPresses()
		window.Update()
	}

	// on closed
	renderTicker.Stop()
	physicsTicker.Stop()

}

func renderRoutine() {

	// Update the image
	setImageToPixelMatrix(img, pixelMatrix)

	// Update the picture data
	pic := pixel.PictureDataFromImage(img)

	// Render the image
	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()))

}

func physicsRoutine() {

	for i := 0; i < len(ce.effectors); i++ {
		if i < ce.anchorIdx {
			*ce.effectors[i].pos = tensionTransform(ce.effectors[i].pos, ce.effectors[i+1].pos)
		} else if i > ce.anchorIdx {
			*ce.effectors[i].pos = tensionTransform(ce.effectors[i].pos, ce.effectors[i-1].pos)
		} // else if anchor, do nothing
	}

}

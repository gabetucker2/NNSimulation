package main

import (
	"github.com/faiface/pixel/pixelgl"
)

var windowSize *Vector2
var imgMatrix [][][]uint8
var fps int
var window *pixelgl.Window
var emptyCol, slimeCol *Color

func initParams() {

	///////////////////////////////////////
	// CHANGE THESE

	// set up window width/height
	windowSize = NewVector2(500, 500)

	// set the fps
	fps = 60

	// set our colors
	emptyCol = NewColor(255, 255, 255)
	slimeCol = NewColor(0, 200, 200)

	///////////////////////////////////////
	// DON'T CHANGE THESE

	// set up imgMatrix
	imgMatrix = make([][][]uint8, 3)
	for i := 0; i < 3; i++ {
		imgMatrix[i] = make([][]uint8, windowSize.Y)
		for x := 0; x < windowSize.X; x++ {
			imgMatrix[i][x] = make([]uint8, windowSize.Y)
		}
	}

	///////////////////////////////////////

}

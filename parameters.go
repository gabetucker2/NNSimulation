package main

import (
	"math"

	"github.com/faiface/pixel/pixelgl"
)

var windowSize *Vector2
var imgMatrix [][][]uint8
var fps int
var window *pixelgl.Window
var nullCol, emptyCol, slimeCol *Color
var ce *CaenorhabditisElegans
var modelCall, modelUpdateLeft, modelUpdateRight, modelUpdateUp, modelUpdateDown func()
var inf float64

func initParams() {

	///////////////////////////////////////
	// CHANGE THESE

	// set up window width/height
	windowSize = NewVector2(500, 500)

	// set the fps
	fps = 60

	// set our colors
	emptyCol = NewColor(116, 116, 116)
	slimeCol = NewColor(59, 59, 59)

	// create ce
	ce = NewCE([]*Joint {
		NewJoint(100, 100, 50, true, []int {1}),  // 0
		NewJoint(150, 120, 20, false, []int {0}), // 1
	})

	// define which model we would like to run
	modelCall = ceModelCall
	modelUpdateLeft = ceModelUpdateLeft
	modelUpdateRight = ceModelUpdateRight
	modelUpdateUp = ceModelUpdateUp
	modelUpdateDown = ceModelUpdateDown

	///////////////////////////////////////
	// DON'T CHANGE THESE

	// set nullCol to some arbitrary color
	nullCol = NewColor(0, 0, 0)

	// set quasi-infinity constant
	inf = math.MaxFloat64

	// set up imgMatrix
	imgMatrix = make([][][]uint8, 3)
	for i := 0; i < 3; i++ {
		imgMatrix[i] = make([][]uint8, windowSize.Y)
		for x := 0; x < windowSize.X; x++ {
			imgMatrix[i][x] = make([]uint8, windowSize.Y)
		}
	}
	// initialize model
	modelCall()

	///////////////////////////////////////

}

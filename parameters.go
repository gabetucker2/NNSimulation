package main

import (
	"image"
	"math"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/gabetucker2/gogenerics"
)

var window *pixelgl.Window
var windowSize *Vector2Int
var pixelMatrix [][][]uint8
var nullCol, emptyCol, cytoCol, connectionCol *Color
var cytoDensity, connectionDensity float64
var ce *CaenorhabditisElegans
var modelCall, modelUpdateLeft, modelUpdateRight, modelUpdateUp, modelUpdateDown func()
var inf float64
var tensLowerThreshold, tensVertAsymptote, tensUpperThresholdOffset, tensSmoothness, tensTransFactor float64
var delta float64
var fps int
var frameΔT, physicsΔT float64 // in seconds
var timePrecision time.Duration
var img *image.RGBA

func initParams() {

	///////////////////////////////////////
	// CHANGE THESE

	// set up window width/height
	windowSize = NewVector2Int(500, 300)

	// set our colors
	emptyCol = NewColor(246, 240, 213)
	cytoCol = NewColor(132, 115, 34)
	connectionCol = NewColor(120, 105, 30)

	// set our densities
	cytoDensity = 0.9
	connectionDensity = 0.8

	// create ce
	ce = NewCE([]*Effector{
		NewEffector(250, 155, 25, true, []int{}),
		NewEffector(200, 170, 23, false, []int{0}),
		NewEffector(150, 150, 20, false, []int{1}),
		NewEffector(120, 140, 18, false, []int{2}),
		NewEffector(100, 130, 16, false, []int{3}),
		NewEffector(80, 120, 14, false, []int{4}),
		NewEffector(60, 110, 12, false, []int{5}),
		NewEffector(40, 100, 10, false, []int{6}),
	})

	// define which model we would like to run
	modelCall = ceModelCall
	modelUpdateLeft = ceModelUpdateLeft
	modelUpdateRight = ceModelUpdateRight
	modelUpdateUp = ceModelUpdateUp
	modelUpdateDown = ceModelUpdateDown

	// initialize tension equation parameters
	tensLowerThreshold = 9.2
	tensVertAsymptote = 39.3
	tensUpperThresholdOffset = 1.84
	tensSmoothness = 57.5
	tensTransFactor = 0.0005

	// initialize delta value
	delta = 0.000001

	// initialize time settings
	timePrecision = time.Microsecond // our time units update with temporal accuracy down to this unit
	fps = 10
	physicsΔT = 0.01 // seconds between each physics update

	///////////////////////////////////////
	// DON'T CHANGE THESE

	// Initialize window image
	img = image.NewRGBA(image.Rect(0, 0, windowSize.x, windowSize.y))

	// set frame update period
	frameΔT = 1.0 / float64(fps)

	// set nullCol to some arbitrary color
	nullCol = NewColor(0, 0, 0)
	gogenerics.RemoveUnusedError(nullCol)

	// set quasi-infinity constant
	inf = math.MaxFloat64

	// set up pixelMatrix
	pixelMatrix = make([][][]uint8, 3)
	for i := 0; i < 3; i++ {
		pixelMatrix[i] = make([][]uint8, windowSize.x)
		for x := 0; x < windowSize.x; x++ {
			pixelMatrix[i][x] = make([]uint8, windowSize.y)
		}
	}

	// initialize model
	modelCall()

	///////////////////////////////////////

}

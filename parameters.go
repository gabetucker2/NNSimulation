package main

import (
	"math"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

var windowSize *Vector2Int
var imgMatrix [][][]uint8
var window *pixelgl.Window
var nullCol, emptyCol, slimeCol *Color
var ce *CaenorhabditisElegans
var modelCall, modelUpdateLeft, modelUpdateRight, modelUpdateUp, modelUpdateDown func()
var inf float64
var tensLowerThreshold, tensVertAsymptote, tensUpperThresholdOffset, tensSmoothness, tensTransFactor float64
var delta float64
var fps int;
var deltaT float64;

func initParams() {

	///////////////////////////////////////
	// CHANGE THESE

	// set up window width/height
	windowSize = NewVector2Int(500, 300)

	// set our colors
	emptyCol = NewColor(116, 116, 116)
	slimeCol = NewColor(59, 59, 59)

	// create ce
	ce = NewCE([]*Effector {
		NewEffector(100, 150, 20, false, []int {1}), // 0
		NewEffector(200, 170, 30, false, []int {2}), // 1
		NewEffector(250, 155, 25, true , []int { }), // 2
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
	tensTransFactor = 120
	
	// initialize delta value
	delta = 0.000001

	// initialize time settings
	fps = 60
	deltaT = 0.01; // 10ms between each physics update

	///////////////////////////////////////
	// DON'T CHANGE THESE

	// set nullCol to some arbitrary color
	nullCol = NewColor(0, 0, 0)

	// set quasi-infinity constant
	inf = math.MaxFloat64

	// set up imgMatrix
	imgMatrix = make([][][]uint8, 3)
	for i := 0; i < 3; i++ {
		imgMatrix[i] = make([][]uint8, windowSize.x)
		for x := 0; x < windowSize.x; x++ {
			imgMatrix[i][x] = make([]uint8, windowSize.y)
		}
	}

	// initialize model
	modelCall()

	// initialize routines
	renderTicker := time.NewTicker(time.Duration(1/float64(fps)) * time.Second)
	physicsTicker := time.NewTicker(time.Duration(deltaT) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
		select {
			case <- renderTicker.C:
				renderRoutine()
			case <- physicsTicker.C:
				physicsRoutine()
			case <- quit:
				renderTicker.Stop()
				physicsTicker.Stop()
				return
			}
		}
	}()

	///////////////////////////////////////

}

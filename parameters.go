package main

import (
	"github.com/faiface/pixel/pixelgl"
)

var W, H int
var imgMatrix [][]uint8
var fps int
var window *pixelgl.Window

type CElegans struct {
	
}

func initParams() {

	///////////////////////////////////////
	// CHANGE THESE

	// set up window width/height
	W = 500
	H = 500

	// set the fps
	fps = 60

	///////////////////////////////////////
	// DON'T CHANGE THESE

	// set up imgMatrix
	imgMatrix = make([][]uint8, W)
	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			imgMatrix[x] = make([]uint8, H)
		}
	}

	///////////////////////////////////////

}

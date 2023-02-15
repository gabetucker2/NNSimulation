package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	// call our functions
	initParams()
	initWindow()
	pixelgl.Run(run)
}

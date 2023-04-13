package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func run() {

	// call our functions
	initParams()
	initWindow()
	initRoutines()

}

func main() {
	pixelgl.Run(run)
}

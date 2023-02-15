package main

import (
	"math/rand"
)

// create Caenorhabditis Elegans
func ceModelCall() {

	// get previous IM save before refresh
	// prevIM := imgMatrix

	// fill to empty
	fillIM(emptyCol)

	// update ce's physical representation
	for x := 0; x < windowSize.X; x++ {
		for y := 0; y < windowSize.Y; y++ {
			// prevColor := getColCoords(x, y, prevIM)
			pixel := NewPixel(x, y, nullCol)

			// check whether to render slime
			makeSlime := false
			for _, joint := range ce.joints {
				if SqrDistance(pixel.pos, joint.pos) - joint.sqrSize <= 0 {
					// distance to center of joint, from 0, inner-most, to 1, outer-most (non-linear)
					distance := SqrDistance(pixel.pos, joint.pos) / joint.sqrSize
					chanceSlime := rand.Float32()
					if distance < chanceSlime {
						makeSlime = true
					} else {
						makeSlime = false
					}
					break
				}
			}

			// set color
			newCol := nullCol
			if makeSlime {
				newCol = slimeCol
			}

			// update the image
			if newCol != nullCol {
				updatePixelCol(pixel, newCol)
			}
			
		}
	}

}

func ceModelUpdateLeft() {
	ce.joints[0].pos.X -= 10
	ceModelCall()
}

func ceModelUpdateRight() {
	ce.joints[0].pos.X += 10
	ceModelCall()
}

func ceModelUpdateUp() {
	ce.joints[0].pos.Y -= 10
	ceModelCall()
}

func ceModelUpdateDown() {
	ce.joints[0].pos.Y += 10
	ceModelCall()
}

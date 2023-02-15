package main

// create Caenorhabditis Elegans
func runCEModel() {

	// create ce struct
	ce := new(CaenorhabditisElegans)
	ce.joints = []*Joint {
		NewJoint(100, 100, 50, true, []int {1}), // 0
		NewJoint(150, 120, 20, false, []int {0}), // 1
	}

	// update ce's physical representation
	for x := 0; x < windowSize.X; x++ {
		for y := 0; y < windowSize.Y; y++ {
			pixel := NewPixel(x, y, getIMColCoords(x, y))

			// check whether to render slime
			inRangeOfJoint := false
			for _, joint := range ce.joints {
				if SqrDistance(pixel.pos, joint.pos) < joint.size {
					inRangeOfJoint = true
					break
				}
			}

			// set color
			var newCol *Color
			if inRangeOfJoint {
				newCol = slimeCol
			} else {
				newCol = emptyCol
			}

			// update the image
			updatePixelCol(pixel, newCol)
			
		}
	}

	// set ce's anchor to its anchor joint
	for _, j := range ce.joints {
		if j.anchor {
			ce.anchor = j
		}
	}

}

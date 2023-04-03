package main

import (
	"math/rand"
)

func (ce *CaenorhabditisElegans) getAnchorEffector() (e *Effector) {
	for _, effector := range ce.effectors {
		if effector.anchor {
			e = effector
			break
		}
	}
	return
}

// render slime pixel if meets conditions
func intraEffectorSlime(circle *Effector, pixel *Pixel) {

	// initialize test condition
	makeSlime := true

	// is it outside the circle
	if sqrMagnitude(pixel.pos, circle.pos) > circle.sqrRadius {
		makeSlime = false
	}

	// random chance not to render
	// chanceNotRender := 0.0
	// if rand.Float64() < chanceNotRender {
	// 	makeSlime = false
	// }

	// render pixel if passed
	if makeSlime {
		updatePixelCol(pixel, slimeCol)
	}

}

// render slime pixel if meets conditions
func interEffectorSlime(connection *Polygon, pixel *Pixel) {

	// initialize test condition
	makeSlime := true

 	// is it inside the connection polygon
	if !inPolygon(connection, NewVector2(pixel.pos.x, pixel.pos.y)) {
		makeSlime = false
	}

	// random chance not to render
	chanceNotRender := 0.1
	if rand.Float64() < chanceNotRender {
		makeSlime = false
	}

	// render pixel if passed
	if makeSlime {
		updatePixelCol(pixel, slimeCol)
	}

}

// create Caenorhabditis Elegans
func ceModelCall() {

	// get previous IM save before refresh
	// prevIM := imgMatrix

	// fill to empty
	fillIM(emptyCol)
			
	// check whether to render slime between effectors
	for _, effectorA := range ce.effectors {
		for _, idx := range effectorA.connectionIdxs {
			effectorB := ce.effectors[idx]

			// define our circle properties
			A := NewVector2(effectorA.pos.x, effectorA.pos.y)
			B := NewVector2(effectorB.pos.x, effectorB.pos.y)
			r_A := effectorA.radius
			r_B := effectorB.radius

			// define the slope of our intersect function between A and B
			m := m(A, B)

			// define the inverse slope of m
			M := M(m)

			// define the intersects between our inverse function and their corresponding circles
			A_1 := N_n(A, A, B, r_A, M, false)
			B_1 := N_n(B, A, B, r_B, M, false)
			A_2 := N_n(A, A, B, r_A, M, true)
			B_2 := N_n(B, A, B, r_B, M, true)


			// define the connection zone as a polygon
			connection := MakePolygon([]*Vector2 {A_1, B_1, B_2, A_2})

			// update ce's circle connection physical representations
			for x := 0; x < windowSize.x; x++ {
				for y := 0; y < windowSize.y; y++ {
					interEffectorSlime(connection, getPixel(x, y))
				}
			}
			
		}

		// update ce's circle physical representations
		for x := 0; x < windowSize.x; x++ {
			for y := 0; y < windowSize.y; y++ {
				intraEffectorSlime(effectorA, getPixel(x, y))
			}
		}
	}

}

func ceModelUpdateLeft() {
	ce.getAnchorEffector().pos.x -= 10
	ceModelCall()
}

func ceModelUpdateRight() {
	ce.getAnchorEffector().pos.x += 10
	ceModelCall()
}

func ceModelUpdateUp() {
	ce.getAnchorEffector().pos.y -= 10
	ceModelCall()
}

func ceModelUpdateDown() {
	ce.getAnchorEffector().pos.y += 10
	ceModelCall()
}

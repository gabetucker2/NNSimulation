package main

import (
	"math"
)

// render slime pixel if meets conditions
func intraJointSlime(joint *Joint, pixel *Pixel) {

	// initialize test case
	makeSlime := true

	// is it outside the circle
	if SqrDistance(pixel.pos, joint.pos) > math.Pow(joint.radius, 2) {
		makeSlime = false
	}

	// render pixel if passed
	if makeSlime {
		updatePixelCol(pixel, slimeCol)
	}

}

// render slime pixel if meets conditions
func interJointSlime(jointA, jointB *Joint, pixel *Pixel) {

	// initialize test case
	makeSlime := true

	// according with: https://www.desmos.com/calculator/cy12nq52sx
	// define our circle properties
	A := jointA.pos
	B := jointB.pos
	x := float64(pixel.pos.X)
	y := float64(pixel.pos.Y)
	r_A := jointA.radius
	r_B := jointB.radius

	// define the slope of our intersect function f(x) between the circle centers
	m_f := m(A, B)

	// define the inverse functions of f(x) F_A(X) and F_A(X) intersecting with A or B, respectively
	M := M(m_f)
	F_Ax := F_Nx(x, M, A)
	F_Bx := F_Nx(x, M, B)

	// define the intersects between our inverse functions and their corresponding circles
	A_tx := N_n(A, r_A, M, true)
	B_tx := N_n(B, r_B, M, true)
	A_bx := N_n(A, r_A, M, false)
	B_bx := N_n(B, r_B, M, false)

	// define the functions connecting the intersect points
	m_t := m(A_tx, B_tx)
	m_b := m(A_bx, B_bx)
	b_t := b(A_tx, B_tx, m_t)
	b_b := b(A_bx, B_bx, m_b)
	f_tx := fx(x, m_t, b_t)
	f_bx := fx(x, m_b, b_b)

	// test whether slime is within our connection bounds
	belowUpper := y < f_tx
	aboveLower := f_bx < y
	rightOfLeft := F_Ax < x
	leftOfRight := x < F_Bx
	if !(belowUpper && aboveLower && rightOfLeft && leftOfRight) {
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

	// update ce's physical representation
	for x := 0; x < windowSize.X; x++ {
		for y := 0; y < windowSize.Y; y++ {

			// prevColor := getColCoords(x, y, prevIM)
			pixel := getPixel(x, y)
			
			// check whether to render slime between joints
			for _, jointA := range ce.joints {
				intraJointSlime(jointA, pixel)
				for _, idx := range jointA.connectionIdxs {
					jointB := ce.joints[idx]
					interJointSlime(jointA, jointB, pixel)
				}
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

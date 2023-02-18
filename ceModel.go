package main

import (
	"fmt"
	"math"
	"time"
)

// render slime pixel if meets conditions
func intraEffectorSlime(effector *Effector, pixel *Pixel) {

	// initialize test case
	makeSlime := true

	// is it outside the circle
	if SqrDistance(pixel.pos, effector.pos) > math.Pow(effector.radius, 2) {
		makeSlime = false
	}

	// render pixel if passed
	if makeSlime {
		updatePixelCol(pixel, slimeCol)
	}

}

// render slime pixel if meets conditions
func interEffectorSlime(effectorA, effectorB *Effector, pixel *Pixel) {

	// initialize test case
	makeSlime := true

	// according with: https://www.desmos.com/calculator/cy12nq52sx
	// define our circle properties
	A := effectorA.pos
	B := effectorB.pos
	x := float64(pixel.pos.X)
	y := float64(pixel.pos.Y)
	r_A := effectorA.radius
	r_B := effectorB.radius

	// define the slope of our intersect function f(x) between the circle centers
	m_f := m(A, B)

	// define the inverse functions of f(x) F_A(X) and F_A(X) intersecting with A or B, respectively
	M := M(m_f)
	F_Ax := F_Nx(x, M, A)
	F_Bx := F_Nx(x, M, B)

	// define the intersects between our inverse functions and their corresponding circles
	A_t := N_n(A, r_A, M, true)
	B_t := N_n(B, r_B, M, true)
	A_b := N_n(A, r_A, M, false)
	B_b := N_n(B, r_B, M, false)

	// define the functions connecting the intersect points
	m_t := m(A_t, B_t)
	m_b := m(A_b, B_b)
	b_t := b(A_t, B_t, m_t)
	b_b := b(A_b, B_b, m_b)
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

	fmt.Printf("A.Y: %d\nA.Y: %d\nB.X: %d\nB.Y: %d\nx: %g\ny: %g\nr_A: %g\nr_B: %g\nm_f: %g\nM: %g\nF_Ax: %g\nF_Bx: %g\nA_t.X: %d\nB_t.X: %d\nA_b.X: %d\nB_b.X: %d\nA_t.Y: %d\nB_t.Y: %d\nA_b.Y: %d\nB_b.Y: %d\nm_t: %g\nm_b: %g\nb_t: %g\nb_b: %g\nf_tx: %g\nf_bx: %g\n", A.Y, A.Y, B.X, B.Y, x, y, r_A, r_B, m_f, M, F_Ax, F_Bx, A_t.X, B_t.X, A_b.X, B_b.X, A_t.Y, B_t.Y, A_b.Y, B_b.Y, m_t, m_b, b_t, b_b, f_tx, f_bx)

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
			
			// check whether to render slime between effectors
			for _, effectorA := range ce.effectors {
				intraEffectorSlime(effectorA, pixel)
				for _, idx := range effectorA.connectionIdxs {
					effectorB := ce.effectors[idx]
					interEffectorSlime(effectorA, effectorB, pixel)
					time.Sleep(time.Duration(time.Second * 10000))
				}
			}
			
		}
	}

}

func ceModelUpdateLeft() {
	ce.effectors[0].pos.X -= 10
	ceModelCall()
}

func ceModelUpdateRight() {
	ce.effectors[0].pos.X += 10
	ceModelCall()
}

func ceModelUpdateUp() {
	ce.effectors[0].pos.Y -= 10
	ceModelCall()
}

func ceModelUpdateDown() {
	ce.effectors[0].pos.Y += 10
	ceModelCall()
}

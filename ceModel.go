package main

import (
	"fmt"
)

// render slime pixel if meets conditions
func intraEffectorSlime(circle *Effector, pixel *Pixel) {

	// initialize test case
	makeSlime := true

	// is it outside the circle
	if SqrMagnitude(pixel.pos, circle.pos) > circle.sqrRadius {
		makeSlime = false
	}

	// render pixel if passed
	if makeSlime {
		updatePixelCol(pixel, slimeCol)
	}

}

// render slime pixel if meets conditions
func interEffectorSlime(connection *Polygon, pixel *Pixel) {

	// initialize test case
	makeSlime := true

	x := float64(pixel.pos.X)                 // left to right in window
	y := yMathToYWindow(float64(pixel.pos.Y)) // top to bottom in window

	// calculate the inverse func
	F_Ax := F_Nx(x, M, A)
	F_Bx := F_Nx(x, M, B)

	// define the functions connecting the intersect points
	m_1 := m(A_1, B_1)
	m_2 := m(A_2, B_2)
	b_1 := b(A_1, B_1, m_1)
	b_2 := b(A_2, B_2, m_2)
	f_1x := fx(x, m_1, b_1)
	f_2x := fx(x, m_2, b_2)

	// test whether slime is within our connection bounds (only god knows why this works)
	cond1 := false
	if (r_A > r_B && A.Y > B.Y && A_1.X < B_1.X && B_2.X < A_2.X) || (r_B > r_A && A.Y > B.Y && B_1.X < A_1.X && A_2.X < B_2.X) {
		fmt.Println("1")
		cond1 = y <= f_1x && y <= f_2x
	} else if (r_A > r_B && B.Y > A.Y && A_1.X < B_1.X && B_2.X < A_2.X) || (r_B > r_A && B.Y > A.Y && B_1.X < A_1.X && A_2.X < B_2.X) {
		fmt.Println("2")
		cond1 = y >= f_1x && y >= f_2x
	} else if B.X < A.X || A.Y == B.Y {
		fmt.Println("3")
		cond1 = f_1x <= y && y <= f_2x
	} else if A.X < B.X {
		fmt.Println("4")
		cond1 = f_2x <= y && y <= f_1x
	}
	cond2 := (F_Ax <= y && y <= F_Bx) || (F_Bx <= y && y <= F_Ax)
	if !(cond1 && cond2) {
		makeSlime = false
	}

	// fmt.Printf("A.X: %d\nA.Y: %d\nB.X: %d\nB.Y: %d\nx: %g\ny: %g\nr_A: %g\nr_B: %g\nm_f: %g\nM: %g\nF_Ax: %g\nF_Bx: %g\nA_1.X: %d\nA_1.Y: %d\nB_1.X: %d\nB_1.Y: %d\nA_2.X: %d\nA_2.Y: %d\nB_2.X: %d\nB_2.Y: %d\nm_1: %g\nm_2: %g\nb_1: %g\nb_2: %g\nf_1x: %g\nf_2x: %g\n", A.X, A.Y, B.X, B.Y, x, y, r_A, r_B, m_f, M, F_Ax, F_Bx, A_1.X, A_1.Y, B_1.X, B_1.Y, A_2.X, A_2.Y, B_2.X, B_2.Y, m_1, m_2, b_1, b_2, f_1x, f_2x)

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
			A := NewVector2(effectorA.pos.X, int(yMathToYWindow(float64(effectorA.pos.Y))))
			B := NewVector2(effectorB.pos.X, int(yMathToYWindow(float64(effectorB.pos.Y))))
			r_A := effectorA.radius
			r_B := effectorB.radius

			// define the slope of our intersect function f(x) between the circle centers
			m_f := m(A, B)

			// define the inverse functions of f(x) F_A(X) and F_A(X) intersecting with A or B, respectively
			M := M(m_f)

			// define the intersects between our inverse functions and their corresponding circles
			A_1 := N_n(A, A, B, r_A, M, false)
			B_1 := N_n(B, A, B, r_B, M, false)
			A_2 := N_n(A, A, B, r_A, M, true)
			B_2 := N_n(B, A, B, r_B, M, true)

			// define the connection zone as a polygon
			connection := MakePolygon([]*Vector2 {A_1, B_1, B_2, A_2})

			// update ce's circle connection physical representations
			for x := 0; x < windowSize.X; x++ {
				for y := 0; y < windowSize.Y; y++ {
					interEffectorSlime(connection, getPixel(x, y))
				}
			}
			
		}

		// update ce's circle physical representations
		for x := 0; x < windowSize.X; x++ {
			for y := 0; y < windowSize.Y; y++ {
				intraEffectorSlime(effectorA, getPixel(x, y))
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

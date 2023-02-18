package main

import "math"

func SqrDistance(v1, v2 *Vector2) float64 {
	// square distance comparison is far more optimized than dist comparison since we aren't taking the square root of 2
	xSqrDist := (v1.X - v2.X) * (v1.X - v2.X)
	ySqrDist := (v1.Y - v2.Y) * (v1.Y - v2.Y)
	return float64(xSqrDist) + float64(ySqrDist)
}

func m(A, B *Vector2) float64 {
	if B.X - A.X == 0 {
		return inf
	} else {
		return float64((B.Y - A.Y) / (B.X - A.X))
	}
}

func b(A, B *Vector2, m float64) float64 {
	return float64(A.Y) - m*float64(A.X)
}

func fx(x, m, b float64) float64 {
	return m*x + b
}

func M(m float64) float64 {
	return -(1 / m)
}

// calculates the inverse function of fx for a circle
func F_Nx(x, M float64, N *Vector2) float64 {
	return M*(x-float64(N.X)) + float64(N.Y)
}

// calculates the top point
func N_n(N *Vector2, r_N, M float64, tNotb bool) (V *Vector2) {
	mult := 1.0
	if !tNotb {
		mult = -1
	}
	V_x := 0
	if 1+math.Pow(M,2) == 0 {
		V_x = int(inf)
	} else {
		V_x = int(math.Round(float64(N.X) + mult*r_N*math.Sqrt(1/(1+math.Pow(M,2)))))
	}
	V = NewVector2(
		V_x,
		int(math.Round(F_Nx(float64(V_x), M, N))),
	)
	return
}

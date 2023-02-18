package main

import "math"

func yMathToYWindow(y float64) float64 {
	return float64(windowSize.Y) - y
}

func SqrMagnitude(v1, v2 *Vector2) float64 {
	xDiff := v2.X - v1.X
	yDiff := v2.Y - v1.Y
	xSqrDist := xDiff * xDiff
	ySqrDist := yDiff * yDiff
	return float64(xSqrDist) + float64(ySqrDist)
}

func m(A, B *Vector2) float64 {
	xDist := B.X - A.X
	yDist := B.Y - A.Y
	if xDist == 0 {
		return inf
	} else if yDist == 0 {
		return 0
	} else {
		return float64(yDist) / float64(xDist)
	}
}

func b(A, B *Vector2, m float64) float64 {
	return float64(A.Y) - m*float64(A.X)
}

func fx(x, m, b float64) float64 {
	return m*x + b
}

func M(m float64) float64 {
	if m == 0 {
		return inf
	} else {
		return -(1 / m)
	}
}

// calculates the inverse function of fx for a circle
func F_Nx(x, M float64, N *Vector2) float64 {
	return M*(x-float64(N.X)) + float64(N.Y)
}

// calculates the top point
func N_n(N, A, B *Vector2, r_N, M float64, bNotT bool) (V *Vector2) {
	mult := 0.0
	if bNotT {
		if A.Y < B.Y {
			mult = 1
		} else {
			mult = -1
		}
	} else {
		if A.Y < B.Y {
			mult = -1
		} else {
			mult = 1
		}
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

func InPolygon(polygon *Polygon, x, y float64) (in bool) {
	in = true
	return
}
